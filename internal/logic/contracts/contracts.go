package contracts

import (
	"context"
	"fmt"

	"github.com/scroll-tech/go-ethereum/accounts/abi/bind"
	"github.com/scroll-tech/go-ethereum/ethclient"
	"github.com/scroll-tech/go-ethereum/log"

	"github.com/scroll-tech/chain-monitor/internal/config"
	"github.com/scroll-tech/chain-monitor/internal/types"
)

// Contracts fetch/watch the logs from l1/l2
type Contracts struct {
	l1Contracts *l1Contracts
	l2Contracts *l2Contracts
}

// NewContracts create contracts filter logs fetcher
func NewContracts(client *ethclient.Client) *Contracts {
	c := &Contracts{
		l1Contracts: newL1Contracts(client),
		l2Contracts: newL2Contracts(client),
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
func (l *Contracts) Iterator(ctx context.Context, opts *bind.FilterOpts, layerType types.LayerType, txEventCategory types.TxEventCategory) ([]WrapIterator, error) {
	if layerType == types.Layer1 {
		switch txEventCategory {
		case types.ERC20EventCategory:
			return l.l1Erc20Filter(ctx, opts)
		case types.ERC721EventCategory:
		case types.ERC1155EventCategory:
		}
	}

	if layerType == types.Layer2 {
		switch txEventCategory {
		case types.ERC20EventCategory:
			return l.l2Erc20Filter(ctx, opts)
		case types.ERC721EventCategory:
		case types.ERC1155EventCategory:
		}
	}

	return nil, nil
}

func (l *Contracts) l1Erc20Filter(ctx context.Context, opts *bind.FilterOpts) ([]WrapIterator, error) {
	var iterators []WrapIterator
	erc20TokenList := l.l1Contracts.ERC20CategoryTokens

	for _, erc20Token := range erc20TokenList {
		gatewayFilter, filterExist := l.l1Contracts.ERC20Gateways[erc20Token]
		if !filterExist {
			err := fmt.Errorf("can't get erc20 filter failed, erc20Token:%d", erc20Token)
			log.Error("get erc20 event filter from l1 contracts failed", "err", err)
			return nil, err
		}

		// deposit
		depositIter, err := gatewayFilter.FilterDepositERC20(opts, nil, nil, nil)
		if err != nil {
			log.Error("get erc20 gateway deposit iterator failed", "token type", erc20Token, "error", err)
			return nil, err
		}

		depositWrapIter := WrapIterator{
			Transfer:  false,
			Iter:      depositIter,
			EventType: types.L1DepositERC20,
		}
		iterators = append(iterators, depositWrapIter)

		// finalizeWithdraw
		finalizeWithdrawIter, err := gatewayFilter.FilterFinalizeWithdrawERC20(opts, nil, nil, nil)
		if err != nil {
			log.Error("get erc20 gateway finalizeWithdraw iterator failed", "token type", erc20Token, "error", err)
			return nil, err
		}

		finalizeWithdrawWrapIter := WrapIterator{
			Transfer:  false,
			Iter:      finalizeWithdrawIter,
			EventType: types.L1FinalizeWithdrawERC20,
		}
		iterators = append(iterators, finalizeWithdrawWrapIter)

		// refund
		refundIter, err := gatewayFilter.FilterRefundERC20(opts, nil, nil)
		if err != nil {
			log.Error("get erc20 gateway refund iterator failed", "token type", erc20Token, "error", err)
			return nil, err
		}

		refundWrapIter := WrapIterator{
			Transfer:  false,
			Iter:      refundIter,
			EventType: types.L1RefundERC20,
		}
		iterators = append(iterators, refundWrapIter)

		// transfer event
		transferFilter, transferFilerExist := l.l1Contracts.ERC20Transfer[erc20Token]
		if !transferFilerExist {
			err := fmt.Errorf("can't get erc20 transfer filter failed, erc20Token:%d", erc20Token)
			log.Error("get erc20 event transfer filter from l1 contracts failed", "err", err)
			return nil, err
		}

		transferIter, err := transferFilter.FilterTransfer(opts, nil, nil)
		if err != nil {
			log.Error("get erc20 gateway refund iterator failed", "token type", erc20Token, "error", err)
			return nil, err
		}
		transferWrapIter := WrapIterator{
			Transfer: true,
			Iter:     transferIter,
		}
		iterators = append(iterators, transferWrapIter)
	}
	return iterators, nil
}

func (l *Contracts) l2Erc20Filter(ctx context.Context, opts *bind.FilterOpts) ([]WrapIterator, error) {
	var iterators []WrapIterator
	erc20TokenList := l.l2Contracts.ERC20CategoryTokens

	for _, erc20Token := range erc20TokenList {
		gatewayFilter, filterExist := l.l2Contracts.ERC20Gateways[erc20Token]
		if !filterExist {
			err := fmt.Errorf("can't get erc20 filter failed, erc20Token:%d", erc20Token)
			log.Error("get erc20 event filter from l1 contracts failed", "err", err)
			return nil, err
		}

		// withdraw
		depositIter, err := gatewayFilter.FilterWithdrawERC20(opts, nil, nil, nil)
		if err != nil {
			log.Error("get erc20 gateway deposit iterator failed", "token type", erc20Token, "error", err)
			return nil, err
		}

		depositWrapIter := WrapIterator{
			Transfer:  false,
			Iter:      depositIter,
			EventType: types.L2WithdrawERC20,
		}
		iterators = append(iterators, depositWrapIter)

		// finalizeDeposit
		finalizeWithdrawIter, err := gatewayFilter.FilterFinalizeDepositERC20(opts, nil, nil, nil)
		if err != nil {
			log.Error("get erc20 gateway finalizeWithdraw iterator failed", "token type", erc20Token, "error", err)
			return nil, err
		}

		finalizeWithdrawWrapIter := WrapIterator{
			Transfer:  false,
			Iter:      finalizeWithdrawIter,
			EventType: types.L2FinalizeDepositERC20,
		}
		iterators = append(iterators, finalizeWithdrawWrapIter)

		// transfer event
		transferFilter, transferFilerExist := l.l2Contracts.ERC20Transfer[erc20Token]
		if !transferFilerExist {
			err := fmt.Errorf("can't get erc20 transfer filter failed, erc20Token:%d", erc20Token)
			log.Error("get erc20 event transfer filter from l1 contracts failed", "err", err)
			return nil, err
		}

		transferIter, err := transferFilter.FilterTransfer(opts, nil, nil)
		if err != nil {
			log.Error("get erc20 gateway refund iterator failed", "token type", erc20Token, "error", err)
			return nil, err
		}
		transferWrapIter := WrapIterator{
			Transfer: true,
			Iter:     transferIter,
		}
		iterators = append(iterators, transferWrapIter)
	}
	return iterators, nil
}
