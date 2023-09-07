package l2watcher

import (
	"context"
	"fmt"

	"github.com/scroll-tech/go-ethereum/common"
	"github.com/scroll-tech/go-ethereum/core/types"
	"github.com/scroll-tech/go-ethereum/log"

	"chain-monitor/bytecode/scroll/L2"
	"chain-monitor/internal/orm"
	"chain-monitor/internal/utils"
)

func (l2 *l2Contracts) registerMessengerHandlers() {
	l2.ScrollMessenger.RegisterSentMessage(func(vLog *types.Log, data *L2.L2ScrollMessengerSentMessageEvent) error {
		msgHash := utils.ComputeMessageHash(l2.ScrollMessenger.ABI, data.Sender, data.Target, data.Value, data.MessageNonce, data.Message)
		number := vLog.BlockNumber
		l2.txHashMsgHash[vLog.TxHash.String()] = msgHash
		l2.msgSentEvents[number] = append(l2.msgSentEvents[number], &orm.L2MessengerEvent{
			Number:   number,
			MsgHash:  msgHash.String(),
			Type:     orm.L2SentMessage,
			MsgNonce: data.MessageNonce.Uint64(),
		})
		return nil
	})
	l2.ScrollMessenger.RegisterRelayedMessage(func(vLog *types.Log, data *L2.L2ScrollMessengerRelayedMessageEvent) error {
		l2.txHashMsgHash[vLog.TxHash.String()] = data.MessageHash
		return orm.SaveL2Messenger(l2.tx, orm.L2RelayedMessage, vLog, data.MessageHash)
	})
	l2.ScrollMessenger.RegisterFailedRelayedMessage(func(vLog *types.Log, data *L2.L2ScrollMessengerFailedRelayedMessageEvent) error {
		l2.txHashMsgHash[vLog.TxHash.String()] = data.MessageHash
		return orm.SaveL2Messenger(l2.tx, orm.L2FailedRelayedMessage, vLog, data.MessageHash)
	})
}

func (l2 *l2Contracts) storeMessengerEvents(ctx context.Context, start, end uint64) error {
	if len(l2.msgSentEvents) == 0 {
		return nil
	}

	// Calculate withdraw root.
	var (
		chainMonitors = make([]*orm.L2ChainConfirm, 0, end-start+1)
		msgSentEvents []*orm.L2MessengerEvent
	)
	for number := start; number <= end; number++ {
		if l2.msgSentEvents[number] == nil {
			chainMonitors = append(chainMonitors, &orm.L2ChainConfirm{
				Number:             number,
				WithdrawRootStatus: true,
			})
			continue
		}
		msgs := l2.msgSentEvents[number]
		for i, msg := range msgs {
			proofs := l2.withdraw.AppendMessages([]common.Hash{common.HexToHash(msg.MsgHash)})
			// Store the latest one for every block.
			if i == len(msgs)-1 {
				msg.MsgProof = common.Bytes2Hex(proofs[0])
			}
			msgSentEvents = append(msgSentEvents, msgs[i])
		}
		chainMonitors = append(chainMonitors, &orm.L2ChainConfirm{
			Number:       number,
			WithdrawRoot: l2.withdraw.MessageRoot(),
		})
	}

	// Check withdraw root and store confirm monitor.
	if err := l2.storeWithdrawRoots(ctx, chainMonitors); err != nil {
		return err
	}

	// Store messenger events.
	if err := l2.tx.Model(&orm.L2MessengerEvent{}).Save(msgSentEvents).Error; err != nil {
		return err
	}
	return nil
}

func (l2 *l2Contracts) storeWithdrawRoots(ctx context.Context, chainMonitors []*orm.L2ChainConfirm) error {
	var (
		numbers       []uint64
		withdrawRoots []common.Hash
		err           error
	)
	for _, monitor := range chainMonitors {
		if !monitor.WithdrawRootStatus {
			numbers = append(numbers, monitor.Number)
		}
	}

	utils.TryTimes(3, func() bool {
		// get withdraw root by batch.
		withdrawRoots, err = utils.GetBatchWithdrawRoots(ctx, l2.rpcCli, l2.MessageQueue.Address, numbers)
		return err == nil
	})
	if err != nil {
		return err
	}

	for _, monitor := range chainMonitors {
		if len(withdrawRoots) == 0 {
			break
		}
		if monitor.WithdrawRootStatus {
			continue
		}
		expectRoot := withdrawRoots[0]
		withdrawRoots = withdrawRoots[1:]
		monitor.WithdrawRootStatus = monitor.WithdrawRoot == expectRoot
		// If the withdraw root doesn't match, alert it.
		if !monitor.WithdrawRootStatus {
			msg := fmt.Sprintf(
				"withdraw root doestn't match, l2_number: %d, expect_withdraw_root: %s, actual_withdraw_root: %s",
				monitor.Number,
				expectRoot.String(),
				monitor.WithdrawRoot.String(),
			)
			log.Error("withdraw root doesn't match", "number", monitor.Number, "expect_root", expectRoot.String(), "actual_root", monitor.WithdrawRoot.String())
			go l2.monitorAPI.SlackNotify(msg)
		}
	}
	if err = l2.tx.Model(&orm.L2ChainConfirm{}).Save(chainMonitors).Error; err != nil {
		return err
	}
	return nil
}
