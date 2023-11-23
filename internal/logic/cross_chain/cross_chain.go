package crosschain

import (
	"context"
	"fmt"
	"math/big"

	"github.com/scroll-tech/go-ethereum/common"
	"github.com/scroll-tech/go-ethereum/ethclient"
	"github.com/scroll-tech/go-ethereum/log"
	"github.com/scroll-tech/go-ethereum/rpc"
	"github.com/shopspring/decimal"
	"gorm.io/gorm"

	"github.com/scroll-tech/chain-monitor/internal/logic/checker"
	"github.com/scroll-tech/chain-monitor/internal/logic/slack"
	"github.com/scroll-tech/chain-monitor/internal/orm"
	"github.com/scroll-tech/chain-monitor/internal/types"
)

const ethBalanceGap = 50

// LogicCrossChain is a struct for checking the l1/l2 event match from the database.
// FinalizeWithdraw value corresponds to the withdrawal value.
// FinalizeDeposit value corresponds to the deposit value.
// This is because not every deposit/withdrawal event in the system will have a finalize event,
// as users have the ability to refund deposits independently.
type LogicCrossChain struct {
	messageOrm      *orm.MessageMatch
	checker         *checker.Checker
	l1Client        *rpc.Client
	l2Client        *rpc.Client
	l1MessengerAddr common.Address
	l2MessengerAddr common.Address
}

// NewCrossChainLogic is a constructor for Logic.
func NewCrossChainLogic(db *gorm.DB, l1Client, l2Client *rpc.Client, l1MessengerAddr, l2MessengerAddr common.Address) *LogicCrossChain {
	return &LogicCrossChain{
		messageOrm:      orm.NewMessageMatch(db),
		checker:         checker.NewChecker(db),
		l1Client:        l1Client,
		l2Client:        l2Client,
		l1MessengerAddr: l1MessengerAddr,
		l2MessengerAddr: l2MessengerAddr,
	}
}

// CheckCrossChainGatewayMessage is a method for checking cross-chain messages.
func (c *LogicCrossChain) CheckCrossChainGatewayMessage(ctx context.Context, layerType types.LayerType) {
	messages, err := c.messageOrm.GetUncheckedLatestGatewayMessageMatch(ctx, 1000)
	if err != nil {
		log.Error("CheckCrossChainGatewayMessage.GetUncheckedLatestGatewayMessageMatch failed", "error", err)
		return
	}

	var messageMatchIds []int64
	for _, message := range messages {
		checkResult := c.checker.CrossChainCheck(ctx, layerType, message)
		if checkResult == types.MismatchTypeValid {
			messageMatchIds = append(messageMatchIds, message.ID)
			continue
		}
		slack.Notify(slack.MrkDwnGatewayCrossChainMessage(message, checkResult))
	}

	if err := c.messageOrm.UpdateCrossChainStatus(ctx, messageMatchIds, layerType, types.CrossChainStatusTypeValid); err != nil {
		log.Error("Logic.CheckCrossChainMessage UpdateCrossChainStatus failed", "error", err)
		return
	}
}

// CheckETHBalance checks the ETH balance for the given Ethereum layer (either Layer1 or Layer2).
func (c *LogicCrossChain) CheckETHBalance(ctx context.Context, layerType types.LayerType) {
	latestMsg, err := c.messageOrm.GetLatestDoubleLayerValidMessageMatch(ctx)
	if err != nil {
		log.Error("c.messageOrm GetLatestDoubleValidMessageMatch failed", "error", err)
		return
	}

	if latestMsg == nil {
		return
	}

	latestValidMessage, err := c.messageOrm.GetLatestValidETHBalanceMessageMatch(ctx, layerType)
	if err != nil {
		log.Error("c.messageOrm GetLatestValidETHBalanceMessageMatch failed", "layer type", layerType, "error", err)
		return
	}

	if latestValidMessage == nil {
		return
	}

	// NOTE: Suppose the `messages` get 1000 messageMatch, it contains the l1_block_number is [0-1000].
	// but the block number 1000 has more than one messageMatch. so the 1000 block number need to be ignored.
	// the bad case: the 1000 MessageMatch all have the same block number, so need another 1000 MessageMatch.
	var ignoredLastMessageMatch []orm.MessageMatch
	batchSize := 1000
	for {
		messages, err := c.messageOrm.GetUncheckedLatestETHMessageMatch(ctx, layerType, batchSize)
		if err != nil {
			log.Error("CheckETHBalance.GetUncheckedLatestETHMessageMatch failed", "error", err)
			return
		}

		for _, message := range messages {
			switch layerType {
			case types.Layer1:
				maxBlockNumber := messages[len(messages)-1].L1BlockNumber
				if message.L1BlockNumber < maxBlockNumber {
					ignoredLastMessageMatch = append(ignoredLastMessageMatch, message)
				}
			case types.Layer2:
				maxBlockNumber := messages[len(messages)-1].L2BlockNumber
				if message.L2BlockNumber < maxBlockNumber {
					ignoredLastMessageMatch = append(ignoredLastMessageMatch, message)
				}
			}
		}
		if len(ignoredLastMessageMatch) > 0 {
			break
		}
		batchSize += 1000
	}

	c.checkETH(ctx, layerType, latestMsg, latestValidMessage, ignoredLastMessageMatch)
}

func (c *LogicCrossChain) checkETH(ctx context.Context, layer types.LayerType, latestMsg, latestValidMessage *orm.MessageMatch, messages []orm.MessageMatch) {
	var startBlockNumber, endBlockNumber, latestBlockNumber uint64
	var messengerAddr common.Address
	var client *rpc.Client
	if layer == types.Layer1 {
		startBlockNumber = messages[0].L1BlockNumber
		endBlockNumber = messages[len(messages)-1].L1BlockNumber
		latestBlockNumber = latestMsg.L1BlockNumber
		messengerAddr = c.l1MessengerAddr
		client = c.l1Client
	} else {
		startBlockNumber = messages[0].L2BlockNumber
		endBlockNumber = messages[len(messages)-1].L2BlockNumber
		latestBlockNumber = latestMsg.L2BlockNumber
		messengerAddr = c.l2MessengerAddr
		client = c.l2Client
	}

	// because balanceAt can't get the too early block balance, so only can compute the locally l1 messenger balance and
	// update the l1_messenger_eth_balance/l2_messenger_eth_balance
	if endBlockNumber+ethBalanceGap < latestBlockNumber {
		c.computeBlockBalance(ctx, layer, messages, latestValidMessage)
		return
	}

	endBalance, err := ethclient.NewClient(client).BalanceAt(ctx, messengerAddr, new(big.Int).SetUint64(endBlockNumber))
	if err != nil {
		log.Error("get messenger balance failed", "layer types", layer, "addr", messengerAddr, "err", err)
		return
	}

	var startBalance *big.Int
	if layer == types.Layer1 {
		startBalance = latestValidMessage.L1MessengerETHBalance.BigInt()
	} else {
		startBalance = latestValidMessage.L2MessengerETHBalance.BigInt()
	}

	ok, expectedEndBalance, actualBalance, err := c.checkBalance(layer, startBalance, endBalance, messages)
	if err != nil {
		log.Error("checkLayer1Balance failed", "startBlock", startBlockNumber, "endBlock", endBlockNumber, "expectedEndBalance", expectedEndBalance, "actualBalance", actualBalance, "err", err)
		return
	}

	if !ok {
		c.checkBlockBalanceOneByOne(ctx, ethclient.NewClient(client), messengerAddr, layer, messages)
		return
	}

	// get all the eth status valid, and update the eth balance status and eth balance
	c.computeBlockBalance(ctx, layer, messages, latestValidMessage)
}

func (c *LogicCrossChain) checkBlockBalanceOneByOne(ctx context.Context, client *ethclient.Client, messengerAddr common.Address, layer types.LayerType, messages []orm.MessageMatch) {
	var startBalance *big.Int
	var startIndex int
	for idx, message := range messages {
		var blockNumber uint64
		if layer == types.Layer1 {
			blockNumber = message.L1BlockNumber
		} else {
			blockNumber = message.L2BlockNumber
		}

		tmpBalance, err := client.BalanceAt(ctx, messengerAddr, new(big.Int).SetUint64(blockNumber))
		if err != nil {
			continue
		}

		startBalance = tmpBalance
		startIndex = idx
		break
	}

	for i := startIndex + 1; i < len(messages); i++ {
		var blockNumber uint64
		if layer == types.Layer1 {
			blockNumber = messages[i].L1BlockNumber
		} else {
			blockNumber = messages[i].L2BlockNumber
		}

		if layer == types.Layer1 && (i+1 != len(messages) && (messages[i].L1BlockNumber == messages[i+1].L1BlockNumber)) {
			continue
		}

		if layer == types.Layer2 && (i+1 != len(messages) && (messages[i].L2BlockNumber == messages[i+1].L2BlockNumber)) {
			continue
		}

		endBalance, err := client.BalanceAt(ctx, messengerAddr, new(big.Int).SetUint64(blockNumber))
		if err != nil {
			continue
		}

		ok, expectedEndBalance, actualBalance, err := c.checkBalance(layer, startBalance, endBalance, messages[startIndex:i+1])
		if !ok || err != nil {
			log.Error("balance check failed", "block", blockNumber)
			slack.MrkDwnETHGatewayMessage(&messages[i], expectedEndBalance, actualBalance)
			continue
		}
	}
}

func (c *LogicCrossChain) checkBalance(layer types.LayerType, startBalance, endBalance *big.Int, messages []orm.MessageMatch) (bool, *big.Int, *big.Int, error) {
	balanceDiff := big.NewInt(0)
	for _, message := range messages {
		var amount *big.Int
		var ok bool
		if layer == types.Layer1 {
			amount, ok = new(big.Int).SetString(message.L1Amounts, 10)
			if !ok {
				return false, nil, nil, fmt.Errorf("invalid L1Amount value: %v", message.L1Amounts)
			}
		} else {
			amount, ok = new(big.Int).SetString(message.L2Amounts, 10)
			if !ok {
				return false, nil, nil, fmt.Errorf("invalid L2Amounts value: %v", message.L2Amounts)
			}
		}

		if types.EventType(message.L1EventType) == types.L1SentMessage || types.EventType(message.L2EventType) == types.L2SentMessage {
			balanceDiff = new(big.Int).Add(balanceDiff, amount)
		}

		if types.EventType(message.L1EventType) == types.L1RelayedMessage || types.EventType(message.L2EventType) == types.L2RelayedMessage {
			balanceDiff = new(big.Int).Sub(balanceDiff, amount)
		}
	}

	expectedEndBalance := new(big.Int).Add(startBalance, balanceDiff)
	if expectedEndBalance.Cmp(endBalance) == 0 {
		return true, expectedEndBalance, endBalance, nil
	}

	return false, expectedEndBalance, endBalance, nil
}

func (c *LogicCrossChain) computeBlockBalance(ctx context.Context, layer types.LayerType, messages []orm.MessageMatch, latestValidMessage *orm.MessageMatch) {
	var messengerETHBalance *big.Int
	if latestValidMessage != nil {
		if layer == types.Layer1 {
			messengerETHBalance = latestValidMessage.L1MessengerETHBalance.BigInt()
		} else {
			messengerETHBalance = latestValidMessage.L2MessengerETHBalance.BigInt()
		}
	}

	blockNumberAmountMap := make(map[uint64]*big.Int)
	for _, message := range messages {
		if layer == types.Layer1 {
			if _, ok := blockNumberAmountMap[message.L1BlockNumber]; !ok {
				blockNumberAmountMap[message.L1BlockNumber] = new(big.Int)
			}

			amount, ok := new(big.Int).SetString(message.L1Amounts, 10)
			if !ok {
				log.Error("invalid L1Amount value", "l1 amounts", message.L1Amounts)
				return
			}

			if types.EventType(message.L1EventType) == types.L1SentMessage {
				blockNumberAmountMap[message.L1BlockNumber] = new(big.Int).Add(blockNumberAmountMap[message.L1BlockNumber], amount)
			}

			if types.EventType(message.L1EventType) == types.L1RelayedMessage {
				blockNumberAmountMap[message.L1BlockNumber] = new(big.Int).Sub(blockNumberAmountMap[message.L1BlockNumber], amount)
			}
		}

		if layer == types.Layer2 {
			if _, ok := blockNumberAmountMap[message.L2BlockNumber]; !ok {
				blockNumberAmountMap[message.L2BlockNumber] = new(big.Int)
			}

			amount, ok := new(big.Int).SetString(message.L2Amounts, 10)
			if !ok {
				log.Error("invalid L2Amount value", "l2 amounts", message.L2Amounts)
				return
			}

			if types.EventType(message.L2EventType) == types.L2SentMessage {
				blockNumberAmountMap[message.L2BlockNumber] = new(big.Int).Add(blockNumberAmountMap[message.L2BlockNumber], amount)
			}

			if types.EventType(message.L2EventType) == types.L2RelayedMessage {
				blockNumberAmountMap[message.L2BlockNumber] = new(big.Int).Sub(blockNumberAmountMap[message.L2BlockNumber], amount)
			}
		}
	}

	for k, v := range messages {
		var ethBalance *big.Int
		var blockETHBalance *big.Int
		var lastBlockETHBalance *big.Int

		if layer == types.Layer1 {
			blockETHBalance = blockNumberAmountMap[v.L1BlockNumber]
			if k != 0 {
				lastBlockETHBalance = messages[k-1].L1MessengerETHBalance.BigInt()
			} else {
				lastBlockETHBalance = messengerETHBalance
			}
		} else {
			blockETHBalance = blockNumberAmountMap[v.L2BlockNumber]
			if k != 0 {
				lastBlockETHBalance = messages[k-1].L2MessengerETHBalance.BigInt()
			} else {
				lastBlockETHBalance = messengerETHBalance
			}
		}

		ethBalance = new(big.Int).Add(lastBlockETHBalance, blockETHBalance)

		// update the db
		mm := orm.MessageMatch{
			MessageHash: v.MessageHash,
			CheckStatus: int(types.CheckStatusChecked),
		}

		if layer == types.Layer1 {
			mm.L1MessengerETHBalance = decimal.NewFromBigInt(ethBalance, 0)
			mm.L1ETHBalanceStatus = int(types.ETHBalanceStatusTypeValid)
		} else {
			mm.L2MessengerETHBalance = decimal.NewFromBigInt(ethBalance, 0)
			mm.L2ETHBalanceStatus = int(types.ETHBalanceStatusTypeValid)
		}
		if err := c.messageOrm.UpdateETHBalance(ctx, layer, mm); err != nil {
			log.Error("computeOverageBlockBalance.UpdateETHBalance failed", "layer", layer, "message match:%v", mm, "error", err)
		}
	}
}
