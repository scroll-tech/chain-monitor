package logic

import (
	"context"

	"github.com/scroll-tech/go-ethereum/common"
	"github.com/scroll-tech/go-ethereum/core/types"

	"chain-monitor/bytecode/scroll/L2"
	"chain-monitor/internal/utils"
	"chain-monitor/orm"
)

func (l2 *L2Contracts) registerMessengerHandlers() {
	l2.ScrollMessenger.RegisterSentMessage(func(vLog *types.Log, data *L2.L2ScrollMessengerSentMessageEvent) error {
		msgHash := computeMessageHash(l2.ScrollMessenger.ABI, data.Sender, data.Target, data.Value, data.MessageNonce, data.Message)
		l2.txHashMsgHash[vLog.TxHash.String()] = msgHash
		l2.msgSentEvents[vLog.BlockNumber] = append(l2.msgSentEvents[vLog.BlockNumber], &orm.L2MessengerEvent{
			Number:   vLog.BlockNumber,
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

func (l2 *L2Contracts) storeMessengerEvents(ctx context.Context, start, end uint64) error {
	if len(l2.msgSentEvents) == 0 {
		return nil
	}

	// Calculate withdraw root.
	var (
		chainMonitors = make([]*orm.ChainConfirm, 0, end-start+1)
		msgSentEvents []*orm.L2MessengerEvent
		latestProof   []byte
	)
	for number := start; number <= end; number++ {
		if l2.msgSentEvents[number] == nil {
			chainMonitors = append(chainMonitors, &orm.ChainConfirm{
				Number:         number,
				WithdrawStatus: true,
			})
			continue
		}
		msgs := l2.msgSentEvents[number]
		for i, msg := range msgs {
			proofs := l2.withdraw.AppendMessages([]common.Hash{common.HexToHash(msg.MsgHash)})
			latestProof = proofs[len(proofs)-1]
			msgSentEvents = append(msgSentEvents, msgs[i])
		}
		chainMonitors = append(chainMonitors, &orm.ChainConfirm{
			Number:       number,
			WithdrawRoot: l2.withdraw.MessageRoot(),
		})
	}

	// Check withdraw root and store confirm monitor.
	if err := l2.storeWithdrawRoots(ctx, chainMonitors); err != nil {
		return err
	}

	// Just store the latest proof.
	msgSentEvents[len(msgSentEvents)-1].MsgProof = common.Bytes2Hex(latestProof)
	// Store messenger events.
	if err := l2.tx.Model(&orm.L2MessengerEvent{}).Save(msgSentEvents).Error; err != nil {
		return err
	}
	return nil
}

func (l2 *L2Contracts) storeWithdrawRoots(ctx context.Context, chainMonitors []*orm.ChainConfirm) error {
	var (
		numbers       []uint64
		withdrawRoots []common.Hash
		err           error
	)
	for _, monitor := range chainMonitors {
		if !monitor.WithdrawStatus {
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
		if !monitor.WithdrawStatus {
			monitor.WithdrawStatus = monitor.WithdrawRoot == withdrawRoots[0]
			withdrawRoots = withdrawRoots[1:]
		}
	}
	if err = l2.tx.Model(&orm.ChainConfirm{}).Save(chainMonitors).Error; err != nil {
		return err
	}
	return nil
}
