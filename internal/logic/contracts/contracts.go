package contracts

import (
	"context"
	"fmt"

	"github.com/scroll-tech/go-ethereum/accounts/abi/bind"
	"github.com/scroll-tech/go-ethereum/ethclient"

	"github.com/scroll-tech/chain-monitor/internal/config"
	"github.com/scroll-tech/chain-monitor/internal/logic/events"
	"github.com/scroll-tech/chain-monitor/internal/types"
)

// Contracts fetch/watch the logs from l1/l2
type Contracts struct {
	l1Contracts *l1Contracts
	l2Contracts *l2Contracts
}

// NewContracts create contracts filter logs fetcher
func NewContracts(l1Client, l2Client *ethclient.Client) *Contracts {
	c := &Contracts{
		l1Contracts: newL1Contracts(l1Client),
		l2Contracts: newL2Contracts(l2Client),
	}
	return c
}

// Register all gateway/messenger/transfer contracts
func (l *Contracts) Register(conf config.Config) error {
	if err := l.l1Contracts.register(conf); err != nil {
		return err
	}

	if err := l.l2Contracts.register(conf); err != nil {
		return err
	}

	return nil
}

// Iterator get the filter iterator
func (l *Contracts) Iterator(ctx context.Context, opts *bind.FilterOpts, layerType types.LayerType, txEventCategory types.TxEventCategory) ([]types.WrapIterator, error) {
	if layerType == types.Layer1 {
		switch txEventCategory {
		case types.ERC20EventCategory:
			return l.l1Erc20Filter(ctx, opts)
		case types.ERC721EventCategory:
			return l.l1Erc721Filter(ctx, opts)
		case types.ERC1155EventCategory:
		case types.MessengerEventCategory:
			return l.l1MessengerFilter(ctx, opts)
		}
	}

	if layerType == types.Layer2 {
		switch txEventCategory {
		case types.ERC20EventCategory:
			return l.l2Erc20Filter(ctx, opts)
		case types.ERC721EventCategory:
			return l.l2Erc721Filter(ctx, opts)
		case types.ERC1155EventCategory:
		case types.MessengerEventCategory:
			return l.l2MessengerFilter(ctx, opts)
		}
	}

	return nil, fmt.Errorf("invalid type, layerType: %v, txEventCategory: %v", layerType, txEventCategory)
}

func (l *Contracts) GetGatewayTransfer(ctx context.Context, startBlockNumber, endBlockNumber uint64, layerType types.LayerType, txEventCategory types.TxEventCategory) ([]events.EventUnmarshaler, error) {
	if layerType == types.Layer1 {
		switch txEventCategory {
		case types.ERC20EventCategory:
			return l.getL1Erc20GatewayTransfer(ctx, startBlockNumber, endBlockNumber)
		case types.ERC721EventCategory:
			return l.getL1Erc721GatewayTransfer(ctx, startBlockNumber, endBlockNumber)
		case types.ERC1155EventCategory:
			return l.getL1Erc1155GatewayTransfer(ctx, startBlockNumber, endBlockNumber)
		}
	}

	if layerType == types.Layer2 {
		switch txEventCategory {
		case types.ERC20EventCategory:
			return l.getL2Erc20GatewayTransfer(ctx, startBlockNumber, endBlockNumber)
		case types.ERC721EventCategory:
			return l.getL2Erc721GatewayTransfer(ctx, startBlockNumber, endBlockNumber)
		case types.ERC1155EventCategory:
			return l.getL2Erc1155GatewayTransfer(ctx, startBlockNumber, endBlockNumber)
		}
	}

	return nil, fmt.Errorf("invalid type, layerType: %v, txEventCategory: %v", layerType, txEventCategory)
}
