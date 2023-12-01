package checker

import (
	"fmt"
	"strings"

	"github.com/scroll-tech/go-ethereum/common"

	"github.com/scroll-tech/chain-monitor/internal/logic/events"
	"github.com/scroll-tech/chain-monitor/internal/orm"
	"github.com/scroll-tech/chain-monitor/internal/types"
)

func (c *Checker) erc721EventUnmarshaler(gatewayEventsData, messengerEventsData, transferEventsData []events.EventUnmarshaler) ([]orm.GatewayMessageMatch, error) {
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
	var gatewayEvents []events.ERC721GatewayEventUnmarshaler
	for _, eventData := range gatewayEventsData {
		erc721EventUnmarshaler := eventData.(*events.ERC721GatewayEventUnmarshaler)
		gatewayEvents = append(gatewayEvents, *erc721EventUnmarshaler)

		var tmpMessageMatch orm.GatewayMessageMatch
		switch erc721EventUnmarshaler.Type {
		case types.L1DepositERC721:
			messageHash, exists := c.findPrevMessageEvent(erc721EventUnmarshaler.TxHash, erc721EventUnmarshaler.Index, messageHashes)
			if !exists {
				return nil, fmt.Errorf("message hash does not exist for erc721 event %v", erc721EventUnmarshaler)
			}
			var tokenIdsStrList []string
			for _, tokenID := range erc721EventUnmarshaler.TokenIds {
				tokenIdsStrList = append(tokenIdsStrList, tokenID.String())
			}
			tmpMessageMatch = orm.GatewayMessageMatch{
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
			messageHash, exists := c.findNextMessageEvent(erc721EventUnmarshaler.TxHash, erc721EventUnmarshaler.Index, messageHashes)
			if !exists {
				return nil, fmt.Errorf("message hash does not exist for erc721 event %v", erc721EventUnmarshaler)
			}
			var tokenIdsStrList []string
			for _, tokenID := range erc721EventUnmarshaler.TokenIds {
				tokenIdsStrList = append(tokenIdsStrList, tokenID.String())
			}
			tmpMessageMatch = orm.GatewayMessageMatch{
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
			messageHash, exists := c.findPrevMessageEvent(erc721EventUnmarshaler.TxHash, erc721EventUnmarshaler.Index, messageHashes)
			if !exists {
				return nil, fmt.Errorf("message hash does not exist for erc721 event %v", erc721EventUnmarshaler)
			}
			var tokenIdsStrList []string
			for _, tokenID := range erc721EventUnmarshaler.TokenIds {
				tokenIdsStrList = append(tokenIdsStrList, tokenID.String())
			}
			tmpMessageMatch = orm.GatewayMessageMatch{
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
			messageHash, exists := c.findNextMessageEvent(erc721EventUnmarshaler.TxHash, erc721EventUnmarshaler.Index, messageHashes)
			if !exists {
				return nil, fmt.Errorf("message hash does not exist for erc721 event %v", erc721EventUnmarshaler)
			}
			var tokenIdsStrList []string
			for _, tokenID := range erc721EventUnmarshaler.TokenIds {
				tokenIdsStrList = append(tokenIdsStrList, tokenID.String())
			}
			tmpMessageMatch = orm.GatewayMessageMatch{
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
