package cross_chain

import (
	"context"

	"github.com/scroll-tech/go-ethereum/log"
	"gorm.io/gorm"

	"github.com/scroll-tech/chain-monitor/internal/logic/checker"
	"github.com/scroll-tech/chain-monitor/internal/orm"
	"github.com/scroll-tech/chain-monitor/internal/types"
)

// CrossChainLogic check the l1/l2 event match from db
// FinalizeWithdraw value ⇒ withdraw value.
// FinalizeDeposit value ⇒ deposit value.
// This is due to the fact that not every deposit/withdrawal event in the system will have a finalize event,
// as users have the capability to independently refund deposits.
type CrossChainLogic struct {
	messageOrm *orm.MessageMatch
	checker    *checker.Checker
}

func NewCrossChainLogic(db *gorm.DB) *CrossChainLogic {
	return &CrossChainLogic{
		messageOrm: orm.NewMessageMatch(db),
		checker:    checker.NewChecker(db),
	}
}

func (c *CrossChainLogic) Fetcher(ctx context.Context, layerType types.LayerType, start, end uint64) {
	messages, err := c.messageOrm.GetLatestMessageMatch(ctx, 100)
	if err != nil {
		log.Error("CrossChainLogic.Fetcher failed", "error", err)
		return
	}

	var messageMatchIds []int64
	for _, message := range messages {
		switch layerType {
		case types.Layer1:
			if message.L1BlockNumber < start || message.L1BlockNumber > end {
				continue
			}
		case types.Layer2:
			if message.L2BlockNumber < start || message.L2BlockNumber > end {
				continue
			}
		}

		checkResult := c.checker.CrossChainCheck(ctx, layerType, message)
		if checkResult == types.MismatchTypeOk {
			messageMatchIds = append(messageMatchIds, message.ID)
			continue
		}

		// todo send to slack
	}

	if err := c.messageOrm.UpdateCrossChainStatus(ctx, messageMatchIds, layerType, types.CrossChainStatusTypeValid); err != nil {
		log.Error("CrossChainLogic.Fetcher UpdateCrossChainStatus failed", "error", err)
		return
	}
}
