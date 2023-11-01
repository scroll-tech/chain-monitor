package checker

import "github.com/scroll-tech/chain-monitor/internal/logic/events"

// TransferEventMatcher check event exist and amount match
type TransferEventMatcher struct {
}

func NewTransferEventMatcher() *TransferEventMatcher {
	return &TransferEventMatcher{}
}

func (t *TransferEventMatcher) Erc20Matcher(transferEvents map[string]*events.ERC20GatewayEventUnmarshaler, gatewayEvent map[string]*events.ERC20GatewayEventUnmarshaler) error {
	var redundantTransferEvent []events.ERC20GatewayEventUnmarshaler
	for txHash, eventUnmarshaler := range transferEvents {

	}
}
