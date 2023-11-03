package cross_chain

import (
	"context"
	"time"

	"github.com/scroll-tech/go-ethereum/log"
	"gorm.io/gorm"

	"github.com/scroll-tech/chain-monitor/internal/logic/checker"
	"github.com/scroll-tech/chain-monitor/internal/orm"
	"github.com/scroll-tech/chain-monitor/internal/types"
	"github.com/scroll-tech/chain-monitor/internal/utils"
)

// CrossChainLogic check the l1/l2 event match from db
// FinalizeWithdraw value ⇒ withdraw value (L1FinalizeWithdrawETH, L1FinalizeWithdrawERC721, L1FinalizeWithdrawERC1155), missing FinalizeBatchWithdraw.
// FinalizeDeposit value ⇒ deposit value (L2FinalizeDepositETH, L2FinalizeDepositERC721, L2FinalizeDepositERC1155), missing FinalizeBatchDeposit
// sent/relayed message
type CrossChainLogic struct {
	transactionOrm *orm.TransactionMatch
	checker        *checker.Checker
}

func NewCrossChainLogic(db *gorm.DB) *CrossChainLogic {
	return &CrossChainLogic{
		transactionOrm: orm.NewTransactionMatch(db),
		checker:        checker.NewChecker(db),
	}
}

func (c *CrossChainLogic) Fetcher(ctx context.Context, layerType types.LayerType, start, end uint64) {
	transactions, err := c.transactionOrm.GetLatestTransactionMatch(ctx, 100)
	if err != nil {
		log.Error("CrossChainLogic.Fetcher failed", "error", err)
		return
	}

	var transactionMatchIds []int64
	for _, transaction := range transactions {
		switch layerType {
		case types.Layer1:
			if transaction.L1BlockNumber < start || transaction.L1BlockNumber > end {
				continue
			}
		case types.Layer2:
			if transaction.L2BlockNumber < start || transaction.L2BlockNumber > end {
				continue
			}
		}

		// todo remove or not?
		if utils.NowUTC().Sub(transaction.CreatedAt) > 2*time.Hour {
			log.Error("id:%d don't check more than 10 minutes", transaction.ID)
		}

		checkResult := c.checker.CrossChainCheck(ctx, layerType, transaction)
		if checkResult == types.MismatchTypeOk {
			transactionMatchIds = append(transactionMatchIds, transaction.ID)
			continue
		}

		// todo send to slack
	}

	if err := c.transactionOrm.UpdateCrossChainStatus(ctx, transactionMatchIds, layerType, types.CrossChainStatusTypeValid); err != nil {
		log.Error("CrossChainLogic.Fetcher UpdateCrossChainStatus failed", "error", err)
		return
	}
}
