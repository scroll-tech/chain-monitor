package checker

import (
	"errors"

	"github.com/scroll-tech/go-ethereum/log"

	"github.com/scroll-tech/chain-monitor/internal/logic/events"
)

// TransferEventMatcher check event exist and amount match
type TransferEventMatcher struct {
}

func NewTransferEventMatcher() *TransferEventMatcher {
	return &TransferEventMatcher{}
}

// @todo: handle one tx with multiple events.
func (t *TransferEventMatcher) Erc20Matcher(transferEvents map[string]*events.ERC20GatewayEventUnmarshaler, gatewayEvents map[string]*events.ERC20GatewayEventUnmarshaler) error {
	var transferNotMatchGateway, gatewayNotMatchTransfer []*events.ERC20GatewayEventUnmarshaler
	var amountNotMatchList []*events.ERC20GatewayEventUnmarshaler
	for txHash, transferEventUnmarshaler := range transferEvents {
		gatewayEventUnmarshaler, exist := gatewayEvents[txHash]
		if !exist {
			transferNotMatchGateway = append(transferNotMatchGateway, transferEventUnmarshaler)
			log.Error("Erc20Matcher failed relate gateway event not exist", "event type", transferEventUnmarshaler.Type.String(), "tx_hash", txHash)
			continue
		}

		if gatewayEventUnmarshaler.Amount != transferEventUnmarshaler.Amount {
			amountNotMatchList = append(amountNotMatchList, gatewayEventUnmarshaler)
			log.Error("transfer value doesn't match gateway value", "event type", transferEventUnmarshaler.Type.String(), "tx_hash", txHash, "transfer event value",
				gatewayEventUnmarshaler.Amount.String(), "gateway event value", gatewayEventUnmarshaler.Amount.String())
			continue
		}
	}

	for txHash, gatewayEventUnmarshaler := range gatewayEvents {
		transferEventUnmarshaler, exist := transferEvents[txHash]
		if !exist {
			gatewayNotMatchTransfer = append(gatewayNotMatchTransfer, transferEventUnmarshaler)
			log.Error("Erc20Matcher failed relate transfer event not exist", "event type", gatewayEventUnmarshaler.Type.String(), "tx_hash", txHash)
			continue
		}
	}

	if len(transferNotMatchGateway) != 0 || len(gatewayNotMatchTransfer) != 0 || len(amountNotMatchList) != 0 {
		// todo -> send slack.
		return errors.New("erc20 match failed")
	}

	return nil
}
