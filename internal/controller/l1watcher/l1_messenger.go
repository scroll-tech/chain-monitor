package l1watcher

import (
	"chain-monitor/bytecode"
	"chain-monitor/bytecode/scroll/L1"
	"chain-monitor/bytecode/scroll/L1/rollup"
	"chain-monitor/internal/orm"
	"chain-monitor/internal/utils"
	"context"
	"github.com/scroll-tech/go-ethereum/common"
	"github.com/scroll-tech/go-ethereum/core/types"
)

func (l1 *l1Contracts) registerMessengerHandlers() {
	l1.ScrollMessenger.RegisterSentMessage(func(vLog *types.Log, data *L1.L1ScrollMessengerSentMessageEvent) error {
		msgHash := utils.ComputeMessageHash(data.Sender, data.Target, data.Value, data.MessageNonce, data.Message)
		l1.msgSentEvents[vLog.TxHash.String()] = &orm.L1MessengerEvent{
			TxHead: &orm.TxHead{
				Number:  vLog.BlockNumber,
				TxHash:  vLog.TxHash.String(),
				MsgHash: msgHash.String(),
				Type:    orm.L1SentMessage,
			},
			Amount:  data.Value,
			Target:  data.Target,
			Message: data.Message,
			Log:     vLog,
		}
		return nil
	})
	l1.ScrollMessenger.RegisterRelayedMessage(func(vLog *types.Log, data *L1.L1ScrollMessengerRelayedMessageEvent) error {
		txHash := vLog.TxHash.String()
		l1.msgSentEvents[txHash] = &orm.L1MessengerEvent{
			TxHead: &orm.TxHead{
				Number:  vLog.BlockNumber,
				TxHash:  txHash,
				MsgHash: common.BytesToHash(data.MessageHash[:]).String(),
				Type:    orm.L1RelayedMessage,
			},
			Log: vLog,
		}
		return nil
	})
	l1.ScrollMessenger.RegisterFailedRelayedMessage(func(vLog *types.Log, data *L1.L1ScrollMessengerFailedRelayedMessageEvent) error {
		txHash := vLog.TxHash.String()
		l1.msgSentEvents[txHash] = &orm.L1MessengerEvent{
			TxHead: &orm.TxHead{
				Number:  vLog.BlockNumber,
				TxHash:  txHash,
				MsgHash: common.BytesToHash(data.MessageHash[:]).String(),
				Type:    orm.L1FailedRelayedMessage,
			},
			Log: vLog,
		}
		return nil
	})
}

func (l1 *l1Contracts) registerScrollHandlers() {
	l1.ScrollChain.RegisterCommitBatch(func(vLog *types.Log, data *rollup.ScrollChainCommitBatchEvent) error {
		l1Tx, _, err := l1.client.TransactionByHash(context.Background(), vLog.TxHash)
		if err != nil {
			return err
		}
		scrollBatch, err := bytecode.GetBatchRangeFromCalldataV2(l1.ScrollChain.ABI, l1Tx.Data())
		if err != nil {
			return err
		}
		return l1.tx.Create(&orm.L1ScrollChainEvent{
			Number:        vLog.BlockNumber,
			BatchIndex:    data.BatchIndex.Uint64(),
			BatchHash:     common.BytesToHash(data.BatchHash[:]).String(),
			CommitNumber:  vLog.BlockNumber,
			CommitHash:    vLog.TxHash.String(),
			L2StartNumber: scrollBatch.L2StartNumber,
			L2EndNumber:   scrollBatch.L2EndNumber,
			Status:        orm.BatchCommit,
		}).Error
	})
	l1.ScrollChain.RegisterFinalizeBatch(func(vLog *types.Log, data *rollup.ScrollChainFinalizeBatchEvent) error {
		eventMsg := orm.L1ScrollChainEvent{
			BatchHash:      common.BytesToHash(data.BatchHash[:]).String(),
			FinalizeNumber: vLog.BlockNumber,
			FinalizeHash:   vLog.TxHash.String(),
			Status:         orm.BatchFinalize,
		}
		return l1.tx.Select("finalize_number", "finalize_hash", "status").Updates(&eventMsg).Error
	})
	l1.ScrollChain.RegisterRevertBatch(func(vLog *types.Log, data *rollup.ScrollChainRevertBatchEvent) error {
		eventMsg := orm.L1ScrollChainEvent{
			BatchHash:    common.BytesToHash(data.BatchHash[:]).String(),
			RevertNumber: vLog.BlockNumber,
			RevertHash:   vLog.TxHash.String(),
			Status:       orm.BatchRevert,
		}
		return l1.tx.Select("revert_number", "revert_hash", "status").Updates(&eventMsg).Error
	})
}

func (l1 *l1Contracts) storeL1WatcherEvents() error {
	// store l1 eth events.
	if len(l1.ethEvents) > 0 {
		for _, event := range l1.ethEvents {
			if msg, exist := l1.msgSentEvents[event.TxHash]; exist {
				event.MsgHash = msg.MsgHash
			}
		}
		if err := l1.tx.Save(l1.ethEvents).Error; err != nil {
			return err
		}
	}

	// store l1 erc20 events.
	if len(l1.erc20Events) > 0 {
		for _, event := range l1.erc20Events {
			if msg, exist := l1.msgSentEvents[event.TxHash]; exist {
				event.MsgHash = msg.MsgHash
			}
		}
		if err := l1.tx.Save(l1.erc20Events).Error; err != nil {
			return err
		}
	}

	// store l1 err721 events.
	if len(l1.erc721Events) > 0 {
		for _, event := range l1.erc721Events {
			if msg, exist := l1.msgSentEvents[event.TxHash]; exist {
				event.MsgHash = msg.MsgHash
			}
		}
		if err := l1.tx.Save(l1.erc721Events).Error; err != nil {
			return err
		}
	}

	// store l1 erc1155 events.
	if len(l1.erc1155Events) > 0 {
		for _, event := range l1.erc1155Events {
			if msg, exist := l1.msgSentEvents[event.TxHash]; exist {
				event.MsgHash = msg.MsgHash
			}
		}
		if err := l1.tx.Save(l1.erc1155Events).Error; err != nil {
			return err
		}
	}

	// store l1 scroll_messenger sentMessage events.
	if length := len(l1.msgSentEvents); length > 0 {
		var msgs = make([]*orm.L1MessengerEvent, 0, length)
		for k := range l1.msgSentEvents {
			msgs = append(msgs, l1.msgSentEvents[k])
		}
		if err := l1.tx.Save(msgs).Error; err != nil {
			return err
		}
	}

	return nil
}
