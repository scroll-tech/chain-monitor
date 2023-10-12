package l2watcher

import (
	"context"
	"fmt"

	"github.com/scroll-tech/go-ethereum/common"
	"github.com/scroll-tech/go-ethereum/core/types"
	"github.com/scroll-tech/go-ethereum/log"

	"chain-monitor/bytecode/scroll/L2"
	"chain-monitor/internal/controller"
	"chain-monitor/internal/orm"
	"chain-monitor/internal/utils"
)

func (l2 *l2Contracts) registerMessengerHandlers() {
	l2.ScrollMessenger.RegisterSentMessage(func(vLog *types.Log, data *L2.L2ScrollMessengerSentMessageEvent) error {
		msgHash := utils.ComputeMessageHash(data.Sender, data.Target, data.Value, data.MessageNonce, data.Message)
		number := vLog.BlockNumber
		l2.txHashMsgHash[vLog.TxHash.String()] = msgHash
		l2.msgSentEvents[number] = append(l2.msgSentEvents[number], &orm.L2MessengerEvent{
			Target:   data.Target,
			Message:  data.Message,
			Log:      vLog,
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
	// Calculate withdraw root.
	var msgSentEvents []*orm.L2MessengerEvent
	for number := start; number <= end; number++ {
		if l2.msgSentEvents[number] == nil {
			l2.l2Confirms[number].WithdrawRootStatus = true
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
		l2.l2Confirms[number].WithdrawRoot = l2.withdraw.MessageRoot()
	}

	// Store messenger events.
	if len(msgSentEvents) > 0 {
		tx := l2.tx.WithContext(ctx)
		if err := tx.Model(&orm.L2MessengerEvent{}).Save(msgSentEvents).Error; err != nil {
			return err
		}
	}

	return nil
}

func (l2 *l2Contracts) storeWithdrawRoots(ctx context.Context) error {
	var (
		numbers       []uint64
		withdrawRoots = map[uint64]common.Hash{}
		l2Confirms    = make([]*orm.L2ChainConfirm, 0, len(l2.l2Confirms))
		err           error
	)
	for i, monitor := range l2.l2Confirms {
		if !monitor.WithdrawRootStatus {
			numbers = append(numbers, monitor.Number)
		}
		l2Confirms = append(l2Confirms, l2.l2Confirms[i])
	}

	utils.TryTimes(3, func() bool {
		if len(numbers) == 0 {
			return true
		}
		// get withdraw root by batch.
		var roots []common.Hash
		roots, err = utils.GetBatchWithdrawRoots(ctx, l2.rpcCli, l2.MessageQueue.Address, numbers)
		if err != nil {
			return false
		}
		for i, number := range numbers {
			withdrawRoots[number] = roots[i]
		}
		return true
	})
	if err != nil {
		return err
	}

	for _, monitor := range l2Confirms {
		if len(withdrawRoots) == 0 {
			break
		}
		if monitor.WithdrawRootStatus {
			continue
		}
		expectRoot := withdrawRoots[monitor.Number]
		monitor.WithdrawRootStatus = monitor.WithdrawRoot == expectRoot
		// If the withdraw root doesn't match, alert it.
		if !monitor.WithdrawRootStatus {
			controller.WithdrawRootMisMatchTotal.Inc()
			msg := fmt.Sprintf(
				"withdraw root doestn't match, l2_number: %d, expect_withdraw_root: %s, actual_withdraw_root: %s",
				monitor.Number,
				expectRoot.String(),
				monitor.WithdrawRoot.String(),
			)
			log.Error("withdraw root doesn't match", "Number", monitor.Number, "expect_root", expectRoot.String(), "actual_root", monitor.WithdrawRoot.String())
			go controller.SlackNotify(msg)
		}
	}
	if err = l2.tx.Model(&orm.L2ChainConfirm{}).Save(l2Confirms).Error; err != nil {
		return err
	}
	return nil
}
