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

// Contracts is a struct that helps fetch and watch logs from Layer 1 and Layer 2 contracts.
type Contracts struct {
	l1Contracts *l1Contracts
	l2Contracts *l2Contracts
}

// NewContracts creates a new instance of Contracts which can be used to filter log fetchers
// from L1 and L2 smart contracts.
func NewContracts(l1Client, l2Client *ethclient.Client) *Contracts {
	c := &Contracts{
		l1Contracts: newL1Contracts(l1Client),
		l2Contracts: newL2Contracts(l2Client),
	}
	return c
}

// Register registers all gateway/messenger/transfer contracts present in the configuration.
func (l *Contracts) Register(conf config.Config) error {
	if err := l.l1Contracts.register(conf); err != nil {
		return err
	}

	if err := l.l2Contracts.register(conf); err != nil {
		return err
	}

	return nil
}

// Iterator returns a filter iterator for the provided layer type and transaction event category.
func (l *Contracts) Iterator(ctx context.Context, opts *bind.FilterOpts, layerType types.LayerType, txEventCategory types.TxEventCategory) ([]types.WrapIterator, error) {
	if layerType == types.Layer1 {
		switch txEventCategory {
		case types.ERC20EventCategory:
			return l.l1Erc20Filter(ctx, opts)
		case types.ERC721EventCategory:
			return l.l1Erc721Filter(ctx, opts)
		case types.ERC1155EventCategory:
			return l.l1Erc1155Filter(ctx, opts)
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
			return l.l2Erc1155Filter(ctx, opts)
		case types.MessengerEventCategory:
			return l.l2MessengerFilter(ctx, opts)
		}
	}

	return nil, fmt.Errorf("invalid type, layerType: %v, txEventCategory: %v", layerType, txEventCategory)
}

// GetGatewayTransfer returns a list of unmarshaled events for gateway transfers of a specific event category
// between the startBlockNumber and endBlockNumber. It returns an error if the layer type or transaction
// event category is invalid.
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
