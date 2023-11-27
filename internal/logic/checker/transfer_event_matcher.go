package checker

import (
	"fmt"
	"math/big"

	"github.com/scroll-tech/go-ethereum/common"

	"github.com/scroll-tech/chain-monitor/internal/logic/events"
	"github.com/scroll-tech/chain-monitor/internal/logic/slack"
	"github.com/scroll-tech/chain-monitor/internal/types"
)

const (
	transferEventDontHaveGatewayEvent        = "the gateway event associated with the transfer event does not exist"
	transferEventBalanceMismatchGatewayEvent = "the transfer event's balance don't match the balance of gateway event"

	gatewayEventDontHaveTransferEvent        = "the transfer event associated with the gateway event does not exist"
	gatewayEventBalanceMismatchTransferEvent = "the gateway event's balance don't match the balance of transfer event"
)

type erc20MatcherKey struct {
	tokenAddress common.Address
	txHash       common.Hash
}

type erc721MatchKey struct {
	tokenAddress common.Address
	tokenID      *big.Int
	txHash       common.Hash
}

type erc1155MatchKey struct {
	tokenAddress common.Address
	tokenID      *big.Int
	txHash       common.Hash
}

type matcherValue struct {
	balance     *big.Int
	tokenType   types.TokenType
	layer       types.LayerType
	eventType   types.EventType
	blockNumber uint64
	messageHash common.Hash
}

// TransferEventMatcher checks the existence of an event and consistency of the transferred amount.
type TransferEventMatcher struct{}

// NewTransferEventMatcher creates a new instance of TransferEventMatcher.
func NewTransferEventMatcher() *TransferEventMatcher {
	return &TransferEventMatcher{}
}

func (t *TransferEventMatcher) erc20Matcher(transferEvents, gatewayEvents []events.ERC20GatewayEventUnmarshaler) error {
	transferBalances := make(map[erc20MatcherKey]matcherValue)
	gatewayBalances := make(map[erc20MatcherKey]matcherValue)

	for _, event := range gatewayEvents {
		key := erc20MatcherKey{
			tokenAddress: event.TokenAddress,
			txHash:       event.TxHash,
		}
		if _, exists := gatewayBalances[key]; !exists {
			gatewayBalances[key] = matcherValue{
				tokenType:   types.TokenTypeERC20,
				layer:       event.Layer,
				eventType:   event.Type,
				blockNumber: event.Number,
				messageHash: event.MessageHash,
				balance:     common.Big0,
			}
		}
		if event.Type == types.L1DepositERC20 || event.Type == types.L2WithdrawERC20 {
			gatewayBalances[key].balance.Add(gatewayBalances[key].balance, event.Amount)
		} else if event.Type == types.L2FinalizeDepositERC20 || event.Type == types.L1FinalizeWithdrawERC20 || event.Type == types.L1RefundERC20 {
			gatewayBalances[key].balance.Sub(gatewayBalances[key].balance, event.Amount)
		}
	}

	for _, event := range transferEvents {
		key := erc20MatcherKey{
			tokenAddress: event.TokenAddress,
			txHash:       event.TxHash,
		}
		// filter airdrop Transfers.
		_, exists := gatewayBalances[key]
		if !exists && event.Amount.Sign() >= 0 {
			continue
		}
		if _, exists := transferBalances[key]; !exists {
			transferBalances[key] = matcherValue{
				tokenType:   types.TokenTypeERC20,
				layer:       event.Layer,
				eventType:   event.Type,
				blockNumber: event.Number,
				messageHash: event.MessageHash,
				balance:     common.Big0,
			}
		}
		transferBalances[key].balance.Add(transferBalances[k].balance, event.Amount)
	}

	for transferMatcherKey, transferMatcherValue := range transferBalances {
		gatewayMatcherValue, exists := gatewayBalances[transferMatcherKey]
		// If the corresponding Transfer does not exist,
		// or if the tokens of Transfer events < the difference of tokens in gateway events (more tokens out).
		if !exists || transferMatcherValue.balance.Cmp(gatewayMatcherValue.balance) < 0 {
			info := slack.GatewayTransferInfo{
				TokenAddress:    transferMatcherKey.tokenAddress,
				TokenType:       transferMatcherValue.tokenType,
				Layer:           transferMatcherValue.layer,
				EventType:       transferMatcherValue.eventType,
				BlockNumber:     transferMatcherValue.blockNumber,
				TxHash:          transferMatcherKey.txHash,
				MessageHash:     transferMatcherValue.messageHash,
				TransferBalance: transferMatcherValue.balance,
			}
			if !exists {
				// Ignore additional Transfer events in Layer2.
				if info.Layer == types.Layer2 {
					continue
				}
				info.Error = transferEventDontHaveGatewayEvent
			} else {
				info.Error = transferEventBalanceMismatchGatewayEvent
				info.GatewayBalance = gatewayMatcherValue.balance
			}
			slack.Notify(slack.MrkDwnGatewayTransferMessage(info))
			return fmt.Errorf("balance mismatch for token %s: transfer balance = %s, gateway balance = %s, info = %v",
				info.TokenAddress.Hex(), info.TransferBalance.String(), info.GatewayBalance.String(), info)
		}
	}

	for gatewayMatcherKey, gatewayMatcherValue := range gatewayBalances {
		transferMatcherValue, exists := transferBalances[gatewayMatcherKey]
		// If the corresponding Transfer does not exist,
		// or if the tokens of Transfer events < the difference of tokens in gateway events (more tokens out).
		if !exists || gatewayMatcherValue.balance.Cmp(transferMatcherValue.balance) > 0 {
			info := slack.GatewayTransferInfo{
				TokenAddress:   gatewayMatcherKey.tokenAddress,
				TokenType:      gatewayMatcherValue.tokenType,
				Layer:          gatewayMatcherValue.layer,
				EventType:      gatewayMatcherValue.eventType,
				BlockNumber:    gatewayMatcherValue.blockNumber,
				TxHash:         gatewayMatcherKey.txHash,
				MessageHash:    gatewayMatcherValue.messageHash,
				GatewayBalance: gatewayMatcherValue.balance,
			}
			if !exists {
				info.Error = gatewayEventDontHaveTransferEvent
			} else {
				info.Error = gatewayEventBalanceMismatchTransferEvent
				info.TransferBalance = transferMatcherValue.balance
			}
			slack.Notify(slack.MrkDwnGatewayTransferMessage(info))
			return fmt.Errorf("balance mismatch for token %s: gateway balance = %s, transfer balance = %s, info = %v",
				info.TokenAddress.Hex(), info.GatewayBalance.String(), info.TransferBalance.String(), info)
		}
	}
	return nil
}

func (t *TransferEventMatcher) erc721Matcher(transferEvents, gatewayEvents []events.ERC721GatewayEventUnmarshaler) error {
	transferTokenIds := make(map[erc721MatchKey]matcherValue)
	gatewayTokenIds := make(map[erc721MatchKey]matcherValue)

	for _, event := range gatewayEvents {
		if len(event.TokenIds) != len(event.Amounts) {
			return fmt.Errorf("erc1155 gateway event tokenIds and amounts not match, %v", event)
		}

		for idx, tokenID := range event.TokenIds {
			key := erc721MatchKey{
				tokenAddress: event.TokenAddress,
				tokenID:      tokenID,
				txHash:       event.TxHash,
			}

			if _, exists := gatewayTokenIds[key]; !exists {
				gatewayTokenIds[key] = matcherValue{
					tokenType:   types.TokenTypeERC721,
					layer:       event.Layer,
					eventType:   event.Type,
					blockNumber: event.Number,
					messageHash: event.MessageHash,
					balance:     common.Big0,
				}
			}

			if event.Type == types.L1DepositERC721 || event.Type == types.L2WithdrawERC721 ||
				event.Type == types.L1BatchDepositERC721 || event.Type == types.L2BatchWithdrawERC721 {
				gatewayTokenIds[key].balance.Add(gatewayTokenIds[key].balance, event.Amounts[idx])
			} else if event.Type == types.L2FinalizeDepositERC721 || event.Type == types.L1FinalizeWithdrawERC721 || event.Type == types.L1RefundERC721 ||
				event.Type == types.L2FinalizeBatchDepositERC721 || event.Type == types.L1FinalizeBatchWithdrawERC721 || event.Type == types.L1BatchRefundERC721 {
				gatewayTokenIds[key].balance.Sub(gatewayTokenIds[key].balance, event.Amounts[idx])
			}
		}
	}

	for _, event := range transferEvents {
		if len(event.TokenIds) != len(event.Amounts) {
			return fmt.Errorf("erc721 transfer event tokenIds and amounts not match, %v", event)
		}

		for idx, tokenID := range event.TokenIds {
			key := erc721MatchKey{
				tokenAddress: event.TokenAddress,
				tokenID:      tokenID,
				txHash:       event.TxHash,
			}
			// filter airdrop Transfers.
			_, exists := gatewayTokenIds[key]
			if !exists && event.Amounts[idx].Sign() >= 0 {
				continue
			}
			if _, exists := transferTokenIds[key]; !exists {
				transferTokenIds[key] = matcherValue{
					tokenType:   types.TokenTypeERC721,
					layer:       event.Layer,
					eventType:   event.Type,
					blockNumber: event.Number,
					messageHash: event.MessageHash,
					balance:     common.Big0,
				}
			}
			transferTokenIds[key].balance.Add(transferTokenIds[key].balance, event.Amounts[idx])
		}
	}

	for transferMatcherKey, transferMatcherValue := range transferTokenIds {
		gatewayMatcherValue, exists := gatewayTokenIds[transferMatcherKey]
		// If the corresponding Transfer does not exist,
		// or if the tokens of Transfer events < the difference of tokens in gateway events (more tokens out).
		if !exists || transferMatcherValue.balance.Cmp(gatewayMatcherValue.balance) < 0 {
			info := slack.GatewayTransferInfo{
				TokenAddress:    transferMatcherKey.tokenAddress,
				TokenType:       transferMatcherValue.tokenType,
				Layer:           transferMatcherValue.layer,
				EventType:       transferMatcherValue.eventType,
				BlockNumber:     transferMatcherValue.blockNumber,
				TxHash:          transferMatcherKey.txHash,
				MessageHash:     transferMatcherValue.messageHash,
				TransferBalance: transferMatcherValue.balance,
			}
			if !exists {
				// Ignore additional Transfer events in Layer2.
				if info.Layer == types.Layer2 {
					continue
				}
				info.Error = transferEventDontHaveGatewayEvent
			} else {
				info.Error = transferEventBalanceMismatchGatewayEvent
				info.GatewayBalance = gatewayMatcherValue.balance
			}
			slack.Notify(slack.MrkDwnGatewayTransferMessage(info))
			return fmt.Errorf("erc721 mismatch for tokenAddress %s: transfer amount = %s, gateway amount = %s",
				info.TokenAddress.Hex(), info.TransferBalance.String(), info.GatewayBalance.String())
		}
	}

	for gatewayMatcherKey, gatewayMatcherValue := range gatewayTokenIds {
		transferMatcherValue, exists := transferTokenIds[gatewayMatcherKey]
		// If the corresponding Transfer does not exist,
		// or if the tokens of Transfer events < the difference of tokens in gateway events (more tokens out).
		if !exists || gatewayMatcherValue.balance.Cmp(transferMatcherValue.balance) > 0 {
			info := slack.GatewayTransferInfo{
				TokenAddress:   gatewayMatcherKey.tokenAddress,
				TokenType:      gatewayMatcherValue.tokenType,
				Layer:          gatewayMatcherValue.layer,
				EventType:      gatewayMatcherValue.eventType,
				BlockNumber:    gatewayMatcherValue.blockNumber,
				TxHash:         gatewayMatcherKey.txHash,
				MessageHash:    gatewayMatcherValue.messageHash,
				GatewayBalance: gatewayMatcherValue.balance,
			}
			if !exists {
				info.Error = gatewayEventDontHaveTransferEvent
			} else {
				info.Error = gatewayEventBalanceMismatchTransferEvent
				info.TransferBalance = transferMatcherValue.balance
			}
			slack.Notify(slack.MrkDwnGatewayTransferMessage(info))
			return fmt.Errorf("erc721 mismatch for tokenAddress %s: gateway amount = %s, transfer amount = %s",
				info.TokenAddress.Hex(), info.GatewayBalance.String(), info.TransferBalance.String())
		}
	}

	return nil
}

func (t *TransferEventMatcher) erc1155Matcher(transferEvents, gatewayEvents []events.ERC1155GatewayEventUnmarshaler) error {
	transferTokenIds := make(map[erc1155MatchKey]matcherValue)
	gatewayTokenIds := make(map[erc1155MatchKey]matcherValue)

	for _, event := range gatewayEvents {
		if len(event.TokenIds) != len(event.Amounts) {
			return fmt.Errorf("erc1155 gateway event tokenIds and amounts not match, %v", event)
		}

		for idx, tokenID := range event.TokenIds {
			key := erc1155MatchKey{
				tokenAddress: event.TokenAddress,
				tokenID:      tokenID,
				txHash:       event.TxHash,
			}

			gatewayTokenIds[key] = matcherValue{
				tokenType:   types.TokenTypeERC1155,
				layer:       event.Layer,
				eventType:   event.Type,
				blockNumber: event.Number,
				messageHash: event.MessageHash,
				balance:     common.Big0,
			}

			if event.Type == types.L1DepositERC1155 || event.Type == types.L2WithdrawERC1155 ||
				event.Type == types.L1BatchDepositERC1155 || event.Type == types.L2BatchWithdrawERC1155 {
				gatewayTokenIds[key].balance.Add(gatewayTokenIds[key].balance, event.Amounts[idx])
			} else if event.Type == types.L2FinalizeDepositERC1155 || event.Type == types.L1FinalizeWithdrawERC1155 || event.Type == types.L1RefundERC1155 ||
				event.Type == types.L2FinalizeBatchDepositERC1155 || event.Type == types.L1FinalizeBatchWithdrawERC1155 || event.Type == types.L1BatchRefundERC1155 {
				gatewayTokenIds[key].balance.Sub(gatewayTokenIds[key].balance, event.Amounts[idx])
			}
		}
	}

	for _, event := range transferEvents {
		if len(event.TokenIds) != len(event.Amounts) {
			return fmt.Errorf("erc1155 transfer event tokenIds and amounts not match, %v", event)
		}

		for idx, tokenID := range event.TokenIds {
			key := erc1155MatchKey{
				tokenAddress: event.TokenAddress,
				tokenID:      tokenID,
				txHash:       event.TxHash,
			}
			// filter airdrop Transfers.
			_, exists := gatewayTokenIds[key]
			if !exists && event.Amounts[idx].Sign() >= 0 {
				continue
			}
			if _, exists := transferTokenIds[key]; !exists {
				transferTokenIds[key] = matcherValue{
					tokenType:   types.TokenTypeERC1155,
					layer:       event.Layer,
					eventType:   event.Type,
					blockNumber: event.Number,
					messageHash: event.MessageHash,
					balance:     common.Big0,
				}
			}
			transferTokenIds[key].balance.Add(transferTokenIds[key].balance, event.Amounts[idx])
		}
	}

	for transferMatcherKey, transferMatcherValue := range transferTokenIds {
		gatewayMatcherValue, exists := gatewayTokenIds[transferMatcherKey]
		// If the corresponding Transfer does not exist,
		// or if the tokens of Transfer events < the difference of tokens in gateway events (more tokens out).
		if !exists || transferMatcherValue.balance.Cmp(gatewayMatcherValue.balance) < 0 {
			info := slack.GatewayTransferInfo{
				TokenAddress:    transferMatcherKey.tokenAddress,
				TokenType:       transferMatcherValue.tokenType,
				Layer:           transferMatcherValue.layer,
				EventType:       transferMatcherValue.eventType,
				BlockNumber:     transferMatcherValue.blockNumber,
				TxHash:          transferMatcherKey.txHash,
				MessageHash:     transferMatcherValue.messageHash,
				TransferBalance: transferMatcherValue.balance,
			}
			if !exists {
				// Ignore additional Transfer events in Layer2.
				if info.Layer == types.Layer2 {
					continue
				}
				info.Error = transferEventDontHaveGatewayEvent
			} else {
				info.Error = transferEventBalanceMismatchGatewayEvent
				info.GatewayBalance = gatewayMatcherValue.balance
			}
			slack.Notify(slack.MrkDwnGatewayTransferMessage(info))
			return fmt.Errorf("erc1155 mismatch for tokenAddress %s: transfer amount = %s, gateway amount = %s",
				info.TokenAddress.Hex(), info.TransferBalance.String(), info.GatewayBalance.String())
		}
	}

	for gatewayMatcherKey, gatewayMatcherValue := range gatewayTokenIds {
		transferMatcherValue, exists := transferTokenIds[gatewayMatcherKey]
		// If the corresponding Transfer does not exist,
		// or if the tokens of Transfer events < the difference of tokens in gateway events (more tokens out).
		if !exists || gatewayMatcherValue.balance.Cmp(transferMatcherValue.balance) > 0 {
			info := slack.GatewayTransferInfo{
				TokenAddress:   gatewayMatcherKey.tokenAddress,
				TokenType:      gatewayMatcherValue.tokenType,
				Layer:          gatewayMatcherValue.layer,
				EventType:      gatewayMatcherValue.eventType,
				BlockNumber:    gatewayMatcherValue.blockNumber,
				TxHash:         gatewayMatcherKey.txHash,
				MessageHash:    gatewayMatcherValue.messageHash,
				GatewayBalance: gatewayMatcherValue.balance,
			}
			if !exists {
				info.Error = gatewayEventDontHaveTransferEvent
			} else {
				info.Error = gatewayEventBalanceMismatchTransferEvent
				info.TransferBalance = transferMatcherValue.balance
			}
			slack.Notify(slack.MrkDwnGatewayTransferMessage(info))
			return fmt.Errorf("erc1155 mismatch for token %s: gateway amount = %s, transfer amount = %s",
				info.TokenAddress.Hex(), info.GatewayBalance.String(), info.TransferBalance.String())
		}
	}

	return nil
}
