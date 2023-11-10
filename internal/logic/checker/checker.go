package checker

import (
	"context"
	"fmt"

	"github.com/scroll-tech/go-ethereum/common"
	"github.com/shopspring/decimal"
	"gorm.io/gorm"

	"github.com/scroll-tech/chain-monitor/internal/logic/events"
	"github.com/scroll-tech/chain-monitor/internal/orm"
	"github.com/scroll-tech/chain-monitor/internal/types"
	"github.com/scroll-tech/chain-monitor/internal/utils/msgproof"
)

type Checker struct {
	messageMatchOrm   *orm.MessageMatch
	transferMatcher   *TransferEventMatcher
	crossChainMatcher *CrossEventMatcher
}

func NewChecker(db *gorm.DB) *Checker {
	return &Checker{
		messageMatchOrm:   orm.NewMessageMatch(db),
		transferMatcher:   NewTransferEventMatcher(),
		crossChainMatcher: NewCrossEventMatcher(),
	}
}

func (c *Checker) CrossChainCheck(_ context.Context, layer types.LayerType, messageMatch orm.MessageMatch) types.MismatchType {
	if layer == types.Layer1 {
		if !c.crossChainMatcher.L1EventMatchL2(messageMatch) {
			return types.MismatchTypeL2EventNotExist
		}
	}

	if layer == types.Layer2 {
		if !c.crossChainMatcher.L2EventMatchL1(messageMatch) {
			return types.MismatchTypeL1EventNotExist
		}
	}

	if !c.crossChainMatcher.CrossChainAmountMatch(messageMatch) {
		return types.MismatchTypeAmount
	}

	if !c.crossChainMatcher.EventTypeMatch(messageMatch) {
		return types.MismatchTypeCrossChainTypeNotMatch
	}

	return types.MismatchTypeOk
}

func (c *Checker) GatewayCheck(ctx context.Context, eventCategory types.TxEventCategory, eventDataList []events.EventUnmarshaler, transferDataList []events.EventUnmarshaler) error {
	switch eventCategory {
	case types.ERC20EventCategory:
		return c.erc20EventUnmarshaler(ctx, eventDataList, transferDataList)
	}
	return nil
}

func (c *Checker) CheckL2WithdrawRoots(ctx context.Context, startBlockNumber, endBlockNumber uint64, sentMessageEvents map[uint64][]events.SentMessageEvent, withdrawRoots map[uint64]common.Hash) error {
	withdrawTrie := msgproof.NewWithdrawTrie()
	if startBlockNumber > 1 { // ignore genesis block itself and block whose parent is genesis block.
		msg, err := c.messageMatchOrm.GetMessageMatchByL2BlockNumber(ctx, startBlockNumber-1)
		if err != nil {
			return err
		}
		withdrawTrie.Initialize(msg.MessageNonce, common.HexToHash(msg.MessageHash), msg.MessageProof)
	}

	var messageMatches []orm.MessageMatch
	lastWithdrawRoot := withdrawTrie.MessageRoot()
	for blockNum := startBlockNumber; blockNum <= endBlockNumber; blockNum++ {
		events := sentMessageEvents[blockNum]
		eventHashes := make([]common.Hash, len(events))
		for i, event := range events {
			eventHashes[i] = event.MessageHash
		}
		proofs := withdrawTrie.AppendMessages(eventHashes)
		lastWithdrawRoot = withdrawTrie.MessageRoot()
		if lastWithdrawRoot != withdrawRoots[blockNum] {
			// @todo: send slack message.
			return fmt.Errorf("withdraw root mismatch in %v, got: %v, expected %v", blockNum, lastWithdrawRoot, withdrawRoots[blockNum])
		}
		// current block has SentMessage events.
		numEvents := len(eventHashes)
		if numEvents > 0 {
			// only update the last message of each block (which contains at least one SentMessage event).
			messageMatches = append(messageMatches, orm.MessageMatch{
				MessageHash:  eventHashes[numEvents-1].Hex(),
				MessageProof: proofs[numEvents-1],
				MessageNonce: withdrawTrie.NextMessageNonce,
			})
		}
	}

	effectRows, err := c.messageMatchOrm.InsertOrUpdate(ctx, messageMatches)
	if err != nil || effectRows != len(messageMatches) {
		return fmt.Errorf("erc20EventUnmarshaler orm insert failed, err: %w", err)
	}
	return nil
}

func (c *Checker) erc20EventUnmarshaler(ctx context.Context, eventDataList []events.EventUnmarshaler, transferDataList []events.EventUnmarshaler) error {
	var messageMatches []orm.MessageMatch
	type EventKey struct {
		BlockNumber uint64
		LogIndex    uint
	}
	messageHashes := make(map[EventKey]common.Hash)

	var gatewayEvents []events.ERC20GatewayEventUnmarshaler
	for _, eventData := range eventDataList {
		erc20EventUnmarshaler, ok := eventData.(*events.ERC20GatewayEventUnmarshaler)
		if !ok {
			return fmt.Errorf("eventData is not of type *events.ERC20GatewayEventUnmarshaler")
		}
		switch erc20EventUnmarshaler.Type {
		case types.L1SentMessage, types.L2SentMessage, types.L1RelayedMessage, types.L2RelayedMessage:
			key := EventKey{BlockNumber: erc20EventUnmarshaler.Number, LogIndex: erc20EventUnmarshaler.Index}
			messageHashes[key] = erc20EventUnmarshaler.MessageHash
		case types.L1DepositERC20, types.L1FinalizeWithdrawERC20, types.L1RefundERC20, types.L2WithdrawERC20, types.L2FinalizeDepositERC20:
			gatewayEvents = append(gatewayEvents, *erc20EventUnmarshaler)
		default:
			return fmt.Errorf("unknown erc20 event type: %v", erc20EventUnmarshaler.Type)
		}
	}

	for _, eventData := range eventDataList {
		erc20EventUnmarshaler := eventData.(*events.ERC20GatewayEventUnmarshaler)

		var tmpMessageMatch orm.MessageMatch
		switch erc20EventUnmarshaler.Type {
		case types.L1DepositERC20:
			key := EventKey{BlockNumber: erc20EventUnmarshaler.Number, LogIndex: erc20EventUnmarshaler.Index + 1}
			messageHash, exists := messageHashes[key]
			if !exists {
				return fmt.Errorf("message hash does not exist for key %d", key)
			}
			tmpMessageMatch = orm.MessageMatch{
				MessageHash:   messageHash.Hex(),
				TokenType:     int(types.TokenTypeERC20),
				L1EventType:   int(erc20EventUnmarshaler.Type),
				L1BlockNumber: erc20EventUnmarshaler.Number,
				L1TxHash:      erc20EventUnmarshaler.TxHash.Hex(),
				L1Amount:      decimal.NewFromBigInt(erc20EventUnmarshaler.Amount, 0),
			}
			messageMatches = append(messageMatches, tmpMessageMatch)
		case types.L1FinalizeWithdrawERC20:
			key := EventKey{BlockNumber: erc20EventUnmarshaler.Number, LogIndex: erc20EventUnmarshaler.Index - 1}
			messageHash, exists := messageHashes[key]
			if !exists {
				return fmt.Errorf("message hash does not exist for key %d", key)
			}
			tmpMessageMatch = orm.MessageMatch{
				MessageHash:   messageHash.Hex(),
				TokenType:     int(types.TokenTypeERC20),
				L1EventType:   int(erc20EventUnmarshaler.Type),
				L1BlockNumber: erc20EventUnmarshaler.Number,
				L1TxHash:      erc20EventUnmarshaler.TxHash.Hex(),
				L1Amount:      decimal.NewFromBigInt(erc20EventUnmarshaler.Amount, 0),
			}
			messageMatches = append(messageMatches, tmpMessageMatch)
		case types.L2WithdrawERC20:
			key := EventKey{BlockNumber: erc20EventUnmarshaler.Number, LogIndex: erc20EventUnmarshaler.Index + 1}
			messageHash, exists := messageHashes[key]
			if !exists {
				return fmt.Errorf("message hash does not exist for key %d", key)
			}
			tmpMessageMatch = orm.MessageMatch{
				MessageHash:   messageHash.Hex(),
				TokenType:     int(types.TokenTypeERC20),
				L2EventType:   int(erc20EventUnmarshaler.Type),
				L2BlockNumber: erc20EventUnmarshaler.Number,
				L2TxHash:      erc20EventUnmarshaler.TxHash.Hex(),
				L2Amount:      decimal.NewFromBigInt(erc20EventUnmarshaler.Amount, 0),
			}
			messageMatches = append(messageMatches, tmpMessageMatch)
		case types.L2FinalizeDepositERC20:
			key := EventKey{BlockNumber: erc20EventUnmarshaler.Number, LogIndex: erc20EventUnmarshaler.Index - 1}
			messageHash, exists := messageHashes[key]
			if !exists {
				return fmt.Errorf("message hash does not exist for key %d", key)
			}
			tmpMessageMatch = orm.MessageMatch{
				MessageHash:   messageHash.Hex(),
				TokenType:     int(types.TokenTypeERC20),
				L2EventType:   int(erc20EventUnmarshaler.Type),
				L2BlockNumber: erc20EventUnmarshaler.Number,
				L2TxHash:      erc20EventUnmarshaler.TxHash.Hex(),
				L2Amount:      decimal.NewFromBigInt(erc20EventUnmarshaler.Amount, 0),
			}
			messageMatches = append(messageMatches, tmpMessageMatch)
		}
	}

	effectRows, err := c.messageMatchOrm.InsertOrUpdate(ctx, messageMatches)
	if err != nil || effectRows != len(messageMatches) {
		return fmt.Errorf("erc20EventUnmarshaler orm insert failed, err: %w", err)
	}

	var transferEvents []events.ERC20GatewayEventUnmarshaler
	for _, eventData := range eventDataList {
		transferEventUnmarshaler, ok := eventData.(*events.ERC20GatewayEventUnmarshaler)
		if !ok {
			return fmt.Errorf("eventData is not of type *events.ERC20GatewayEventUnmarshaler")
		}
		transferEvents = append(transferEvents, *transferEventUnmarshaler)
	}

	return c.transferMatcher.Erc20Matcher(transferEvents, gatewayEvents)
}
