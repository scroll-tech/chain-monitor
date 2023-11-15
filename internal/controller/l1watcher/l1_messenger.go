package l1watcher

import (
	"context"

	"github.com/scroll-tech/go-ethereum/common"
	"github.com/scroll-tech/go-ethereum/core/types"

	"chain-monitor/bytecode"
	"chain-monitor/bytecode/scroll/L1"
	"chain-monitor/bytecode/scroll/L1/rollup"
	"chain-monitor/internal/orm"
	"chain-monitor/internal/utils"
)

func (l1 *l1Contracts) registerMessengerHandlers() {
	l1.ScrollMessenger.RegisterSentMessage(func(vLog *types.Log, data *L1.L1ScrollMessengerSentMessageEvent) error {
		msgHash := utils.ComputeMessageHash(data.Sender, data.Target, data.Value, data.MessageNonce, data.Message)
		l1.txHashMsgHash[vLog.TxHash.String()] = msgHash
		l1.msgSentEvents[vLog.BlockNumber] = append(l1.msgSentEvents[vLog.BlockNumber], &orm.L1MessengerEvent{
			TxHead: orm.NewTxHead(vLog, orm.L1SentMessage),
			Data:   data,
		})
		return orm.SaveL1Messenger(l1.tx, orm.L1SentMessage, vLog, msgHash)
	})
	l1.ScrollMessenger.RegisterRelayedMessage(func(vLog *types.Log, data *L1.L1ScrollMessengerRelayedMessageEvent) error {
		msgHash := common.BytesToHash(data.MessageHash[:])
		l1.txHashMsgHash[vLog.TxHash.String()] = msgHash
		return orm.SaveL1Messenger(l1.tx, orm.L1RelayedMessage, vLog, msgHash)
	})
	l1.ScrollMessenger.RegisterFailedRelayedMessage(func(vLog *types.Log, data *L1.L1ScrollMessengerFailedRelayedMessageEvent) error {
		msgHash := common.BytesToHash(data.MessageHash[:])
		l1.txHashMsgHash[vLog.TxHash.String()] = msgHash
		return orm.SaveL1Messenger(l1.tx, orm.L1FailedRelayedMessage, vLog, msgHash)
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
