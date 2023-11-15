package cross_chain

import (
	"context"
	"math/big"

	"github.com/scroll-tech/go-ethereum/common"
	"github.com/scroll-tech/go-ethereum/ethclient"
	"github.com/scroll-tech/go-ethereum/log"
	"github.com/shopspring/decimal"
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
	messageOrm      *orm.MessageMatch
	checker         *checker.Checker
	client          *ethclient.Client
	l1MessengerAddr common.Address
	l2MessengerAddr common.Address
}

func NewCrossChainLogic(db *gorm.DB, client *ethclient.Client, l1MessengerAddr, l2MessengerAddr common.Address) *CrossChainLogic {
	return &CrossChainLogic{
		messageOrm:      orm.NewMessageMatch(db),
		checker:         checker.NewChecker(db),
		client:          client,
		l1MessengerAddr: l1MessengerAddr,
		l2MessengerAddr: l2MessengerAddr,
	}
}

func (c *CrossChainLogic) CheckCrossChainMessage(ctx context.Context, layerType types.LayerType) {
	latestValidMsg, err := c.messageOrm.GetLatestValidCrossChainMessageMatch(ctx, layerType)
	if err != nil {
		log.Error("c.messageOrm GetLatestValidCrossChainMessageMatch failed", "layer", layerType, "error", err)
		return
	}

	latestMsg, err := c.messageOrm.GetLatestDoubleValidMessageMatch(ctx)
	if err != nil {
		log.Error("c.messageOrm GetLatestDoubleValidMessageMatch failed", "error", err)
		return
	}

	var startHeight, endHeight uint64
	if layerType == types.Layer1 {
		startHeight = latestValidMsg.L1BlockNumber
		endHeight = latestMsg.L1BlockNumber
	} else {
		startHeight = latestValidMsg.L2BlockNumber
		endHeight = latestMsg.L2BlockNumber
	}

	messages, err := c.messageOrm.GetMessageMatchesByBlockNumberRange(ctx, layerType, startHeight, endHeight)
	if err != nil {
		log.Error("GetMessageMatchesByBlockNumberRange failed", "layer type", layerType, "start height", startHeight, "end height", endHeight, "error", err)
		return
	}

	var messageMatchIds []int64
	for _, message := range messages {
		checkResult := c.checker.CrossChainCheck(ctx, layerType, message)
		if checkResult == types.MismatchTypeOk {
			messageMatchIds = append(messageMatchIds, message.ID)
			continue
		}

		// @todo send to slack
	}

	if err := c.messageOrm.UpdateCrossChainStatus(ctx, messageMatchIds, layerType, types.CrossChainStatusTypeValid); err != nil {
		log.Error("CrossChainLogic.CheckCrossChainMessage UpdateCrossChainStatus failed", "error", err)
		return
	}
}

func (c *CrossChainLogic) CheckETHBalance(ctx context.Context, layerType types.LayerType) {
	latestMsg, err := c.messageOrm.GetLatestDoubleValidMessageMatch(ctx)
	if err != nil {
		log.Error("c.messageOrm GetLatestDoubleValidMessageMatch failed", "error", err)
		return
	}

	latestValidMessage, err := c.messageOrm.GetLatesValidETHBalanceMessageMatch(ctx, layerType)
	if err != nil {
		log.Error("c.messageOrm GetLatesValidETHBalanceMessageMatch failed", "layer type", layerType, "error", err)
		return
	}

	if layerType == types.Layer1 {
		startHeight := latestValidMessage.L1BlockNumber
		endHeight := latestMsg.L1BlockNumber
		if startHeight >= endHeight {
			log.Info("no ready block to check", "layer types", layerType, "start height", startHeight, "end height", endHeight)
			return
		}
		startBalance := new(big.Int).SetInt64(latestValidMessage.L1MessengerETHBalance.IntPart())
		endBalance, err := c.client.BalanceAt(ctx, c.l1MessengerAddr, new(big.Int).SetUint64(endHeight))
		if err != nil {
			log.Error("get messenger balance failed", "layer types", layerType, "addr", c.l1MessengerAddr, "err", err)
			return
		}
		messages, err := c.messageOrm.GetMessageMatchesByBlockNumberRange(ctx, layerType, startHeight+1, endHeight)
		if err != nil {
			log.Error("GetMessageMatchesByBlockNumberRange failed", "layer type", layerType, "start height", startHeight+1, "end height", endHeight, "error", err)
			return
		}

		balanceDiff := big.NewInt(0)
		for _, message := range messages {
			if types.EventType(message.L1EventType) == types.L1SentMessage {
				balanceDiff = new(big.Int).Add(balanceDiff, new(big.Int).SetInt64(message.L1Amount.IntPart()))
			} else {
				balanceDiff = new(big.Int).Sub(balanceDiff, new(big.Int).SetInt64(message.L2Amount.IntPart()))
			}
		}

		expectedEndBalance := new(big.Int).Add(startBalance, balanceDiff)
		var messageMatch orm.MessageMatch
		if expectedEndBalance.Cmp(endBalance) == 0 {
			messageMatch = orm.MessageMatch{
				MessageHash:           latestMsg.MessageHash,
				L1ETHBalanceStatus:    int(types.ETHBalanceStatusTypeValid),
				L1MessengerETHBalance: decimal.NewFromBigInt(endBalance, 0),
			}
		} else {
			// @todo: send slack.
			log.Error("Expected end balance does not match actual end balance", "start height", startHeight, "end height", endHeight,
				"start balance", startBalance, "end balance", endBalance, "balance diff", balanceDiff)
			messageMatch = orm.MessageMatch{
				MessageHash:           latestMsg.MessageHash,
				L1ETHBalanceStatus:    int(types.ETHBalanceStatusTypeInvalid),
				L1MessengerETHBalance: decimal.NewFromBigInt(endBalance, 0),
			}
		}
		messageMatches := []orm.MessageMatch{messageMatch}
		effectRows, err := c.messageOrm.InsertOrUpdate(ctx, []orm.MessageMatch{messageMatch})
		if err != nil || effectRows != len(messageMatches) {
			log.Error("insert eth balance result failed", "err", err)
		}
	} else {
		startHeight := latestValidMessage.L2BlockNumber
		endHeight := latestMsg.L2BlockNumber
		if startHeight >= endHeight {
			log.Info("no ready block to check", "layer types", layerType, "start height", startHeight, "end height", endHeight)
			return
		}
		startBalance := new(big.Int).SetInt64(latestValidMessage.L2MessengerETHBalance.IntPart())
		endBalance, err := c.client.BalanceAt(ctx, c.l2MessengerAddr, new(big.Int).SetUint64(endHeight))
		if err != nil {
			log.Error("get messenger balance failed", "layer types", layerType, "addr", c.l2MessengerAddr, "err", err)
			return
		}
		messages, err := c.messageOrm.GetMessageMatchesByBlockNumberRange(ctx, layerType, startHeight+1, endHeight)
		if err != nil {
			log.Error("GetMessageMatchesByBlockNumberRange failed", "layer type", layerType, "start height", startHeight+1, "end height", endHeight, "error", err)
			return
		}

		balanceDiff := big.NewInt(0)
		for _, message := range messages {
			if types.EventType(message.L2EventType) == types.L2SentMessage {
				balanceDiff = new(big.Int).Add(balanceDiff, new(big.Int).SetInt64(message.L2Amount.IntPart()))
			} else {
				balanceDiff = new(big.Int).Sub(balanceDiff, new(big.Int).SetInt64(message.L1Amount.IntPart()))
			}
		}

		expectedEndBalance := new(big.Int).Add(startBalance, balanceDiff)
		var messageMatch orm.MessageMatch
		if expectedEndBalance.Cmp(endBalance) == 0 {
			messageMatch = orm.MessageMatch{
				MessageHash:           latestMsg.MessageHash,
				L2ETHBalanceStatus:    int(types.ETHBalanceStatusTypeValid),
				L2MessengerETHBalance: decimal.NewFromBigInt(endBalance, 0),
			}
		} else {
			// @todo: send slack.
			log.Error("Expected end balance does not match actual end balance", "start height", startHeight, "end height", endHeight,
				"start balance", startBalance, "end balance", endBalance, "balance diff", balanceDiff)
			messageMatch = orm.MessageMatch{
				MessageHash:           latestMsg.MessageHash,
				L2ETHBalanceStatus:    int(types.ETHBalanceStatusTypeInvalid),
				L2MessengerETHBalance: decimal.NewFromBigInt(endBalance, 0),
			}
		}
		messageMatches := []orm.MessageMatch{messageMatch}
		effectRows, err := c.messageOrm.InsertOrUpdate(ctx, []orm.MessageMatch{messageMatch})
		if err != nil || effectRows != len(messageMatches) {
			log.Error("insert eth balance result failed", "err", err)
		}
	}
}
