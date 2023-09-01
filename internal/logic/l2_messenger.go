package logic

import (
	"chain-monitor/bytecode/scroll/L2"
	"chain-monitor/orm"
	"context"
	"fmt"
	"github.com/scroll-tech/go-ethereum/common"
	"github.com/scroll-tech/go-ethereum/core/types"
	"sort"
)

func (l2 *L2Contracts) registerMessengerHandlers() {
	l2.ScrollMessenger.RegisterSentMessage(func(vLog *types.Log, data *L2.L2ScrollMessengerSentMessageEvent) error {
		msgHash := computeMessageHash(l2.ScrollMessenger.ABI, data.Sender, data.Target, data.Value, data.MessageNonce, data.Message)
		l2.txHashMsgHash[vLog.TxHash.String()] = msgHash
		l2.msgSentEvents = append(l2.msgSentEvents, &orm.L2MessengerEvent{
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

func (l2 *L2Contracts) storeMessengerEvents(ctx context.Context) error {
	if len(l2.msgSentEvents) == 0 {
		return nil
	}
	// Store sentMessage events.
	sort.Slice(l2.msgSentEvents, func(i, j int) bool {
		return l2.msgSentEvents[i].MsgNonce < l2.msgSentEvents[j].MsgNonce
	})
	// Check
	if l2.withdraw.NextMessageNonce != l2.msgSentEvents[0].MsgNonce {
		return fmt.Errorf("msg nonce doesn't match, expect_nonce: %d, actual_nonce: %d", l2.withdraw.NextMessageNonce, l2.msgSentEvents[0].MsgNonce)
	}

	var (
		chainMonitors = make([]orm.ChainConfirm, len(l2.msgSentEvents))
		numbers       = make(map[uint64]bool, len(l2.msgSentEvents))
	)
	for i, msg := range l2.msgSentEvents {
		proofs := l2.withdraw.AppendMessages([]common.Hash{common.HexToHash(msg.MsgHash)})
		if len(proofs) > 0 && len(proofs[0]) > 0 {
			msg.MsgProof = common.Bytes2Hex(proofs[0])
		}
		if !numbers[msg.Number] {
			numbers[msg.Number] = true
			expectRoot, err := l2.withdrawRoot(ctx, msg.Number)
			if err != nil {
				return err
			}
			chainMonitors[i] = orm.ChainConfirm{
				Number:         msg.Number,
				WithdrawRoot:   expectRoot.String(),
				WithdrawStatus: expectRoot == l2.withdraw.MessageRoot(),
			}
		}
	}
	if err := l2.tx.Save(chainMonitors).Error; err != nil {
		return err
	}
	if err := l2.tx.Save(l2.msgSentEvents).Error; err != nil {
		return err
	}
	return nil
}
