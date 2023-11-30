package crosschain

import (
	"context"
	"fmt"
	"math/big"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/scroll-tech/go-ethereum/common"
	"github.com/scroll-tech/go-ethereum/ethclient"
	"github.com/scroll-tech/go-ethereum/log"
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
	db              *gorm.DB
	messageOrm      *orm.MessageMatch
	checker         *checker.Checker
	l1Client        *ethclient.Client
	l2Client        *ethclient.Client
	l1MessengerAddr common.Address
	l2MessengerAddr common.Address

	crossChainGatewayCheckID *prometheus.GaugeVec
	crossChainETHCheckID     *prometheus.GaugeVec
}

// NewCrossChainLogic is a constructor for Logic.
func NewCrossChainLogic(db *gorm.DB, l1Client, l2Client *ethclient.Client, l1MessengerAddr, l2MessengerAddr common.Address) *LogicCrossChain {
	return &LogicCrossChain{
		db:              db,
		messageOrm:      orm.NewMessageMatch(db),
		checker:         checker.NewChecker(db),
		l1Client:        l1Client,
		l2Client:        l2Client,
		l1MessengerAddr: l1MessengerAddr,
		l2MessengerAddr: l2MessengerAddr,

		crossChainGatewayCheckID: promauto.With(prometheus.DefaultRegisterer).NewGaugeVec(prometheus.GaugeOpts{
			Name: "cross_chain_checked_gateway_event_database_id",
			Help: "the database id of cross chain gateway checked",
		}, []string{"layer"}),

		crossChainETHCheckID: promauto.With(prometheus.DefaultRegisterer).NewGaugeVec(prometheus.GaugeOpts{
			Name: "cross_chain_checked_eth_database_id",
			Help: "the database id of cross chain eth checked",
		}, []string{"layer"}),
	}
}

// CheckCrossChainGatewayMessage is a method for checking cross-chain messages.
func (c *LogicCrossChain) CheckCrossChainGatewayMessage(ctx context.Context, layerType types.LayerType) {
	messages, err := c.messageOrm.GetUncheckedAndDoubleLayerValidGatewayMessageMatches(ctx, layerType, 1000)
	if err != nil {
		log.Error("CheckCrossChainGatewayMessage.GetUncheckedAndDoubleLayerValidGatewayMessageMatchs failed", "error", err)
		return
	}

	if len(messages) == 0 {
		time.Sleep(time.Second)
		return
	}

	log.Info("checking cross chain gateway messages", "number of messages", len(messages))

	var messageMatchIds []int64
	for _, message := range messages {
		c.crossChainGatewayCheckID.WithLabelValues(layerType.String()).Set(float64(message.ID))
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
	log.Info("CheckETHBalance started", "layer type", layerType)

	latestBlockNumber, err := c.getLatestBlockNumber(ctx, layerType)
	if err != nil {
		log.Error("get latest block number from geth node failed", "layer", layerType, "error", err)
		return
	}

	startBalance, err := c.messageOrm.GetETHCheckStartBlockNumberAndBalance(ctx, layerType)
	if err != nil {
		log.Error("c.messageOrm GetETHCheckStartBlockNumberAndBalance failed", "layer type", layerType, "error", err)
		return
	}

	if layerType == types.Layer2 && startBalance.Cmp(new(big.Int)) == 0 {
		var err error
		startBalance, err = c.l2Client.BalanceAt(ctx, c.l2MessengerAddr, new(big.Int).SetUint64(0))
		if err != nil {
			log.Error("get messenger balance failed", "layer types", layerType, "err", err)
			return
		}
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

		if len(messages) == 0 {
			return
		}

		for _, message := range messages {
			if message.ETHAmountStatus == int(types.ETHAmountStatusTypeUnset) {
				break
			}
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

	c.checkETH(ctx, layerType, latestBlockNumber, startBalance, ignoredLastMessageMatch)
	log.Info("CheckETHBalance completed", "layer type", layerType)
}

func (c *LogicCrossChain) checkETH(ctx context.Context, layer types.LayerType, latestBlockNumber uint64, startBalance *big.Int, messages []orm.MessageMatch) {
	var startBlockNumber, endBlockNumber uint64
	var messengerAddr common.Address
	var client *ethclient.Client
	if layer == types.Layer1 {
		startBlockNumber = messages[0].L1BlockNumber
		endBlockNumber = messages[len(messages)-1].L1BlockNumber
		messengerAddr = c.l1MessengerAddr
		client = c.l1Client
	} else {
		startBlockNumber = messages[0].L2BlockNumber
		endBlockNumber = messages[len(messages)-1].L2BlockNumber
		messengerAddr = c.l2MessengerAddr
		client = c.l2Client
	}

	log.Info("checking eth balance", "start", startBlockNumber, "end", endBlockNumber, "latest", latestBlockNumber)

	// because balanceAt can't get the too early block balance, so only can compute the locally l1 messenger balance and
	// update the l1_messenger_eth_balance/l2_messenger_eth_balance
	if layer == types.Layer1 && endBlockNumber+ethBalanceGap < latestBlockNumber {
		c.computeBlockBalance(ctx, layer, messages, startBalance)
		return
	}

	endBalance, err := client.BalanceAt(ctx, messengerAddr, new(big.Int).SetUint64(endBlockNumber))
	if err != nil {
		log.Error("get messenger balance failed", "layer types", layer, "addr", messengerAddr, "end", endBlockNumber, "err", err)
		return
	}

	ok, expectedEndBalance, actualBalance, err := c.checkBalance(layer, startBalance, endBalance, messages)
	if err != nil {
		log.Error("checkLayer1Balance failed", "startBlock", startBlockNumber, "endBlock", endBlockNumber, "expectedEndBalance", expectedEndBalance, "actualBalance", actualBalance, "err", err)
		return
	}

	if !ok {
		c.checkBlockBalanceOneByOne(ctx, client, messengerAddr, layer, messages)
		return
	}

	// get all the eth status valid, and update the eth balance status and eth balance
	c.computeBlockBalance(ctx, layer, messages, startBalance)
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
			log.Error("get balance failed", "err", err)
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
			log.Error("balance check failed", "block", blockNumber, "expectedEndBalance", expectedEndBalance.String(), "actualBalance", actualBalance.String())
			slack.MrkDwnETHGatewayMessage(&messages[i], expectedEndBalance, actualBalance)
			continue
		}
	}
}

func (c *LogicCrossChain) checkBalance(layer types.LayerType, startBalance, endBalance *big.Int, messages []orm.MessageMatch) (bool, *big.Int, *big.Int, error) {
	balanceDiff := big.NewInt(0)
	for _, message := range messages {
		c.crossChainETHCheckID.WithLabelValues(layer.String()).Set(float64(message.ID))

		var amount *big.Int
		var ok bool
		amount, ok = new(big.Int).SetString(message.ETHAmount, 10)
		if !ok {
			return false, nil, nil, fmt.Errorf("invalid ETHAmount value: %v, layer: %v", message.ETHAmount, layer)
		}

		if layer == types.Layer1 {
			if types.EventType(message.L1EventType) == types.L1SentMessage || types.EventType(message.L1EventType) == types.L1DepositERC20 {
				balanceDiff = new(big.Int).Add(balanceDiff, amount)
			}

			if types.EventType(message.L1EventType) == types.L1RelayedMessage || types.EventType(message.L1EventType) == types.L1FinalizeWithdrawERC20 {
				balanceDiff = new(big.Int).Sub(balanceDiff, amount)
			}
		}

		if layer == types.Layer2 {
			if types.EventType(message.L2EventType) == types.L2SentMessage || types.EventType(message.L2EventType) == types.L2WithdrawERC20 {
				balanceDiff = new(big.Int).Add(balanceDiff, amount)
			}

			if types.EventType(message.L2EventType) == types.L2RelayedMessage || types.EventType(message.L2EventType) == types.L2FinalizeDepositERC20 {
				balanceDiff = new(big.Int).Sub(balanceDiff, amount)
			}
		}
	}

	expectedEndBalance := new(big.Int).Add(startBalance, balanceDiff)
	if expectedEndBalance.Cmp(endBalance) == 0 {
		return true, expectedEndBalance, endBalance, nil
	}

	log.Info("balance check failed", "expectedEndBalance", expectedEndBalance.String(), "actualBalance", endBalance.String())
	return false, expectedEndBalance, endBalance, nil
}

func (c *LogicCrossChain) computeBlockBalance(ctx context.Context, layer types.LayerType, messages []orm.MessageMatch, messengerETHBalance *big.Int) {
	blockNumberAmountMap := make(map[uint64]*big.Int)
	for _, message := range messages {
		if layer == types.Layer1 {
			if _, ok := blockNumberAmountMap[message.L1BlockNumber]; !ok {
				blockNumberAmountMap[message.L1BlockNumber] = new(big.Int)
			}

			amount, ok := new(big.Int).SetString(message.ETHAmount, 10)
			if !ok {
				log.Error("invalid L1 ETH Amount value", "amount", message.ETHAmount)
				return
			}

			if types.EventType(message.L1EventType) == types.L1SentMessage || types.EventType(message.L1EventType) == types.L1DepositERC20 {
				blockNumberAmountMap[message.L1BlockNumber] = new(big.Int).Add(blockNumberAmountMap[message.L1BlockNumber], amount)
			}

			if types.EventType(message.L1EventType) == types.L1RelayedMessage || types.EventType(message.L1EventType) == types.L1FinalizeWithdrawERC20 {
				blockNumberAmountMap[message.L1BlockNumber] = new(big.Int).Sub(blockNumberAmountMap[message.L1BlockNumber], amount)
			}
		}

		if layer == types.Layer2 {
			if _, ok := blockNumberAmountMap[message.L2BlockNumber]; !ok {
				blockNumberAmountMap[message.L2BlockNumber] = new(big.Int)
			}

			amount, ok := new(big.Int).SetString(message.ETHAmount, 10)
			if !ok {
				log.Error("invalid L2 ETH Amount value", "amount", message.ETHAmount)
				return
			}

			if types.EventType(message.L2EventType) == types.L2SentMessage || types.EventType(message.L2EventType) == types.L2WithdrawERC20 {
				blockNumberAmountMap[message.L2BlockNumber] = new(big.Int).Add(blockNumberAmountMap[message.L2BlockNumber], amount)
			}

			if types.EventType(message.L2EventType) == types.L2RelayedMessage || types.EventType(message.L2EventType) == types.L2FinalizeDepositERC20 {
				blockNumberAmountMap[message.L2BlockNumber] = new(big.Int).Sub(blockNumberAmountMap[message.L2BlockNumber], amount)
			}
		}
	}

	var updateETHMessageMatches []orm.MessageMatch
	lastBlockBalance := new(big.Int).Set(messengerETHBalance)
	lastBlockNumber := uint64(0)
	for _, v := range messages {
		blockNumber := v.L1BlockNumber
		if layer == types.Layer2 {
			blockNumber = v.L2BlockNumber
		}

		if blockNumber != lastBlockNumber {
			lastBlockBalance.Add(lastBlockBalance, blockNumberAmountMap[blockNumber])
			lastBlockNumber = blockNumber
		}

		// update the db
		mm := orm.MessageMatch{ID: v.ID}
		if layer == types.Layer1 {
			mm.L1MessengerETHBalance = decimal.NewFromBigInt(lastBlockBalance, 0)
			mm.L1ETHBalanceStatus = int(types.ETHBalanceStatusTypeValid)
		} else {
			mm.L2MessengerETHBalance = decimal.NewFromBigInt(lastBlockBalance, 0)
			mm.L2ETHBalanceStatus = int(types.ETHBalanceStatusTypeValid)
		}
		updateETHMessageMatches = append(updateETHMessageMatches, mm)
	}

	err := c.db.Transaction(func(tx *gorm.DB) error {
		for _, updateEthMessageMatch := range updateETHMessageMatches {
			if err := c.messageOrm.UpdateETHBalance(ctx, layer, updateEthMessageMatch, tx); err != nil {
				log.Error("computeOverageBlockBalance.UpdateETHBalance failed", "layer", layer, "message match:%v", updateEthMessageMatch, "error", err)
				return err
			}
		}
		return nil
	})
	if err != nil {
		log.Error("computeOverageBlockBalance.UpdateETHBalance failed", "layer", layer, "error", err)
	}
}

func (c *LogicCrossChain) getLatestBlockNumber(ctx context.Context, layerType types.LayerType) (uint64, error) {
	switch layerType {
	case types.Layer1:
		latestHeader, err := c.l1Client.HeaderByNumber(ctx, nil)
		if err != nil {
			log.Error("Failed to get latest header from Layer1 client", "error", err)
			return 0, err
		}
		return latestHeader.Number.Uint64(), nil

	case types.Layer2:
		latestHeader, err := c.l2Client.HeaderByNumber(ctx, nil)
		if err != nil {
			log.Error("Failed to get latest header from Layer2 client", "error", err)
			return 0, err
		}
		return latestHeader.Number.Uint64(), nil

	default:
		log.Error("Invalid layerType", "layerType", layerType)
		return 0, fmt.Errorf("invalid layerType: %v", layerType)
	}
}
