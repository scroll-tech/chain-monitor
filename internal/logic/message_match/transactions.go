package message_match

import (
	"context"

	"gorm.io/gorm"

	"github.com/scroll-tech/chain-monitor/internal/orm"
	"github.com/scroll-tech/chain-monitor/internal/types"
)

// MessageMatchLogic define the message match logics
type MessageMatchLogic struct {
	transactionMatchOrm *orm.MessageMatch
}

func NewTransactionsMatchLogic(db *gorm.DB) *MessageMatchLogic {
	return &MessageMatchLogic{
		transactionMatchOrm: orm.NewMessageMatch(db),
	}
}

// GetLatestBlockNumber get the latest number
func (t *MessageMatchLogic) GetLatestBlockNumber(ctx context.Context, layer types.LayerType) (uint64, error) {
	blockValidTransactionMatch, blockValidErr := t.transactionMatchOrm.GetLatestBlockValidMessageMatch(ctx, layer)
	if blockValidErr != nil {
		return 0, blockValidErr
	}

	var number uint64
	switch layer {
	case types.Layer1:
		number = blockValidTransactionMatch.L1BlockNumber
	case types.Layer2:
		number = blockValidTransactionMatch.L2BlockNumber
	}

	return number, nil
}
