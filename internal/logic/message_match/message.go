package messagematch

import (
	"context"

	"gorm.io/gorm"

	"github.com/scroll-tech/chain-monitor/internal/config"
	"github.com/scroll-tech/chain-monitor/internal/orm"
	"github.com/scroll-tech/chain-monitor/internal/types"
)

// LogicMessageMatch defines the logic related to message matching.
type LogicMessageMatch struct {
	conf            *config.Config
	messageMatchOrm *orm.MessageMatch
}

// NewMessageMatchLogic initializes a new instance of Logic with an instance of orm.MessageMatch.
func NewMessageMatchLogic(cfg *config.Config, db *gorm.DB) *LogicMessageMatch {
	return &LogicMessageMatch{
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

	if layer == types.Layer2 && blockValidMessageMatch == nil {
		return t.conf.L1Config.StartNumber, nil
	}

	if layer == types.Layer1 && blockValidMessageMatch == nil {
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
