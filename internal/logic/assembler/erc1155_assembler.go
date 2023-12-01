package assembler

import (
	"fmt"
	"strings"

	"github.com/scroll-tech/go-ethereum/common"

	"github.com/scroll-tech/chain-monitor/internal/logic/events"
	"github.com/scroll-tech/chain-monitor/internal/orm"
	"github.com/scroll-tech/chain-monitor/internal/types"
)

func (c *MessageMatchAssembler) erc1155EventMessageMatchAssembler(gatewayEventsData, messengerEventsData, transferEventsData []events.EventUnmarshaler) ([]orm.GatewayMessageMatch, error) {
	messageHashes := make(map[messageEventKey]common.Hash)
	for _, eventData := range messengerEventsData {
		messengerEventUnmarshaler, ok := eventData.(*events.MessengerEventUnmarshaler)
		if !ok {
			return nil, fmt.Errorf("eventData is not of type *events.MessengerEventUnmarshaler")
		}
		key := messageEventKey{TxHash: messengerEventUnmarshaler.TxHash, LogIndex: messengerEventUnmarshaler.Index}
		messageHashes[key] = messengerEventUnmarshaler.MessageHash
	}

	var messageMatches []orm.GatewayMessageMatch
	var gatewayEvents []events.ERC1155GatewayEventUnmarshaler
	for _, eventData := range gatewayEventsData {
		erc1155EventUnmarshaler := eventData.(*events.ERC1155GatewayEventUnmarshaler)
		gatewayEvents = append(gatewayEvents, *erc1155EventUnmarshaler)

		var tmpMessageMatch orm.GatewayMessageMatch
		switch erc1155EventUnmarshaler.Type {
		case types.L1DepositERC1155:
			messageHash, exists := c.findPrevMessageEvent(erc1155EventUnmarshaler.TxHash, erc1155EventUnmarshaler.Index, messageHashes)
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
			tmpMessageMatch = orm.GatewayMessageMatch{
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
			messageHash, exists := c.findNextMessageEvent(erc1155EventUnmarshaler.TxHash, erc1155EventUnmarshaler.Index, messageHashes)
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
			tmpMessageMatch = orm.GatewayMessageMatch{
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
			messageHash, exists := c.findPrevMessageEvent(erc1155EventUnmarshaler.TxHash, erc1155EventUnmarshaler.Index, messageHashes)
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
			tmpMessageMatch = orm.GatewayMessageMatch{
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
			messageHash, exists := c.findNextMessageEvent(erc1155EventUnmarshaler.TxHash, erc1155EventUnmarshaler.Index, messageHashes)
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
			tmpMessageMatch = orm.GatewayMessageMatch{
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
