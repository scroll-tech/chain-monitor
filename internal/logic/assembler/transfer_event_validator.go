package assembler

import (
	"fmt"
	"io"
	"math/big"
	"net/http"
	"time"

	"github.com/scroll-tech/go-ethereum/common"
	"github.com/scroll-tech/go-ethereum/ethclient"
	"github.com/scroll-tech/go-ethereum/log"

	"github.com/scroll-tech/chain-monitor/internal/config"
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
type TransferEventMatcher struct {
	l1IgnoredTokens map[common.Address]struct{}
	l2IgnoredTokens map[common.Address]struct{}

	httpClient *http.Client
}

// NewTransferEventMatcher creates a new instance of TransferEventMatcher.
func NewTransferEventMatcher(conf *config.Config, l1Client, l2Client *ethclient.Client) *TransferEventMatcher {
	t := &TransferEventMatcher{
		l1IgnoredTokens: make(map[common.Address]struct{}),
		l2IgnoredTokens: make(map[common.Address]struct{}),
		httpClient: &http.Client{
			Timeout: 10 * time.Second,
		},
	}

	for _, token := range conf.L1Config.IgnoredTokens {
		t.l1IgnoredTokens[token] = struct{}{}
	}

	for _, token := range conf.L2Config.IgnoredTokens {
		t.l2IgnoredTokens[token] = struct{}{}
	}

	// Log the ignored tokens
	log.Info("Ignored Tokens", "L1", t.l1IgnoredTokens, "L2", t.l2IgnoredTokens)

	return t
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
				balance:     big.NewInt(0),
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
				balance:     big.NewInt(0),
			}
		}
		transferBalances[key].balance.Add(transferBalances[key].balance, event.Amount)
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
			log.Error("erc20 gateway and transfer not match",
				"token address", info.TokenAddress.Hex(),
				"block number", info.BlockNumber,
				"transfer balance", info.TransferBalance.String(),
				"gateway balance", info.GatewayBalance.String(),
				"err info", info.Error,
			)
			return t.sendSlackAlert(info)
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
			log.Error("erc20 gateway and transfer not match",
				"token address", info.TokenAddress.Hex(),
				"block number", info.BlockNumber,
				"transfer balance", info.TransferBalance.String(),
				"gateway balance", info.GatewayBalance.String(),
				"err info", info.Error,
			)
			return t.sendSlackAlert(info)
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
					balance:     big.NewInt(0),
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
					balance:     big.NewInt(0),
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
			log.Error("erc721 gateway and transfer not match",
				"token address", info.TokenAddress.Hex(),
				"block number", info.BlockNumber,
				"transfer balance", info.TransferBalance.String(),
				"gateway balance", info.GatewayBalance.String(),
				"err info", info.Error,
			)
			return t.sendSlackAlert(info)
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
			log.Error("erc721 gateway and transfer not match",
				"token address", info.TokenAddress.Hex(),
				"block number", info.BlockNumber,
				"transfer balance", info.TransferBalance.String(),
				"gateway balance", info.GatewayBalance.String(),
				"err info", info.Error,
			)
			return t.sendSlackAlert(info)
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
				balance:     big.NewInt(0),
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
					balance:     big.NewInt(0),
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
			log.Error("erc1155 gateway and transfer not match",
				"token address", info.TokenAddress.Hex(),
				"block number", info.BlockNumber,
				"transfer balance", info.TransferBalance.String(),
				"gateway balance", info.GatewayBalance.String(),
				"err info", info.Error,
			)
			return t.sendSlackAlert(info)
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
			log.Error("erc1155 gateway and transfer not match",
				"token address", info.TokenAddress.Hex(),
				"block number", info.BlockNumber,
				"transfer balance", info.TransferBalance.String(),
				"gateway balance", info.GatewayBalance.String(),
				"err info", info.Error,
			)
			return t.sendSlackAlert(info)
		}
	}

	return nil
}

func (t *TransferEventMatcher) sendSlackAlert(info slack.GatewayTransferInfo) error {
	info.TokenIgnored = t.isTokenIgnored(info.Layer, info.TokenAddress)

	// If it's an L1 token, check CoinGecko
	if info.Layer == types.Layer1 && info.TokenType == types.TokenTypeERC20 {
		tokenExists, err := t.checkTokenInCoinGecko(info.TokenAddress)
		if err != nil {
			// API failure or network error, use conservative strategy: send alert
			log.Error("CoinGecko API check failed, sending alert with conservative strategy",
				"token", info.TokenAddress.Hex(),
				"error", err.Error())

			info.CoinGeckoStatus = fmt.Sprintf("API check failed: %s", err.Error())
		} else if tokenExists {
			// Token exists in CoinGecko -> send alert
			log.Info("L1 token found in CoinGecko - sending alert",
				"token", info.TokenAddress.Hex(),
				"blockNumber", info.BlockNumber)
			info.CoinGeckoStatus = "found in CoinGecko"
		} else {
			// Token explicitly not in CoinGecko -> send notification only
			log.Info("L1 token not found in CoinGecko - sending notification only",
				"token", info.TokenAddress.Hex(),
				"blockNumber", info.BlockNumber)

			notificationMsg := fmt.Sprintf("ℹ️ L1 Token not found in CoinGecko\n"+
				"Token: %s\n"+
				"Block: %d\n"+
				"Transfer Balance: %s\n"+
				"Gateway Balance: %s\n"+
				"Note: This token is not indexed by CoinGecko, may be a non-mainstream token",
				info.TokenAddress.Hex(),
				info.BlockNumber,
				info.TransferBalance.String(),
				info.GatewayBalance.String())

			slack.Notify(notificationMsg)
			return nil // Don't treat this as an error
		}
	}

	slack.Notify(slack.MrkDwnGatewayTransferMessage(info))
	if info.TokenIgnored {
		return nil
	}

	time.Sleep(1 * time.Minute) // Sleep for 1 minute to avoid spamming Slack with alerts

	return fmt.Errorf("balance mismatch for token %s, token type = %s, transfer amount = %s, gateway amount = %s, coingecko status = %s, info = %+v",
		info.TokenAddress.Hex(), info.TokenType.String(), info.TransferBalance.String(), info.GatewayBalance.String(), info.CoinGeckoStatus, info)
}

// checkTokenInCoinGecko checks if a token exists in CoinGecko database
// checkTokenInCoinGecko checks if a token exists in CoinGecko database
func (t *TransferEventMatcher) checkTokenInCoinGecko(tokenAddress common.Address) (bool, error) {
	url := fmt.Sprintf("https://api.coingecko.com/api/v3/coins/ethereum/contract/%s", tokenAddress.Hex())

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return false, fmt.Errorf("failed to create request: %w", err)
	}

	req.Header.Set("Accept", "application/json")
	req.Header.Set("User-Agent", "chain-monitor/1.0")

	resp, err := t.httpClient.Do(req)
	if err != nil {
		return false, fmt.Errorf("get token info from CoinGecko failed: url %s, error: %w", url, err)
	}
	defer func() {
		if closeErr := resp.Body.Close(); closeErr != nil {
			log.Warn("failed to close response body", "error", closeErr)
		}
	}()

	// Use status code to determine token existence
	switch resp.StatusCode {
	case 200:
		// Token exists in CoinGecko
		return true, nil
	case 404:
		// Token not found in CoinGecko
		return false, nil
	default:
		// Other errors (API limits, server errors, etc.)
		body, readErr := io.ReadAll(resp.Body)
		if readErr != nil {
			return false, fmt.Errorf("CoinGecko API error (status %d), failed to read response body: %w", resp.StatusCode, readErr)
		}
		return false, fmt.Errorf("CoinGecko API error (status %d): %s", resp.StatusCode, string(body))
	}
}

func (t *TransferEventMatcher) isTokenIgnored(layer types.LayerType, tokenAddress common.Address) bool {
	if layer == types.Layer2 {
		if _, ok := t.l2IgnoredTokens[tokenAddress]; ok {
			log.Warn("l2 token is ignored", "token address", tokenAddress.Hex())
			return true
		}
	} else {
		if _, ok := t.l1IgnoredTokens[tokenAddress]; ok {
			log.Warn("l1 token is ignored", "token address", tokenAddress.Hex())
			return true
		}
	}
	return false
}
