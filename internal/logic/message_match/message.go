package messagematch

import (
	"context"

	"gorm.io/gorm"

	"github.com/scroll-tech/chain-monitor/internal/orm"
	"github.com/scroll-tech/chain-monitor/internal/types"
)

// Logic defines the logic related to message matching.
type Logic struct {
	messageMatchOrm *orm.MessageMatch
}

// NewLogic initializes a new instance of Logic with an instance of orm.MessageMatch.
func NewLogic(db *gorm.DB) *Logic {
	return &Logic{
		messageMatchOrm: orm.NewMessageMatch(db),
	}
}

// GetLatestBlockNumber retrieves the latest block number for a given layer type.
func (t *Logic) GetLatestBlockNumber(ctx context.Context, layer types.LayerType) (uint64, error) {
	blockValidMessageMatch, blockValidErr := t.messageMatchOrm.GetLatestBlockValidMessageMatch(ctx, layer)
	if blockValidErr != nil {
		return 0, blockValidErr
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
