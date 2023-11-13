package contracts

import (
	"context"
	"fmt"
	"math/big"

	"github.com/scroll-tech/chain-monitor/internal/logic/contracts/abi/iscrollerc20"
	"github.com/scroll-tech/chain-monitor/internal/logic/events"
	"github.com/scroll-tech/go-ethereum"
	"github.com/scroll-tech/go-ethereum/common"

	"github.com/scroll-tech/chain-monitor/internal/types"
	"github.com/scroll-tech/go-ethereum/accounts/abi/bind"
	"github.com/scroll-tech/go-ethereum/log"
)

func (l *Contracts) l1Erc20Filter(ctx context.Context, opts *bind.FilterOpts) ([]types.WrapIterator, error) {
	var iterators []types.WrapIterator
	erc20TokenList := l.l1Contracts.ERC20GatewayTokens
	for _, erc20Token := range erc20TokenList {
		gatewayFilter, filterExist := l.l1Contracts.ERC20Gateways[erc20Token.TokenType]
		if !filterExist {
			err := fmt.Errorf("can't get erc20 filter failed, erc20Token type:%v, address:%v", erc20Token.TokenType, erc20Token.Address)
			log.Error("get erc20 event filter from l1 contracts failed", "err", err)
			return nil, err
		}

		// deposit
		depositIter, err := gatewayFilter.FilterDepositERC20(opts, nil, nil, nil)
		if err != nil {
			log.Error("get erc20 gateway deposit iterator failed", "token type", erc20Token.TokenType, "address", erc20Token.Address, "error", err)
			return nil, err
		}

		depositWrapIter := types.WrapIterator{
			Iter:      depositIter,
			EventType: types.L1DepositERC20,
		}
		iterators = append(iterators, depositWrapIter)

		// finalizeWithdraw
		finalizeWithdrawIter, err := gatewayFilter.FilterFinalizeWithdrawERC20(opts, nil, nil, nil)
		if err != nil {
			log.Error("get erc20 gateway finalizeWithdraw iterator failed", "token type", erc20Token.TokenType, "address", erc20Token.Address, "error", err)
			return nil, err
		}

		finalizeWithdrawWrapIter := types.WrapIterator{
			Iter:      finalizeWithdrawIter,
			EventType: types.L1FinalizeWithdrawERC20,
		}
		iterators = append(iterators, finalizeWithdrawWrapIter)

		// refund
		refundIter, err := gatewayFilter.FilterRefundERC20(opts, nil, nil)
		if err != nil {
			log.Error("get erc20 gateway refund iterator failed", "token type", erc20Token.TokenType, "address", erc20Token.Address, "error", err)
			return nil, err
		}

		refundWrapIter := types.WrapIterator{
			Iter:      refundIter,
			EventType: types.L1RefundERC20,
		}
		iterators = append(iterators, refundWrapIter)
	}
	return iterators, nil
}

func (l *Contracts) l2Erc20Filter(ctx context.Context, opts *bind.FilterOpts) ([]types.WrapIterator, error) {
	var iterators []types.WrapIterator
	erc20TokenList := l.l2Contracts.ERC20GatewayTokens

	for _, erc20Token := range erc20TokenList {
		gatewayFilter, filterExist := l.l2Contracts.ERC20Gateways[erc20Token.TokenType]
		if !filterExist {
			err := fmt.Errorf("can't get erc20 filter failed, erc20Token type:%v, address:%v", erc20Token.TokenType, erc20Token.Address)
			log.Error("get erc20 event filter from l1 contracts failed", "err", err)
			return nil, err
		}

		// withdraw
		depositIter, err := gatewayFilter.FilterWithdrawERC20(opts, nil, nil, nil)
		if err != nil {
			log.Error("get erc20 gateway deposit iterator failed", "token type", erc20Token.TokenType, "address", erc20Token.Address, "error", err)
			return nil, err
		}

		depositWrapIter := types.WrapIterator{
			Iter:      depositIter,
			EventType: types.L2WithdrawERC20,
		}
		iterators = append(iterators, depositWrapIter)

		// finalizeDeposit
		finalizeWithdrawIter, err := gatewayFilter.FilterFinalizeDepositERC20(opts, nil, nil, nil)
		if err != nil {
			log.Error("get erc20 gateway finalizeWithdraw iterator failed", "token type", erc20Token.TokenType, "address", erc20Token.Address, "error", err)
			return nil, err
		}

		finalizeWithdrawWrapIter := types.WrapIterator{
			Iter:      finalizeWithdrawIter,
			EventType: types.L2FinalizeDepositERC20,
		}
		iterators = append(iterators, finalizeWithdrawWrapIter)
	}
	return iterators, nil
}

func (l *Contracts) getL1Erc20GatewayTransfer(ctx context.Context, startBlockNumber, endBlockNumber uint64) ([]events.EventUnmarshaler, error) {
	erc20ABI, err := iscrollerc20.Iscrollerc20MetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	transferEventQuery := ethereum.FilterQuery{
		FromBlock: new(big.Int).SetUint64(startBlockNumber),
		ToBlock:   new(big.Int).SetUint64(endBlockNumber),
		Topics:    [][]common.Hash{{erc20ABI.Events["Transfer"].ID}},
	}

	logs, err := l.l1Contracts.client.FilterLogs(ctx, transferEventQuery)
	if err != nil {
		return nil, err
	}

	tokenAddressMap := make(map[common.Address]struct{})
	for _, token := range l.l1Contracts.ERC20GatewayTokens {
		tokenAddressMap[token.Address] = struct{}{}
	}

	var transferEvents []events.EventUnmarshaler
	for _, vLog := range logs {
		event := struct {
			From  common.Address
			To    common.Address
			Value *big.Int
		}{}
		err := erc20ABI.UnpackIntoInterface(&event, "Transfer", vLog.Data)
		if err != nil {
			return nil, err
		}

		if _, ok := tokenAddressMap[event.From]; ok {
			transferEvents = append(transferEvents, &events.ERC20GatewayEventUnmarshaler{
				// @todo: add other fields.
				Amount:       new(big.Int).Neg(event.Value),
				TokenAddress: vLog.Address,
			})
		} else if _, ok := tokenAddressMap[event.To]; ok {
			transferEvents = append(transferEvents, &events.ERC20GatewayEventUnmarshaler{
				Amount:       event.Value,
				TokenAddress: vLog.Address,
			})
		}
	}

	return transferEvents, nil
}

func (l *Contracts) getL2Erc20GatewayTransfer(ctx context.Context, startBlockNumber, endBlockNumber uint64) ([]events.EventUnmarshaler, error) {
	erc20ABI, err := iscrollerc20.Iscrollerc20MetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	transferEventQuery := ethereum.FilterQuery{
		FromBlock: new(big.Int).SetUint64(startBlockNumber),
		ToBlock:   new(big.Int).SetUint64(endBlockNumber),
		Topics:    [][]common.Hash{{erc20ABI.Events["Transfer"].ID}},
	}

	logs, err := l.l2Contracts.client.FilterLogs(ctx, transferEventQuery)
	if err != nil {
		return nil, err
	}

	tokenAddressMap := make(map[common.Address]struct{})
	for _, token := range l.l2Contracts.ERC20GatewayTokens {
		tokenAddressMap[token.Address] = struct{}{}
	}

	var transferEvents []events.EventUnmarshaler
	for _, vLog := range logs {
		event := struct {
			From  common.Address
			To    common.Address
			Value *big.Int
		}{}
		err := erc20ABI.UnpackIntoInterface(&event, "Transfer", vLog.Data)
		if err != nil {
			return nil, err
		}

		if _, ok := tokenAddressMap[event.From]; ok {
			transferEvents = append(transferEvents, &events.ERC20GatewayEventUnmarshaler{
				// @todo: add other fields.
				Amount:       new(big.Int).Neg(event.Value),
				TokenAddress: vLog.Address,
			})
		}

		if _, ok := tokenAddressMap[event.To]; ok {
			transferEvents = append(transferEvents, &events.ERC20GatewayEventUnmarshaler{
				Amount:       event.Value,
				TokenAddress: vLog.Address,
			})
		}
	}

	return transferEvents, nil
}
