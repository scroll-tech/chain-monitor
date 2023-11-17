package crosschain

import (
	"context"
	"fmt"
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

// Logic is a struct for checking the l1/l2 event match from the database.
// FinalizeWithdraw value corresponds to the withdrawal value.
// FinalizeDeposit value corresponds to the deposit value.
// This is because not every deposit/withdrawal event in the system will have a finalize event,
// as users have the ability to refund deposits independently.
type Logic struct {
	messageOrm      *orm.MessageMatch
	checker         *checker.Checker
	client          *ethclient.Client
	l1MessengerAddr common.Address
	l2MessengerAddr common.Address
}

// NewLogic is a constructor for Logic.
func NewLogic(db *gorm.DB, client *ethclient.Client, l1MessengerAddr, l2MessengerAddr common.Address) *Logic {
	return &Logic{
		messageOrm:      orm.NewMessageMatch(db),
		checker:         checker.NewChecker(db),
		client:          client,
		l1MessengerAddr: l1MessengerAddr,
		l2MessengerAddr: l2MessengerAddr,
	}
}

// CheckCrossChainMessage is a method for checking cross-chain messages.
func (c *Logic) CheckCrossChainMessage(ctx context.Context, layerType types.LayerType) {
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
		if checkResult == types.MismatchTypeValid {
			messageMatchIds = append(messageMatchIds, message.ID)
			continue
		}

		// @todo send to slack
	}

	if err := c.messageOrm.UpdateCrossChainStatus(ctx, messageMatchIds, layerType, types.CrossChainStatusTypeValid); err != nil {
		log.Error("Logic.CheckCrossChainMessage UpdateCrossChainStatus failed", "error", err)
		return
	}
}

// CheckETHBalance checks the ETH balance for the given Ethereum layer (either Layer1 or Layer2).
func (c *Logic) CheckETHBalance(ctx context.Context, layerType types.LayerType) {
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
		startBlock := latestValidMessage.L1BlockNumber + 1 // no need to add startBlock
		endBlock := latestMsg.L1BlockNumber
		startBalance := new(big.Int).SetInt64(latestValidMessage.L1MessengerETHBalance.IntPart())
		endBalance, err := c.client.BalanceAt(ctx, c.l1MessengerAddr, new(big.Int).SetUint64(endBlock))
		if err != nil {
			// Would occur often when starting up.
			log.Error("get messenger balance failed", "layer types", types.Layer1, "addr", c.l1MessengerAddr, "err", err)
			return
		}
		if startBlock <= endBlock {
			ok, err := c.checkLayer1Balance(ctx, startBlock, endBlock, startBalance, endBalance)
			if err != nil {
				log.Error("checkLayer1Balance failed", "startBlock", startBlock, "endBlock", endBlock, "err", err)
				return
			}
			if ok {
				messageMatch := orm.MessageMatch{
					MessageHash:           latestMsg.MessageHash,
					L1ETHBalanceStatus:    int(types.ETHBalanceStatusTypeValid),
					L1MessengerETHBalance: decimal.NewFromBigInt(endBalance, 0),
				}
				if err := c.messageOrm.UpdateETHBalance(ctx, types.Layer1, messageMatch); err != nil {
					log.Error("insert eth balance result failed", "err", err)
					return
				}
			} else {
				if startBlock+50 <= endBlock {
					for blockHeight := startBlock; blockHeight <= endBlock; blockHeight++ {
						endBalance, err := c.client.BalanceAt(ctx, c.l1MessengerAddr, new(big.Int).SetUint64(blockHeight))
						if err != nil {
							log.Error("get messenger balance failed", "layer types", types.Layer1, "addr", c.l1MessengerAddr, "err", err)
							return
						}

						ok, err := c.checkLayer1Balance(ctx, blockHeight, blockHeight, startBalance, endBalance)
						if err != nil {
							log.Error("checkLayer1Balance failed", "startBlock", startBlock, "endBlock", blockHeight, "err", err)
							return
						}

						if !ok {
							log.Error("balance check failed", "block", blockHeight)
							// @todo: send slack.
							// At this point, we know the balance check failed for `blockHeight`.
							// You can add additional handling logic here.
							return
						}

						startBalance = endBalance
					}
				} else {
					// @todo: send slack.
					return
				}
			}
		}
	}

	if layerType == types.Layer2 {
		startBlock := latestValidMessage.L2BlockNumber + 1 // no need to add startBlock
		endBlock := latestMsg.L2BlockNumber
		startBalance := new(big.Int).SetInt64(latestValidMessage.L2MessengerETHBalance.IntPart())
		endBalance, err := c.client.BalanceAt(ctx, c.l2MessengerAddr, new(big.Int).SetUint64(endBlock))
		if err != nil {
			// Would occur often when starting up.
			log.Error("get messenger balance failed", "layer types", types.Layer2, "addr", c.l2MessengerAddr, "err", err)
			return
		}
		if startBlock <= endBlock {
			ok, err := c.checkLayer2Balance(ctx, startBlock, endBlock, startBalance, endBalance)
			if err != nil {
				log.Error("checkLayer2Balance failed", "startBlock", startBlock, "endBlock", endBlock, "err", err)
				return
			}
			if ok {
				messageMatch := orm.MessageMatch{
					MessageHash:           latestMsg.MessageHash,
					L2ETHBalanceStatus:    int(types.ETHBalanceStatusTypeValid),
					L2MessengerETHBalance: decimal.NewFromBigInt(endBalance, 0),
				}
				if err := c.messageOrm.UpdateETHBalance(ctx, types.Layer2, messageMatch); err != nil {
					log.Error("insert eth balance result failed", "err", err)
					return
				}
			} else {
				if startBlock+50 <= endBlock {
					for blockHeight := startBlock; blockHeight <= endBlock; blockHeight++ {
						endBalance, err := c.client.BalanceAt(ctx, c.l2MessengerAddr, new(big.Int).SetUint64(blockHeight))
						if err != nil {
							log.Error("get messenger balance failed", "layer types", types.Layer2, "addr", c.l2MessengerAddr, "err", err)
							return
						}

						ok, err := c.checkLayer2Balance(ctx, blockHeight, blockHeight, startBalance, endBalance)
						if err != nil {
							log.Error("checkLayer2Balance failed", "startBlock", startBlock, "endBlock", blockHeight, "err", err)
							return
						}

						if !ok {
							log.Error("balance check failed", "block", blockHeight)
							// @todo: send slack.
							// At this point, we know the balance check failed for `blockHeight`.
							// You can add additional handling logic here.
							return
						}

						startBalance = endBalance
					}
				} else {
					// @todo: send slack.
					return
				}
			}
		}
	}
}

func (c *Logic) checkLayer1Balance(ctx context.Context, startBlock, endBlock uint64, startBalance, endBalance *big.Int) (bool, error) {
	messages, err := c.messageOrm.GetMessageMatchesByBlockNumberRange(ctx, types.Layer1, startBlock, endBlock)
	if err != nil {
		return false, fmt.Errorf("GetMessageMatchesByBlockNumberRange failed: layer type %v, start height %v, end height %v, error %w", types.Layer1, startBlock, endBlock, err)
	}

	balanceDiff := big.NewInt(0)
	for _, message := range messages {
		amount, ok := new(big.Int).SetString(message.L1Amounts, 10)
		if !ok {
			return false, fmt.Errorf("invalid L1Amount value: %v", message.L1Amounts[0])
		}

		if types.EventType(message.L1EventType) == types.L1SentMessage {
			balanceDiff = new(big.Int).Add(balanceDiff, amount)
		}

		if types.EventType(message.L1EventType) == types.L1RelayedMessage {
			balanceDiff = new(big.Int).Sub(balanceDiff, amount)
		}
	}

	expectedEndBalance := new(big.Int).Add(startBalance, balanceDiff)
	return expectedEndBalance.Cmp(endBalance) == 0, nil
}

func (c *Logic) checkLayer2Balance(ctx context.Context, startBlock, endBlock uint64, startBalance, endBalance *big.Int) (bool, error) {
	messages, err := c.messageOrm.GetMessageMatchesByBlockNumberRange(ctx, types.Layer2, startBlock, endBlock)
	if err != nil {
		return false, fmt.Errorf("GetMessageMatchesByBlockNumberRange failed: layer type %v, start height %v, end height %v, error %w", types.Layer2, startBlock, endBlock, err)
	}

	balanceDiff := big.NewInt(0)
	for _, message := range messages {
		amount, ok := new(big.Int).SetString(message.L2Amounts, 10)
		if !ok {
			return false, fmt.Errorf("invalid L2Amount value: %v", message.L2Amounts[0])
		}

		if types.EventType(message.L2EventType) == types.L2SentMessage {
			balanceDiff = new(big.Int).Add(balanceDiff, amount)
		}

		if types.EventType(message.L2EventType) == types.L2RelayedMessage {
			balanceDiff = new(big.Int).Sub(balanceDiff, amount)
		}
	}

	expectedEndBalance := new(big.Int).Add(startBalance, balanceDiff)
	return expectedEndBalance.Cmp(endBalance) == 0, nil
}
