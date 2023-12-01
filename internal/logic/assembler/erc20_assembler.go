package assembler

import (
	"fmt"

	"github.com/scroll-tech/go-ethereum/common"
	"github.com/shopspring/decimal"

	"github.com/scroll-tech/chain-monitor/internal/logic/events"
	"github.com/scroll-tech/chain-monitor/internal/orm"
	"github.com/scroll-tech/chain-monitor/internal/types"
)

func (c *MessageMatchAssembler) erc20EventMessageMatchAssembler(gatewayEventsData, messengerEventsData, transferEventsData []events.EventUnmarshaler) ([]orm.GatewayMessageMatch, error) {
	messageHashes := make(map[messageEventKey]common.Hash)
	for _, eventData := range messengerEventsData {
		messengerEventUnmarshaler, ok := eventData.(*events.MessengerEventUnmarshaler)
		if !ok {
			return nil, fmt.Errorf("erc20 eventData is not of type *events.MessengerEventUnmarshaler")
		}
		key := messageEventKey{TxHash: messengerEventUnmarshaler.TxHash, LogIndex: messengerEventUnmarshaler.Index}
		messageHashes[key] = messengerEventUnmarshaler.MessageHash
	}

	var messageMatches []orm.GatewayMessageMatch
	var gatewayEvents []events.ERC20GatewayEventUnmarshaler
	for _, eventData := range gatewayEventsData {
		erc20EventUnmarshaler := eventData.(*events.ERC20GatewayEventUnmarshaler)
		gatewayEvents = append(gatewayEvents, *erc20EventUnmarshaler)

		var tmpMessageMatch orm.GatewayMessageMatch
		switch erc20EventUnmarshaler.Type {
		case types.L1DepositERC20:
			messageHash, exists := c.findPrevMessageEvent(erc20EventUnmarshaler.TxHash, erc20EventUnmarshaler.Index, messageHashes)
			if !exists {
				return nil, fmt.Errorf("message hash does not exist for erc20 event %v", erc20EventUnmarshaler)
			}
			tmpMessageMatch = orm.GatewayMessageMatch{
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
			messageHash, exists := c.findNextMessageEvent(erc20EventUnmarshaler.TxHash, erc20EventUnmarshaler.Index, messageHashes)
			if !exists {
				return nil, fmt.Errorf("message hash does not exist for erc20 event %v", erc20EventUnmarshaler)
			}
			tmpMessageMatch = orm.GatewayMessageMatch{
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
			messageHash, exists := c.findPrevMessageEvent(erc20EventUnmarshaler.TxHash, erc20EventUnmarshaler.Index, messageHashes)
			if !exists {
				return nil, fmt.Errorf("message hash does not exist for erc20 event %v", erc20EventUnmarshaler)
			}
			tmpMessageMatch = orm.GatewayMessageMatch{
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
			messageHash, exists := c.findNextMessageEvent(erc20EventUnmarshaler.TxHash, erc20EventUnmarshaler.Index, messageHashes)
			if !exists {
				return nil, fmt.Errorf("message hash does not exist for erc20 event %v", erc20EventUnmarshaler)
			}
			tmpMessageMatch = orm.GatewayMessageMatch{
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
