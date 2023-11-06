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
