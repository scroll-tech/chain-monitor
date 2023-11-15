package checker

import (
	"context"
	"fmt"
	"math"

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
		if !c.crossChainMatcher.checkL1EventMatchL2(messageMatch) {
			return types.MismatchTypeL2EventNotExist
		}
	}

	if layer == types.Layer2 {
		if !c.crossChainMatcher.checkL2EventMatchL1(messageMatch) {
			return types.MismatchTypeL1EventNotExist
		}
	}

	return types.MismatchTypeOk
}

func (c *Checker) GatewayCheck(ctx context.Context, eventCategory types.TxEventCategory, gatewayEvents, messengerEvents, transferEvents []events.EventUnmarshaler) error {
	switch eventCategory {
	case types.ERC20EventCategory:
		return c.erc20EventUnmarshaler(ctx, gatewayEvents, messengerEvents, transferEvents)
	case types.ERC721EventCategory:
	case types.ERC1155EventCategory:
	}
	return nil
}

func (c *Checker) CheckL2WithdrawRoots(ctx context.Context, startBlockNumber, endBlockNumber uint64, messengerEventsData []events.EventUnmarshaler, withdrawRoots map[uint64]common.Hash) error {
	// recover latest withdraw trie.
	withdrawTrie := msgproof.NewWithdrawTrie()
	if startBlockNumber > 1 {
		msg, err := c.messageMatchOrm.GetMessageMatchByL2BlockNumber(ctx, startBlockNumber-1)
		if err != nil {
			return err
		}
		withdrawTrie.Initialize(msg.MessageNonce, common.HexToHash(msg.MessageHash), msg.MessageProof)
	}

	sentMessageEventHashesMap := make(map[uint64][]common.Hash)
	for _, eventData := range messengerEventsData {
		erc20EventUnmarshaler, ok := eventData.(*events.MessengerEventUnmarshaler)
		if !ok {
			return fmt.Errorf("eventData is not of type *events.ERC20GatewayEventUnmarshaler")
		}
		if erc20EventUnmarshaler.Type == types.L2SentMessage {
			blockNum := erc20EventUnmarshaler.Number
			sentMessageEventHashesMap[blockNum] = append(sentMessageEventHashesMap[blockNum], erc20EventUnmarshaler.MessageHash)
		}
	}

	var messageMatches []orm.MessageMatch
	lastWithdrawRoot := withdrawTrie.MessageRoot()
	for blockNum := startBlockNumber; blockNum <= endBlockNumber; blockNum++ {
		eventHashes := sentMessageEventHashesMap[blockNum]
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
		return fmt.Errorf("message proof orm insert failed, err: %w", err)
	}
	return nil
}

func (c *Checker) MessengerCheck(ctx context.Context, messengerEvents []events.EventUnmarshaler) error {
	var messageMatches []orm.MessageMatch
	for _, eventData := range messengerEvents {
		var tmpMessageMatch orm.MessageMatch
		messengerEventUnmarshaler, ok := eventData.(*events.MessengerEventUnmarshaler)
		if !ok {
			return fmt.Errorf("eventData is not of type *events.MessengerEventUnmarshaler")
		}
		switch messengerEventUnmarshaler.Type {
		case types.L1SentMessage:
			tmpMessageMatch = orm.MessageMatch{
				MessageHash:   messengerEventUnmarshaler.MessageHash.Hex(),
				TokenType:     int(types.TokenTypeETH), // default ETH.
				L1EventType:   int(messengerEventUnmarshaler.Type),
				L1BlockNumber: messengerEventUnmarshaler.Number,
				L1TxHash:      messengerEventUnmarshaler.TxHash.Hex(),
				L1Amount:      decimal.NewFromBigInt(messengerEventUnmarshaler.Value, 0),
				L2Amount:      decimal.NewFromBigInt(messengerEventUnmarshaler.Value, 0),
			}
			messageMatches = append(messageMatches, tmpMessageMatch)
		case types.L1RelayedMessage:
			tmpMessageMatch = orm.MessageMatch{
				MessageHash:   messengerEventUnmarshaler.MessageHash.Hex(),
				TokenType:     int(types.TokenTypeETH), // default ETH.
				L1EventType:   int(messengerEventUnmarshaler.Type),
				L1BlockNumber: messengerEventUnmarshaler.Number,
				L1TxHash:      messengerEventUnmarshaler.TxHash.Hex(),
			}
			messageMatches = append(messageMatches, tmpMessageMatch)
		case types.L2SentMessage:
			tmpMessageMatch = orm.MessageMatch{
				MessageHash:   messengerEventUnmarshaler.MessageHash.Hex(),
				TokenType:     int(types.TokenTypeETH), // default ETH.
				L2EventType:   int(messengerEventUnmarshaler.Type),
				L2BlockNumber: messengerEventUnmarshaler.Number,
				L2TxHash:      messengerEventUnmarshaler.TxHash.Hex(),
				L1Amount:      decimal.NewFromBigInt(messengerEventUnmarshaler.Value, 0),
				L2Amount:      decimal.NewFromBigInt(messengerEventUnmarshaler.Value, 0),
			}
			messageMatches = append(messageMatches, tmpMessageMatch)
		case types.L2RelayedMessage:
			tmpMessageMatch = orm.MessageMatch{
				MessageHash:   messengerEventUnmarshaler.MessageHash.Hex(),
				TokenType:     int(types.TokenTypeETH), // default ETH.
				L2EventType:   int(messengerEventUnmarshaler.Type),
				L2BlockNumber: messengerEventUnmarshaler.Number,
				L2TxHash:      messengerEventUnmarshaler.TxHash.Hex(),
			}
			messageMatches = append(messageMatches, tmpMessageMatch)
		}
	}

	effectRows, err := c.messageMatchOrm.InsertOrUpdate(ctx, messageMatches)
	if err != nil || effectRows != len(messageMatches) {
		return fmt.Errorf("messenger event orm insert failed, err: %w", err)
	}
	return nil
}

func (c *Checker) erc20EventUnmarshaler(ctx context.Context, gatewayEventsData, messengerEventsData, transferEventsData []events.EventUnmarshaler) error {
	messageHashes := make(map[messageEventKey]common.Hash)
	for _, eventData := range messengerEventsData {
		messengerEventUnmarshaler, ok := eventData.(*events.MessengerEventUnmarshaler)
		if !ok {
			return fmt.Errorf("eventData is not of type *events.MessengerEventUnmarshaler")
		}
		key := messageEventKey{TxHash: messengerEventUnmarshaler.TxHash, LogIndex: messengerEventUnmarshaler.Index}
		messageHashes[key] = messengerEventUnmarshaler.MessageHash
	}

	var messageMatches []orm.MessageMatch
	var gatewayEvents []events.ERC20GatewayEventUnmarshaler
	for _, eventData := range gatewayEventsData {
		erc20EventUnmarshaler := eventData.(*events.ERC20GatewayEventUnmarshaler)
		gatewayEvents = append(gatewayEvents, *erc20EventUnmarshaler)

		var tmpMessageMatch orm.MessageMatch
		switch erc20EventUnmarshaler.Type {
		case types.L1DepositERC20:
			messageHash, exists := findPrevMessageEvent(erc20EventUnmarshaler.TxHash, erc20EventUnmarshaler.Index, messageHashes)
			if !exists {
				return fmt.Errorf("message hash does not exist for erc20 event %v", erc20EventUnmarshaler)
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
			messageHash, exists := findNextMessageEvent(erc20EventUnmarshaler.TxHash, erc20EventUnmarshaler.Index, messageHashes)
			if !exists {
				return fmt.Errorf("message hash does not exist for erc20 event %v", erc20EventUnmarshaler)
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
			messageHash, exists := findPrevMessageEvent(erc20EventUnmarshaler.TxHash, erc20EventUnmarshaler.Index, messageHashes)
			if !exists {
				return fmt.Errorf("message hash does not exist for erc20 event %v", erc20EventUnmarshaler)
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
			messageHash, exists := findNextMessageEvent(erc20EventUnmarshaler.TxHash, erc20EventUnmarshaler.Index, messageHashes)
			if !exists {
				return fmt.Errorf("message hash does not exist for erc20 event %v", erc20EventUnmarshaler)
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
	for _, eventData := range transferEventsData {
		transferEventUnmarshaler, ok := eventData.(*events.ERC20GatewayEventUnmarshaler)
		if !ok {
			return fmt.Errorf("eventData is not of type *events.ERC20GatewayEventUnmarshaler")
		}
		transferEvents = append(transferEvents, *transferEventUnmarshaler)
	}

	return c.transferMatcher.Erc20Matcher(transferEvents, gatewayEvents)
}

type messageEventKey struct {
	TxHash   common.Hash
	LogIndex uint
}

func findNextMessageEvent(txHash common.Hash, logIndex uint, messageHashes map[messageEventKey]common.Hash) (common.Hash, bool) {
	var nextMessageHash common.Hash
	var found bool
	var smallestDiff uint = math.MaxUint
	for key, msgHash := range messageHashes {
		if key.TxHash == txHash && key.LogIndex > logIndex {
			if diff := key.LogIndex - logIndex; diff < smallestDiff {
				smallestDiff = diff
				nextMessageHash = msgHash
				found = true
			}
		}
	}
	return nextMessageHash, found
}

func findPrevMessageEvent(txHash common.Hash, logIndex uint, messageHashes map[messageEventKey]common.Hash) (common.Hash, bool) {
	var prevMessageHash common.Hash
	var found bool
	var smallestDiff uint = math.MaxUint
	for key, msgHash := range messageHashes {
		if key.TxHash == txHash && key.LogIndex < logIndex {
			if diff := logIndex - key.LogIndex; diff < smallestDiff {
				smallestDiff = diff
				prevMessageHash = msgHash
				found = true
			}
		}
	}
	return prevMessageHash, found
}
