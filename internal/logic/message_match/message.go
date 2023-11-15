package message_match

import (
	"context"

	"gorm.io/gorm"

	"github.com/scroll-tech/chain-monitor/internal/orm"
	"github.com/scroll-tech/chain-monitor/internal/types"
)

// MessageMatchLogic define the message match logics
type MessageMatchLogic struct {
	messageMatchOrm *orm.MessageMatch
}

func NewMessagesMatchLogic(db *gorm.DB) *MessageMatchLogic {
	return &MessageMatchLogic{
		messageMatchOrm: orm.NewMessageMatch(db),
	}
}

// GetLatestBlockNumber get the latest number
func (t *MessageMatchLogic) GetLatestBlockNumber(ctx context.Context, layer types.LayerType) (uint64, error) {
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
