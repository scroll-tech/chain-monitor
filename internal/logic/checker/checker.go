package checker

import (
	"github.com/scroll-tech/chain-monitor/internal/logic/events"
	"github.com/scroll-tech/chain-monitor/internal/types"
)

type Checker struct {
	transferMatcher *TransferEventMatcher
}

func NewChecker() *Checker {
	return &Checker{
		transferMatcher: NewTransferEventMatcher(),
	}
}

func (c *Checker) Check(eventCategory types.TxEventCategory, eventDataList []events.EventUnmarshaler) error {
	switch eventCategory {
	case types.ERC20EventCategory:
		return c.erc20EventUnmarshaler(eventDataList)
	}
	return nil
}

func (c *Checker) erc20EventUnmarshaler(eventDataList []events.EventUnmarshaler) error {
	erc20TransferEventMap := make(map[string]*events.ERC20GatewayEventUnmarshaler)
	erc20GatewayEventMap := make(map[string]*events.ERC20GatewayEventUnmarshaler)
	for _, eventData := range eventDataList {
		erc20EventUnmarshaler := eventData.(*events.ERC20GatewayEventUnmarshaler)
		if erc20EventUnmarshaler.Transfer {
			erc20TransferEventMap[erc20EventUnmarshaler.TxHash] = erc20EventUnmarshaler
		} else {
			erc20GatewayEventMap[erc20EventUnmarshaler.TxHash] = erc20EventUnmarshaler
		}
	}

	return c.transferMatcher.Erc20Matcher(erc20TransferEventMap, erc20GatewayEventMap)
}
