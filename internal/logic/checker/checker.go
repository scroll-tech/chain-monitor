package checker

import (
	"context"
	"fmt"
	"math"
	"sort"
	"strings"

	"github.com/scroll-tech/go-ethereum/common"
	"github.com/shopspring/decimal"
	"gorm.io/gorm"

	"github.com/scroll-tech/chain-monitor/internal/logic/events"
	"github.com/scroll-tech/chain-monitor/internal/logic/slack"
	"github.com/scroll-tech/chain-monitor/internal/orm"
	"github.com/scroll-tech/chain-monitor/internal/types"
	"github.com/scroll-tech/chain-monitor/internal/utils/msgproof"
)

// Checker is a structure that helps in verifying the data integrity
// in the blockchain by checking the message matches and the events.
type Checker struct {
	messageMatchOrm   *orm.MessageMatch
	transferMatcher   *TransferEventMatcher
	crossChainMatcher *CrossEventMatcher
}

// NewChecker returns a new Checker instance.
func NewChecker(db *gorm.DB) *Checker {
	return &Checker{
		messageMatchOrm:   orm.NewMessageMatch(db),
		transferMatcher:   NewTransferEventMatcher(),
		crossChainMatcher: NewCrossEventMatcher(),
	}
}

// CrossChainCheck checks the cross chain events.
func (c *Checker) CrossChainCheck(_ context.Context, layer types.LayerType, messageMatch orm.MessageMatch) types.MismatchType {
	switch layer {
	case types.Layer1:
		return c.crossChainMatcher.checkL1EventAndAmountMatchL2(messageMatch)
	case types.Layer2:
		return c.crossChainMatcher.checkL2EventAndAmountMatchL1(messageMatch)
	}
	return types.MismatchTypeValid
}

// UpdateBlockStatus updates the block status in the database.
func (c *Checker) UpdateBlockStatus(ctx context.Context, layer types.LayerType, start, end uint64) error {
	return c.messageMatchOrm.UpdateBlockStatus(ctx, layer, start, end)
}

// GatewayCheck checks the gateway events.
func (c *Checker) GatewayCheck(ctx context.Context, eventCategory types.EventCategory, gatewayEvents, messengerEvents, transferEvents []events.EventUnmarshaler) ([]orm.MessageMatch, error) {
	switch eventCategory {
	case types.ERC20EventCategory:
		return c.erc20EventUnmarshaler(ctx, gatewayEvents, messengerEvents, transferEvents)
	case types.ERC721EventCategory:
		return c.erc721EventUnmarshaler(ctx, gatewayEvents, messengerEvents, transferEvents)
	case types.ERC1155EventCategory:
		return c.erc1155EventUnmarshaler(ctx, gatewayEvents, messengerEvents, transferEvents)
	}
	return nil, nil
}

// CheckL2WithdrawRoots checks the L2 withdraw roots.
func (c *Checker) CheckL2WithdrawRoots(ctx context.Context, startBlockNumber, endBlockNumber uint64, messengerEventsData []events.EventUnmarshaler, withdrawRoots map[uint64]common.Hash) error {
	// recover latest withdraw trie.
	withdrawTrie := msgproof.NewWithdrawTrie()
	msg, err := c.messageMatchOrm.GetLargestMessageNonceL2MessageMatch(ctx)
	if err != nil {
		return err
	}
	if msg != nil {
		withdrawTrie.Initialize(msg.MessageNonce-1, common.HexToHash(msg.MessageHash), msg.MessageProof)
	}

	sentMessageEventHashesMap := make(map[uint64][]common.Hash)
	for _, eventData := range messengerEventsData {
		messengerEventUnmarshaler, ok := eventData.(*events.MessengerEventUnmarshaler)
		if !ok {
			return fmt.Errorf("eventData is not of type *events.MessengerEventUnmarshaler")
		}
		if messengerEventUnmarshaler.Type == types.L2SentMessage {
			blockNum := messengerEventUnmarshaler.Number
			sentMessageEventHashesMap[blockNum] = append(sentMessageEventHashesMap[blockNum], messengerEventUnmarshaler.MessageHash)
		}
	}

	var messageMatches []orm.MessageMatch
	for blockNum := startBlockNumber; blockNum <= endBlockNumber; blockNum++ {
		eventHashes := sentMessageEventHashesMap[blockNum]
		proofs := withdrawTrie.AppendMessages(eventHashes)
		lastWithdrawRoot := withdrawTrie.MessageRoot()
		if lastWithdrawRoot != withdrawRoots[blockNum] {
			info := slack.WithdrawRootInfo{
				BlockNumber:          blockNum,
				LastWithdrawRoot:     lastWithdrawRoot,
				ExpectedWithdrawRoot: withdrawRoots[blockNum],
			}
			slack.Notify(slack.MrkDwnWithdrawRootMessage(info))
			return fmt.Errorf("withdraw root mismatch in %v, got: %v, expected %v", blockNum, lastWithdrawRoot, withdrawRoots[blockNum])
		}
		// current block has SentMessage events.
		numEvents := len(eventHashes)
		if numEvents > 0 {
			// only update the last message of each block (which contains at least one SentMessage event).
			messageMatches = append(messageMatches, orm.MessageMatch{
				MessageHash:  eventHashes[numEvents-1].Hex(),
				MessageProof: proofs[numEvents-1],
				MessageNonce: withdrawTrie.NextMessageNonce, // +1 to distinguish from the zero value
			})
		}
	}

	effectRows, err := c.messageMatchOrm.InsertOrUpdateMsgProofNonce(ctx, messageMatches)
	if err != nil {
		return fmt.Errorf("message proof orm insert failed, err: %w", err)
	}

	if int(effectRows) != len(messageMatches) {
		return fmt.Errorf("message proof orm insert failed, effectRow:%d not equal messageMatches:%d", effectRows, len(messageMatches))
	}
	return nil
}

// MessengerCheck checks the messenger events.
func (c *Checker) MessengerCheck(_ context.Context, layer types.LayerType, messengerEvents []events.EventUnmarshaler) ([]orm.MessageMatch, error) {
	if len(messengerEvents) == 0 {
		return nil, nil
	}

	var messageMatches []orm.MessageMatch
	for _, eventData := range messengerEvents {
		var tmpMessageMatch orm.MessageMatch
		messengerEventUnmarshaler, ok := eventData.(*events.MessengerEventUnmarshaler)
		if !ok {
			return nil, fmt.Errorf("eventData is not of type *events.MessengerEventUnmarshaler")
		}
		switch messengerEventUnmarshaler.Type {
		case types.L1SentMessage:
			tmpMessageMatch = orm.MessageMatch{
				MessageHash:   messengerEventUnmarshaler.MessageHash.Hex(),
				TokenType:     int(types.TokenTypeETH), // default ETH.
				L1EventType:   int(messengerEventUnmarshaler.Type),
				L1BlockNumber: messengerEventUnmarshaler.Number,
				L1TxHash:      messengerEventUnmarshaler.TxHash.Hex(),
				L1Amounts:     decimal.NewFromBigInt(messengerEventUnmarshaler.Value, 0).String(),
				L2Amounts:     decimal.NewFromBigInt(messengerEventUnmarshaler.Value, 0).String(),
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
				L1Amounts:     decimal.NewFromBigInt(messengerEventUnmarshaler.Value, 0).String(),
				L2Amounts:     decimal.NewFromBigInt(messengerEventUnmarshaler.Value, 0).String(),
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

	sort.SliceStable(messageMatches, func(i, j int) bool {
		if layer == types.Layer1 {
			return messageMatches[i].L1BlockNumber < messageMatches[j].L1BlockNumber
		}
		return messageMatches[i].L2BlockNumber < messageMatches[j].L2BlockNumber
	})

	return messageMatches, nil
}

func (c *Checker) erc1155EventUnmarshaler(_ context.Context, gatewayEventsData, messengerEventsData, transferEventsData []events.EventUnmarshaler) ([]orm.MessageMatch, error) {
	messageHashes := make(map[messageEventKey]common.Hash)
	for _, eventData := range messengerEventsData {
		messengerEventUnmarshaler, ok := eventData.(*events.MessengerEventUnmarshaler)
		if !ok {
			return nil, fmt.Errorf("eventData is not of type *events.MessengerEventUnmarshaler")
		}
		key := messageEventKey{TxHash: messengerEventUnmarshaler.TxHash, LogIndex: messengerEventUnmarshaler.Index}
		messageHashes[key] = messengerEventUnmarshaler.MessageHash
	}

	var messageMatches []orm.MessageMatch
	var gatewayEvents []events.ERC1155GatewayEventUnmarshaler
	for _, eventData := range gatewayEventsData {
		erc1155EventUnmarshaler := eventData.(*events.ERC1155GatewayEventUnmarshaler)
		gatewayEvents = append(gatewayEvents, *erc1155EventUnmarshaler)

		var tmpMessageMatch orm.MessageMatch
		switch erc1155EventUnmarshaler.Type {
		case types.L1DepositERC1155:
			messageHash, exists := findPrevMessageEvent(erc1155EventUnmarshaler.TxHash, erc1155EventUnmarshaler.Index, messageHashes)
			if !exists {
				return nil, fmt.Errorf("message hash does not exist for erc1155 event %v", erc1155EventUnmarshaler)
			}
			var tokenIdsStrList []string
			for _, tokenID := range erc1155EventUnmarshaler.TokenIds {
				tokenIdsStrList = append(tokenIdsStrList, tokenID.String())
			}
			var amountStrList []string
			for _, amount := range erc1155EventUnmarshaler.Amounts {
				amountStrList = append(amountStrList, amount.String())
			}
			tmpMessageMatch = orm.MessageMatch{
				MessageHash:   messageHash.Hex(),
				TokenType:     int(types.TokenTypeERC1155),
				L1EventType:   int(erc1155EventUnmarshaler.Type),
				L1BlockNumber: erc1155EventUnmarshaler.Number,
				L1TxHash:      erc1155EventUnmarshaler.TxHash.Hex(),
				L1TokenIds:    strings.Join(tokenIdsStrList, ","),
				L1Amounts:     strings.Join(amountStrList, ","),
			}
			messageMatches = append(messageMatches, tmpMessageMatch)
			erc1155EventUnmarshaler.MessageHash = messageHash
		case types.L1FinalizeWithdrawERC1155:
			messageHash, exists := findNextMessageEvent(erc1155EventUnmarshaler.TxHash, erc1155EventUnmarshaler.Index, messageHashes)
			if !exists {
				return nil, fmt.Errorf("message hash does not exist for erc1155 event %v", erc1155EventUnmarshaler)
			}
			var tokenIdsStrList []string
			for _, tokenID := range erc1155EventUnmarshaler.TokenIds {
				tokenIdsStrList = append(tokenIdsStrList, tokenID.String())
			}
			var amountStrList []string
			for _, amount := range erc1155EventUnmarshaler.Amounts {
				amountStrList = append(amountStrList, amount.String())
			}
			tmpMessageMatch = orm.MessageMatch{
				MessageHash:   messageHash.Hex(),
				TokenType:     int(types.TokenTypeERC1155),
				L1EventType:   int(erc1155EventUnmarshaler.Type),
				L1BlockNumber: erc1155EventUnmarshaler.Number,
				L1TxHash:      erc1155EventUnmarshaler.TxHash.Hex(),
				L1TokenIds:    strings.Join(tokenIdsStrList, ","),
				L1Amounts:     strings.Join(amountStrList, ","),
			}
			messageMatches = append(messageMatches, tmpMessageMatch)
			erc1155EventUnmarshaler.MessageHash = messageHash
		case types.L2WithdrawERC1155:
			messageHash, exists := findPrevMessageEvent(erc1155EventUnmarshaler.TxHash, erc1155EventUnmarshaler.Index, messageHashes)
			if !exists {
				return nil, fmt.Errorf("message hash does not exist for erc1155 event %v", erc1155EventUnmarshaler)
			}
			var tokenIdsStrList []string
			for _, tokenID := range erc1155EventUnmarshaler.TokenIds {
				tokenIdsStrList = append(tokenIdsStrList, tokenID.String())
			}
			var amountStrList []string
			for _, amount := range erc1155EventUnmarshaler.Amounts {
				amountStrList = append(amountStrList, amount.String())
			}
			tmpMessageMatch = orm.MessageMatch{
				MessageHash:   messageHash.Hex(),
				TokenType:     int(types.TokenTypeERC1155),
				L2EventType:   int(erc1155EventUnmarshaler.Type),
				L2BlockNumber: erc1155EventUnmarshaler.Number,
				L2TxHash:      erc1155EventUnmarshaler.TxHash.Hex(),
				L2TokenIds:    strings.Join(tokenIdsStrList, ","),
				L2Amounts:     strings.Join(amountStrList, ","),
			}
			messageMatches = append(messageMatches, tmpMessageMatch)
			erc1155EventUnmarshaler.MessageHash = messageHash
		case types.L2FinalizeDepositERC1155:
			messageHash, exists := findNextMessageEvent(erc1155EventUnmarshaler.TxHash, erc1155EventUnmarshaler.Index, messageHashes)
			if !exists {
				return nil, fmt.Errorf("message hash does not exist for erc1155 event %v", erc1155EventUnmarshaler)
			}
			var tokenIdsStrList []string
			for _, tokenID := range erc1155EventUnmarshaler.TokenIds {
				tokenIdsStrList = append(tokenIdsStrList, tokenID.String())
			}
			var amountStrList []string
			for _, amount := range erc1155EventUnmarshaler.Amounts {
				amountStrList = append(amountStrList, amount.String())
			}
			tmpMessageMatch = orm.MessageMatch{
				MessageHash:   messageHash.Hex(),
				TokenType:     int(types.TokenTypeERC1155),
				L2EventType:   int(erc1155EventUnmarshaler.Type),
				L2BlockNumber: erc1155EventUnmarshaler.Number,
				L2TxHash:      erc1155EventUnmarshaler.TxHash.Hex(),
				L2TokenIds:    strings.Join(tokenIdsStrList, ","),
				L2Amounts:     strings.Join(amountStrList, ","),
			}
			messageMatches = append(messageMatches, tmpMessageMatch)
			erc1155EventUnmarshaler.MessageHash = messageHash
		}
	}

	var transferEvents []events.ERC1155GatewayEventUnmarshaler
	for _, eventData := range transferEventsData {
		transferEventUnmarshaler, ok := eventData.(*events.ERC1155GatewayEventUnmarshaler)
		if !ok {
			return messageMatches, fmt.Errorf("eventData is not of type *events.ERC1155GatewayEventUnmarshaler")
		}
		transferEvents = append(transferEvents, *transferEventUnmarshaler)
	}

	return messageMatches, c.transferMatcher.erc1155Matcher(transferEvents, gatewayEvents)
}

func (c *Checker) erc721EventUnmarshaler(_ context.Context, gatewayEventsData, messengerEventsData, transferEventsData []events.EventUnmarshaler) ([]orm.MessageMatch, error) {
	messageHashes := make(map[messageEventKey]common.Hash)
	for _, eventData := range messengerEventsData {
		messengerEventUnmarshaler, ok := eventData.(*events.MessengerEventUnmarshaler)
		if !ok {
			return nil, fmt.Errorf("eventData is not of type *events.MessengerEventUnmarshaler")
		}
		key := messageEventKey{TxHash: messengerEventUnmarshaler.TxHash, LogIndex: messengerEventUnmarshaler.Index}
		messageHashes[key] = messengerEventUnmarshaler.MessageHash
	}

	var messageMatches []orm.MessageMatch
	var gatewayEvents []events.ERC721GatewayEventUnmarshaler
	for _, eventData := range gatewayEventsData {
		erc721EventUnmarshaler := eventData.(*events.ERC721GatewayEventUnmarshaler)
		gatewayEvents = append(gatewayEvents, *erc721EventUnmarshaler)

		var tmpMessageMatch orm.MessageMatch
		switch erc721EventUnmarshaler.Type {
		case types.L1DepositERC721:
			messageHash, exists := findPrevMessageEvent(erc721EventUnmarshaler.TxHash, erc721EventUnmarshaler.Index, messageHashes)
			if !exists {
				return nil, fmt.Errorf("message hash does not exist for erc721 event %v", erc721EventUnmarshaler)
			}
			var tokenIdsStrList []string
			for _, tokenID := range erc721EventUnmarshaler.TokenIds {
				tokenIdsStrList = append(tokenIdsStrList, tokenID.String())
			}
			tmpMessageMatch = orm.MessageMatch{
				MessageHash:   messageHash.Hex(),
				TokenType:     int(types.TokenTypeERC721),
				L1EventType:   int(erc721EventUnmarshaler.Type),
				L1BlockNumber: erc721EventUnmarshaler.Number,
				L1TxHash:      erc721EventUnmarshaler.TxHash.Hex(),
				L1TokenIds:    strings.Join(tokenIdsStrList, ","),
			}
			messageMatches = append(messageMatches, tmpMessageMatch)
			erc721EventUnmarshaler.MessageHash = messageHash
		case types.L1FinalizeWithdrawERC721:
			messageHash, exists := findNextMessageEvent(erc721EventUnmarshaler.TxHash, erc721EventUnmarshaler.Index, messageHashes)
			if !exists {
				return nil, fmt.Errorf("message hash does not exist for erc721 event %v", erc721EventUnmarshaler)
			}
			var tokenIdsStrList []string
			for _, tokenID := range erc721EventUnmarshaler.TokenIds {
				tokenIdsStrList = append(tokenIdsStrList, tokenID.String())
			}
			tmpMessageMatch = orm.MessageMatch{
				MessageHash:   messageHash.Hex(),
				TokenType:     int(types.TokenTypeERC721),
				L1EventType:   int(erc721EventUnmarshaler.Type),
				L1BlockNumber: erc721EventUnmarshaler.Number,
				L1TxHash:      erc721EventUnmarshaler.TxHash.Hex(),
				L1TokenIds:    strings.Join(tokenIdsStrList, ","),
			}
			messageMatches = append(messageMatches, tmpMessageMatch)
			erc721EventUnmarshaler.MessageHash = messageHash
		case types.L2WithdrawERC721:
			messageHash, exists := findPrevMessageEvent(erc721EventUnmarshaler.TxHash, erc721EventUnmarshaler.Index, messageHashes)
			if !exists {
				return nil, fmt.Errorf("message hash does not exist for erc721 event %v", erc721EventUnmarshaler)
			}
			var tokenIdsStrList []string
			for _, tokenID := range erc721EventUnmarshaler.TokenIds {
				tokenIdsStrList = append(tokenIdsStrList, tokenID.String())
			}
			tmpMessageMatch = orm.MessageMatch{
				MessageHash:   messageHash.Hex(),
				TokenType:     int(types.TokenTypeERC721),
				L2EventType:   int(erc721EventUnmarshaler.Type),
				L2BlockNumber: erc721EventUnmarshaler.Number,
				L2TxHash:      erc721EventUnmarshaler.TxHash.Hex(),
				L2TokenIds:    strings.Join(tokenIdsStrList, ","),
			}
			messageMatches = append(messageMatches, tmpMessageMatch)
			erc721EventUnmarshaler.MessageHash = messageHash
		case types.L2FinalizeDepositERC721:
			messageHash, exists := findNextMessageEvent(erc721EventUnmarshaler.TxHash, erc721EventUnmarshaler.Index, messageHashes)
			if !exists {
				return nil, fmt.Errorf("message hash does not exist for erc721 event %v", erc721EventUnmarshaler)
			}
			var tokenIdsStrList []string
			for _, tokenID := range erc721EventUnmarshaler.TokenIds {
				tokenIdsStrList = append(tokenIdsStrList, tokenID.String())
			}
			tmpMessageMatch = orm.MessageMatch{
				MessageHash:   messageHash.Hex(),
				TokenType:     int(types.TokenTypeERC721),
				L2EventType:   int(erc721EventUnmarshaler.Type),
				L2BlockNumber: erc721EventUnmarshaler.Number,
				L2TxHash:      erc721EventUnmarshaler.TxHash.Hex(),
				L2TokenIds:    strings.Join(tokenIdsStrList, ","),
			}
			messageMatches = append(messageMatches, tmpMessageMatch)
			erc721EventUnmarshaler.MessageHash = messageHash
		}
	}

	var transferEvents []events.ERC721GatewayEventUnmarshaler
	for _, eventData := range transferEventsData {
		transferEventUnmarshaler, ok := eventData.(*events.ERC721GatewayEventUnmarshaler)
		if !ok {
			return messageMatches, fmt.Errorf("eventData is not of type *events.ERC721GatewayEventUnmarshaler")
		}
		transferEvents = append(transferEvents, *transferEventUnmarshaler)
	}

	return messageMatches, c.transferMatcher.erc721Matcher(transferEvents, gatewayEvents)
}

func (c *Checker) erc20EventUnmarshaler(_ context.Context, gatewayEventsData, messengerEventsData, transferEventsData []events.EventUnmarshaler) ([]orm.MessageMatch, error) {
	messageHashes := make(map[messageEventKey]common.Hash)
	for _, eventData := range messengerEventsData {
		messengerEventUnmarshaler, ok := eventData.(*events.MessengerEventUnmarshaler)
		if !ok {
			return nil, fmt.Errorf("erc20 eventData is not of type *events.MessengerEventUnmarshaler")
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
				return nil, fmt.Errorf("message hash does not exist for erc20 event %v", erc20EventUnmarshaler)
			}
			tmpMessageMatch = orm.MessageMatch{
				MessageHash:   messageHash.Hex(),
				TokenType:     int(types.TokenTypeERC20),
				L1EventType:   int(erc20EventUnmarshaler.Type),
				L1BlockNumber: erc20EventUnmarshaler.Number,
				L1TxHash:      erc20EventUnmarshaler.TxHash.Hex(),
				L1Amounts:     decimal.NewFromBigInt(erc20EventUnmarshaler.Amount, 0).String(),
			}
			messageMatches = append(messageMatches, tmpMessageMatch)
			erc20EventUnmarshaler.MessageHash = messageHash
		case types.L1FinalizeWithdrawERC20:
			messageHash, exists := findNextMessageEvent(erc20EventUnmarshaler.TxHash, erc20EventUnmarshaler.Index, messageHashes)
			if !exists {
				return nil, fmt.Errorf("message hash does not exist for erc20 event %v", erc20EventUnmarshaler)
			}
			tmpMessageMatch = orm.MessageMatch{
				MessageHash:   messageHash.Hex(),
				TokenType:     int(types.TokenTypeERC20),
				L1EventType:   int(erc20EventUnmarshaler.Type),
				L1BlockNumber: erc20EventUnmarshaler.Number,
				L1TxHash:      erc20EventUnmarshaler.TxHash.Hex(),
				L1Amounts:     decimal.NewFromBigInt(erc20EventUnmarshaler.Amount, 0).String(),
			}
			messageMatches = append(messageMatches, tmpMessageMatch)
			erc20EventUnmarshaler.MessageHash = messageHash
		case types.L2WithdrawERC20:
			messageHash, exists := findPrevMessageEvent(erc20EventUnmarshaler.TxHash, erc20EventUnmarshaler.Index, messageHashes)
			if !exists {
				return nil, fmt.Errorf("message hash does not exist for erc20 event %v", erc20EventUnmarshaler)
			}
			tmpMessageMatch = orm.MessageMatch{
				MessageHash:   messageHash.Hex(),
				TokenType:     int(types.TokenTypeERC20),
				L2EventType:   int(erc20EventUnmarshaler.Type),
				L2BlockNumber: erc20EventUnmarshaler.Number,
				L2TxHash:      erc20EventUnmarshaler.TxHash.Hex(),
				L2Amounts:     decimal.NewFromBigInt(erc20EventUnmarshaler.Amount, 0).String(),
			}
			messageMatches = append(messageMatches, tmpMessageMatch)
			erc20EventUnmarshaler.MessageHash = messageHash
		case types.L2FinalizeDepositERC20:
			messageHash, exists := findNextMessageEvent(erc20EventUnmarshaler.TxHash, erc20EventUnmarshaler.Index, messageHashes)
			if !exists {
				return nil, fmt.Errorf("message hash does not exist for erc20 event %v", erc20EventUnmarshaler)
			}
			tmpMessageMatch = orm.MessageMatch{
				MessageHash:   messageHash.Hex(),
				TokenType:     int(types.TokenTypeERC20),
				L2EventType:   int(erc20EventUnmarshaler.Type),
				L2BlockNumber: erc20EventUnmarshaler.Number,
				L2TxHash:      erc20EventUnmarshaler.TxHash.Hex(),
				L2Amounts:     decimal.NewFromBigInt(erc20EventUnmarshaler.Amount, 0).String(),
			}
			messageMatches = append(messageMatches, tmpMessageMatch)
			erc20EventUnmarshaler.MessageHash = messageHash
		}
	}

	var transferEvents []events.ERC20GatewayEventUnmarshaler
	for _, eventData := range transferEventsData {
		transferEventUnmarshaler, ok := eventData.(*events.ERC20GatewayEventUnmarshaler)
		if !ok {
			return messageMatches, fmt.Errorf("eventData is not of type *events.ERC20GatewayEventUnmarshaler")
		}
		transferEvents = append(transferEvents, *transferEventUnmarshaler)
	}

	err := c.transferMatcher.erc20Matcher(transferEvents, gatewayEvents)
	return messageMatches, err
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
