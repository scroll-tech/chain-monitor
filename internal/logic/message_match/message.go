package messagematch

import (
	"context"
	"fmt"

	"gorm.io/gorm"

	"github.com/scroll-tech/chain-monitor/internal/config"
	"github.com/scroll-tech/chain-monitor/internal/orm"
	"github.com/scroll-tech/chain-monitor/internal/types"
)

// LogicMessageMatch defines the logic related to message matching.
type LogicMessageMatch struct {
	db              *gorm.DB
	conf            *config.Config
	messageMatchOrm *orm.MessageMatch
}

// NewMessageMatchLogic initializes a new instance of Logic with an instance of orm.MessageMatch.
func NewMessageMatchLogic(cfg *config.Config, db *gorm.DB) *LogicMessageMatch {
	return &LogicMessageMatch{
		db:              db,
		conf:            cfg,
		messageMatchOrm: orm.NewMessageMatch(db),
	}
}

// GetLatestBlockNumber retrieves the latest block number for a given layer type.
func (t *LogicMessageMatch) GetLatestBlockNumber(ctx context.Context, layer types.LayerType) (uint64, error) {
	blockValidMessageMatch, blockValidErr := t.messageMatchOrm.GetLatestBlockValidMessageMatch(ctx, layer)
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
			effectRow, err := t.messageMatchOrm.InsertOrUpdateEventInfo(ctx, layer, message, tx)
			if err != nil {
				return fmt.Errorf("event orm insert failed, err: %w, layer:%s", err, layer.String())
			}
			effectRows += effectRow
		}
		return nil
	})
	if err != nil {
		return fmt.Errorf("insert or update event info failed, err:%w", err)
	}

	if int(effectRows) != len(messengerMessageMatches) {
		return fmt.Errorf("messenger event orm insert failed, effectRow:%d not equal messageMatches:%d", effectRows, len(messengerMessageMatches))
	}
	return nil
}
