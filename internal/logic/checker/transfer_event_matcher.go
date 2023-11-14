package checker

import (
	"fmt"
	"math/big"

	"github.com/scroll-tech/go-ethereum/common"

	"github.com/scroll-tech/chain-monitor/internal/logic/events"
	"github.com/scroll-tech/chain-monitor/internal/types"
)

// TransferEventMatcher check event exist and amount match
type TransferEventMatcher struct {
}

func NewTransferEventMatcher() *TransferEventMatcher {
	return &TransferEventMatcher{}
}

func (t *TransferEventMatcher) Erc20Matcher(transferEvents, gatewayEvents []events.ERC20GatewayEventUnmarshaler) error {
	transferBalances := make(map[common.Address]*big.Int)
	gatewayBalances := make(map[common.Address]*big.Int)

	for _, event := range transferEvents {
		if _, exists := transferBalances[event.TokenAddress]; !exists {
			transferBalances[event.TokenAddress] = new(big.Int)
		}
		transferBalances[event.TokenAddress].Add(transferBalances[event.TokenAddress], event.Amount)
	}

	for _, event := range gatewayEvents {
		if _, exists := gatewayBalances[event.TokenAddress]; !exists {
			gatewayBalances[event.TokenAddress] = new(big.Int)
		}
		if event.Type == types.L1DepositERC20 || event.Type == types.L2WithdrawERC20 {
			gatewayBalances[event.TokenAddress].Add(gatewayBalances[event.TokenAddress], event.Amount)
		} else if event.Type == types.L2FinalizeDepositERC20 || event.Type == types.L1FinalizeWithdrawERC20 || event.Type == types.L1RefundERC20 {
			gatewayBalances[event.TokenAddress].Sub(gatewayBalances[event.TokenAddress], event.Amount)
		}
	}

	for tokenAddress, transferBalance := range transferBalances {
		gatewayBalance, exists := gatewayBalances[tokenAddress]
		if !exists || transferBalance.Cmp(gatewayBalance) != 0 {
			// send slack.
			return fmt.Errorf("balance mismatch for token %s: transfer balance = %s, gateway balance = %s",
				tokenAddress.Hex(), transferBalance.String(), gatewayBalance.String())
		}
	}

	for tokenAddress, gatewayBalance := range gatewayBalances {
		transferBalance, exists := transferBalances[tokenAddress]
		if !exists || gatewayBalance.Cmp(transferBalance) != 0 {
			// send slack.
			return fmt.Errorf("balance mismatch for token %s: gateway balance = %s, transfer balance = %s",
				tokenAddress.Hex(), gatewayBalance.String(), transferBalance.String())
		}
	}

	return nil
}

func (t *TransferEventMatcher) Erc721Matcher(transferEvents, gatewayEvents []events.ERC721GatewayEventUnmarshaler) error {
	type erc721MatchKey struct {
		tokenAddress common.Address
		tokenID      *big.Int
	}
	transferTokenIds := make(map[erc721MatchKey]*big.Int)
	gatewayTokenIds := make(map[erc721MatchKey]*big.Int)

	for _, event := range transferEvents {
		if len(event.TokenIds) != len(event.Amounts) {
			return fmt.Errorf("erc721 transfer event tokenIds and amounts not match, %v", event)
		}

		for idx, tokenID := range event.TokenIds {
			key := erc721MatchKey{
				tokenAddress: event.TokenAddress,
				tokenID:      tokenID,
			}
			if _, exists := transferTokenIds[key]; !exists {
				transferTokenIds[key] = new(big.Int)
			}
			transferTokenIds[key].Add(transferTokenIds[key], event.Amounts[idx])
		}
	}

	for _, event := range gatewayEvents {
		if len(event.TokenIds) != len(event.Amounts) {
			return fmt.Errorf("erc1155 gateway event tokenIds and amounts not match, %v", event)
		}

		for idx, tokenID := range event.TokenIds {
			key := erc721MatchKey{
				tokenAddress: event.TokenAddress,
				tokenID:      tokenID,
			}

			if event.Type == types.L1DepositERC721 || event.Type == types.L2WithdrawERC721 ||
				event.Type == types.L1BatchDepositERC721 || event.Type == types.L2BatchWithdrawERC721 {
				gatewayTokenIds[key].Add(gatewayTokenIds[key], event.Amounts[idx])
			} else if event.Type == types.L2FinalizeDepositERC721 || event.Type == types.L1FinalizeWithdrawERC721 || event.Type == types.L1RefundERC721 ||
				event.Type == types.L2FinalizeBatchDepositERC721 || event.Type == types.L1FinalizeBatchWithdrawERC721 || event.Type == types.L1BatchRefundERC721 {
				gatewayTokenIds[key].Sub(gatewayTokenIds[key], event.Amounts[idx])
			}
		}
	}

	for key, transferAmount := range transferTokenIds {
		gatewayAmount, exists := gatewayTokenIds[key]
		if !exists || transferAmount.Cmp(gatewayAmount) != 0 {
			// send slack.
			return fmt.Errorf("erc721 mismatch for tokenAddress %s: transfer amount = %s, gateway amount = %s",
				key.tokenAddress.Hex(), transferAmount.String(), gatewayAmount.String())
		}
	}

	for key, gatewayAmount := range gatewayTokenIds {
		transferAmount, exists := transferTokenIds[key]
		if !exists || gatewayAmount.Cmp(transferAmount) != 0 {
			// send slack.
			return fmt.Errorf("erc721 mismatch for tokenAddress %s: gateway amount = %s, transfer amount = %s",
				key.tokenAddress.Hex(), gatewayAmount.String(), transferAmount.String())
		}
	}

	return nil
}

func (t *TransferEventMatcher) Erc1155Matcher(transferEvents, gatewayEvents []events.ERC1155GatewayEventUnmarshaler) error {
	type erc1155MatchKey struct {
		tokenAddress common.Address
		tokenID      *big.Int
	}

	transferTokenIds := make(map[erc1155MatchKey]*big.Int)
	gatewayTokenIds := make(map[erc1155MatchKey]*big.Int)

	for _, event := range transferEvents {
		if len(event.TokenIds) != len(event.Amounts) {
			return fmt.Errorf("erc1155 transfer event tokenIds and amounts not match, %v", event)
		}

		for idx, tokenID := range event.TokenIds {
			key := erc1155MatchKey{
				tokenAddress: event.TokenAddress,
				tokenID:      tokenID,
			}
			if _, exists := transferTokenIds[key]; !exists {
				transferTokenIds[key] = new(big.Int)
			}
			transferTokenIds[key].Add(transferTokenIds[key], event.Amounts[idx])
		}
	}

	for _, event := range gatewayEvents {
		if len(event.TokenIds) != len(event.Amounts) {
			return fmt.Errorf("erc1155 gateway event tokenIds and amounts not match, %v", event)
		}

		for idx, tokenID := range event.TokenIds {
			key := erc1155MatchKey{
				tokenAddress: event.TokenAddress,
				tokenID:      tokenID,
			}

			if event.Type == types.L1DepositERC1155 || event.Type == types.L2WithdrawERC1155 ||
				event.Type == types.L1BatchDepositERC1155 || event.Type == types.L2BatchWithdrawERC1155 {
				gatewayTokenIds[key].Add(gatewayTokenIds[key], event.Amounts[idx])
			} else if event.Type == types.L2FinalizeDepositERC1155 || event.Type == types.L1FinalizeWithdrawERC1155 || event.Type == types.L1RefundERC1155 ||
				event.Type == types.L2FinalizeBatchDepositERC1155 || event.Type == types.L1FinalizeBatchWithdrawERC1155 || event.Type == types.L1BatchRefundERC1155 {
				gatewayTokenIds[key].Sub(gatewayTokenIds[key], event.Amounts[idx])
			}
		}
	}

	for key, transferAmount := range transferTokenIds {
		gatewayAmount, exists := gatewayTokenIds[key]
		if !exists || transferAmount.Cmp(gatewayAmount) != 0 {
			// send slack.
			return fmt.Errorf("erc1155 mismatch for tokenAddress %s: transfer amount = %s, gateway amount = %s",
				key.tokenAddress.Hex(), transferAmount.String(), gatewayAmount.String())
		}
	}

	for key, gatewayAmount := range gatewayTokenIds {
		transferAmount, exists := transferTokenIds[key]
		if !exists || gatewayAmount.Cmp(transferAmount) != 0 {
			// send slack.
			return fmt.Errorf("erc1155 mismatch for token %s: gateway amount = %s, transfer amount = %s",
				key.tokenAddress.Hex(), gatewayAmount.String(), transferAmount.String())
		}
	}

	return nil
}
