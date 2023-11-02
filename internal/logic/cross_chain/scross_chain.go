package cross_chain

import (
	"context"

	"gorm.io/gorm"

	"github.com/scroll-tech/chain-monitor/internal/logic/checker"
	"github.com/scroll-tech/chain-monitor/internal/orm"
	"github.com/scroll-tech/chain-monitor/internal/types"
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
	transactions := c.transactionOrm.GetTransactionMatch(ctx, int(layerType), start, end)
	for _, transaction := range transactions {
		checkResult := c.checker.CrossChainCheck(ctx, layerType, transaction)
		if checkResult == types.MismatchTypeUnknown {
			continue
		}

		// update status
		// send to slack
	}
}
