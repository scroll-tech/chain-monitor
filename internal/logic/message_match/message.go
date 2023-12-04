package messagematch

import (
	"context"
	"fmt"

	"github.com/scroll-tech/go-ethereum/log"
	"gorm.io/gorm"

	"github.com/scroll-tech/chain-monitor/internal/config"
	"github.com/scroll-tech/chain-monitor/internal/orm"
	"github.com/scroll-tech/chain-monitor/internal/types"
)

// LogicMessageMatch defines the logic related to message matching.
type LogicMessageMatch struct {
	db                       *gorm.DB
	conf                     *config.Config
	gatewayMessageMatchOrm   *orm.GatewayMessageMatch
	messengerMessageMatchOrm *orm.MessengerMessageMatch
}

// NewMessageMatchLogic initializes a new instance of Logic with an instance of orm.GatewayMessageMatch/orm.MessengerMessageMatch
func NewMessageMatchLogic(cfg *config.Config, db *gorm.DB) *LogicMessageMatch {
	return &LogicMessageMatch{
		db:                       db,
		conf:                     cfg,
		gatewayMessageMatchOrm:   orm.NewGatewayMessageMatch(db),
		messengerMessageMatchOrm: orm.NewMessengerMessageMatch(db),
	}
}

// GetBlocksStatus get the status from start block number to end block number
func (t *LogicMessageMatch) GetBlocksStatus(ctx context.Context, startBlockNumber, endBlockNumber uint64) bool {
	gatewayMessageMatches, err := t.gatewayMessageMatchOrm.GetBlocksStatus(ctx, startBlockNumber, endBlockNumber)
	if err != nil {
		log.Error("LogicMessageMatch.gatewayMessageMatches failed", "start block number", startBlockNumber, "end block number", endBlockNumber, "error", err)
		return false
	}

	for _, gatewayMessageMatch := range gatewayMessageMatches {
		if gatewayMessageMatch.L2EventType != int(types.L2WithdrawERC20) &&
			gatewayMessageMatch.L2EventType != int(types.L2WithdrawERC721) &&
			gatewayMessageMatch.L2EventType != int(types.L2WithdrawERC1155) &&
			gatewayMessageMatch.L2EventType != int(types.L2BatchWithdrawERC721) &&
			gatewayMessageMatch.L2EventType != int(types.L2FinalizeBatchDepositERC1155) {
			if gatewayMessageMatch.L2BlockNumber == 0 || gatewayMessageMatch.L2BlockStatus != int(types.BlockStatusTypeValid) {
				return false
			}
		}

		if gatewayMessageMatch.L1BlockNumber != 0 {
			if gatewayMessageMatch.L2BlockStatus == 0 ||
				gatewayMessageMatch.L1BlockStatus != int(types.BlockStatusTypeValid) ||
				gatewayMessageMatch.L1CrossChainStatus != int(types.CrossChainStatusTypeValid) ||
				gatewayMessageMatch.L2CrossChainStatus != int(types.CrossChainStatusTypeValid) {
				return false
			}
		}
	}

	messengerMessageMatches, err := t.messengerMessageMatchOrm.GetBlocksStatus(ctx, startBlockNumber, endBlockNumber)
	if err != nil {
		log.Error("LogicMessageMatch.messengerMessageMatches failed", "start block number", startBlockNumber, "end block number", endBlockNumber, "error", err)
		return false
	}

	for _, messengerMessageMatch := range messengerMessageMatches {
		if messengerMessageMatch.L2EventType != int(types.L2SentMessage) {
			if messengerMessageMatch.L2BlockNumber == 0 ||
				messengerMessageMatch.L2BlockStatus != int(types.BlockStatusTypeValid) ||
				messengerMessageMatch.L2ETHBalanceStatus != int(types.ETHBalanceStatusTypeValid) {
				return false
			}
		}

		if messengerMessageMatch.L1BlockNumber != 0 {
			if messengerMessageMatch.L2BlockStatus == 0 ||
				messengerMessageMatch.L1BlockStatus != int(types.BlockStatusTypeValid) ||
				messengerMessageMatch.L1CrossChainStatus != int(types.CrossChainStatusTypeValid) ||
				messengerMessageMatch.L2CrossChainStatus != int(types.CrossChainStatusTypeValid) ||
				messengerMessageMatch.L1ETHBalanceStatus != int(types.ETHBalanceStatusTypeValid) ||
				messengerMessageMatch.L2ETHBalanceStatus != int(types.ETHBalanceStatusTypeValid) {
				return false
			}
		}
	}
	return true
}

// GetLatestBlockNumber retrieves the latest block number for a given layer type.
func (t *LogicMessageMatch) GetLatestBlockNumber(ctx context.Context, layer types.LayerType) (uint64, error) {
	blockValidMessageMatch, blockValidErr := t.messengerMessageMatchOrm.GetLatestBlockValidMessageMatch(ctx, layer)
	if blockValidErr != nil {
		return 0, blockValidErr
	}

	if layer == types.Layer1 && blockValidMessageMatch == nil {
		return t.conf.L1Config.StartNumber, nil
	}

	if layer == types.Layer2 && blockValidMessageMatch == nil {
		return 0, nil
	}

	var number uint64
	switch layer {
	case types.Layer1:
		number = blockValidMessageMatch.L1BlockNumber
	case types.Layer2:
		number = blockValidMessageMatch.L2BlockNumber
	}

	return number, nil
}

// InsertOrUpdateMessageMatches insert or update the gateway/messenger event info
func (t *LogicMessageMatch) InsertOrUpdateMessageMatches(ctx context.Context, layer types.LayerType, gatewayMessageMatches []orm.GatewayMessageMatch, messengerMessageMatches []orm.MessengerMessageMatch) error {
	var effectRows int64
	err := t.db.Transaction(func(tx *gorm.DB) error {
		for _, message := range messengerMessageMatches {
			effectRow, err := t.messengerMessageMatchOrm.InsertOrUpdateEventInfo(ctx, layer, message, tx)
			if err != nil {
				return fmt.Errorf("messenger event orm insert failed, err: %w, layer:%s", err, layer.String())
			}
			effectRows += effectRow
		}

		for _, message := range gatewayMessageMatches {
			effectRow, err := t.gatewayMessageMatchOrm.InsertOrUpdateEventInfo(ctx, layer, message, tx)
			if err != nil {
				return fmt.Errorf("gateway event orm insert failed, err: %w, layer:%s", err, layer.String())
			}
			effectRows += effectRow
		}
		return nil
	})
	if err != nil {
		return fmt.Errorf("insert or update event info failed, err:%w", err)
	}

	if int(effectRows) != len(messengerMessageMatches)+len(gatewayMessageMatches) {
		return fmt.Errorf("gateway and messenger event orm insert failed, effectRow:%d not equal messageMatches:%d", effectRows, len(messengerMessageMatches)+len(gatewayMessageMatches))
	}
	return nil
}
