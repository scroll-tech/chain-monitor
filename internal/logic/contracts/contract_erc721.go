package contracts

import (
	"context"
	"math/big"

	"github.com/scroll-tech/go-ethereum"
	"github.com/scroll-tech/go-ethereum/accounts/abi/bind"
	"github.com/scroll-tech/go-ethereum/common"
	"github.com/scroll-tech/go-ethereum/log"

	"github.com/scroll-tech/chain-monitor/internal/logic/contracts/abi/iscrollerc721"
	"github.com/scroll-tech/chain-monitor/internal/logic/events"
	"github.com/scroll-tech/chain-monitor/internal/types"
)

func (l *Contracts) l1Erc721Filter(_ context.Context, opts *bind.FilterOpts) ([]types.WrapIterator, error) {
	var iterators []types.WrapIterator

	// deposit
	depositIter, err := l.l1Contracts.erc721Gateway.FilterDepositERC721(opts, nil, nil, nil)
	if err != nil {
		log.Error("get erc721 gateway deposit iterator failed", "error", err)
		return nil, err
	}

	depositWrapIter := types.WrapIterator{
		Iter:      depositIter,
		EventType: types.L1DepositERC721,
	}
	iterators = append(iterators, depositWrapIter)

	// batch deposit
	batchDepositIter, err := l.l1Contracts.erc721Gateway.FilterBatchDepositERC721(opts, nil, nil, nil)
	if err != nil {
		log.Error("get erc721 gateway batch deposit iterator failed", "error", err)
		return nil, err
	}

	batchDepositWrapIter := types.WrapIterator{
		Iter:      batchDepositIter,
		EventType: types.L1BatchDepositERC721,
	}
	iterators = append(iterators, batchDepositWrapIter)

	// finalizeWithdraw
	finalizeWithdrawIter, err := l.l1Contracts.erc721Gateway.FilterFinalizeWithdrawERC721(opts, nil, nil, nil)
	if err != nil {
		log.Error("get erc721 gateway finalizeWithdraw iterator failed", "error", err)
		return nil, err
	}

	finalizeWithdrawWrapIter := types.WrapIterator{
		Iter:      finalizeWithdrawIter,
		EventType: types.L1FinalizeWithdrawERC721,
	}
	iterators = append(iterators, finalizeWithdrawWrapIter)

	// finalize batch Withdraw
	finalizeBatchWithdrawIter, err := l.l1Contracts.erc721Gateway.FilterFinalizeBatchWithdrawERC721(opts, nil, nil, nil)
	if err != nil {
		log.Error("get erc721 gateway finalizeWithdraw iterator failed", "error", err)
		return nil, err
	}

	finalizeBatchWithdrawWrapIter := types.WrapIterator{
		Iter:      finalizeBatchWithdrawIter,
		EventType: types.L1FinalizeBatchWithdrawERC721,
	}
	iterators = append(iterators, finalizeBatchWithdrawWrapIter)

	// refund
	refundIter, err := l.l1Contracts.erc721Gateway.FilterRefundERC721(opts, nil, nil)
	if err != nil {
		log.Error("get erc721 gateway refund iterator failed", "error", err)
		return nil, err
	}

	refundWrapIter := types.WrapIterator{
		Iter:      refundIter,
		EventType: types.L1RefundERC721,
	}
	iterators = append(iterators, refundWrapIter)

	// batch refund
	batchRefundIter, err := l.l1Contracts.erc721Gateway.FilterBatchRefundERC721(opts, nil, nil)
	if err != nil {
		log.Error("get erc721 gateway refund iterator failed", "error", err)
		return nil, err
	}

	batchRefundWrapIter := types.WrapIterator{
		Iter:      batchRefundIter,
		EventType: types.L1BatchRefundERC721,
	}
	iterators = append(iterators, batchRefundWrapIter)

	return iterators, nil
}

func (l *Contracts) l2Erc721Filter(_ context.Context, opts *bind.FilterOpts) ([]types.WrapIterator, error) {
	var iterators []types.WrapIterator

	// withdraw
	withdrawIter, err := l.l2Contracts.erc721Gateway.FilterWithdrawERC721(opts, nil, nil, nil)
	if err != nil {
		log.Error("get erc721 gateway withdraw iterator failed", "error", err)
		return nil, err
	}

	withdrawWrapIter := types.WrapIterator{
		Iter:      withdrawIter,
		EventType: types.L2WithdrawERC721,
	}
	iterators = append(iterators, withdrawWrapIter)

	// batch withdraw
	batchWithdrawIter, err := l.l2Contracts.erc721Gateway.FilterBatchWithdrawERC721(opts, nil, nil, nil)
	if err != nil {
		log.Error("get erc721 gateway batch withdraw iterator failed", "error", err)
		return nil, err
	}

	batchWithdrawWrapIter := types.WrapIterator{
		Iter:      batchWithdrawIter,
		EventType: types.L2BatchWithdrawERC721,
	}
	iterators = append(iterators, batchWithdrawWrapIter)

	// finalizeDeposit
	finalizeDepositIter, err := l.l2Contracts.erc721Gateway.FilterFinalizeDepositERC721(opts, nil, nil, nil)
	if err != nil {
		log.Error("get erc721 gateway finalize deposit iterator failed", "error", err)
		return nil, err
	}

	finalizeWithdrawWrapIter := types.WrapIterator{
		Iter:      finalizeDepositIter,
		EventType: types.L2FinalizeDepositERC721,
	}
	iterators = append(iterators, finalizeWithdrawWrapIter)

	// batch finalize Deposit
	batchFinalizeDepositIter, err := l.l2Contracts.erc721Gateway.FilterFinalizeBatchDepositERC721(opts, nil, nil, nil)
	if err != nil {
		log.Error("get erc721 gateway finalize batch deposit iterator failed", "error", err)
		return nil, err
	}

	batchFinalizeWithdrawWrapIter := types.WrapIterator{
		Iter:      batchFinalizeDepositIter,
		EventType: types.L2FinalizeBatchDepositERC721,
	}
	iterators = append(iterators, batchFinalizeWithdrawWrapIter)

	return iterators, nil
}

func (l *Contracts) getL1Erc721GatewayTransfer(ctx context.Context, startBlockNumber, endBlockNumber uint64) ([]events.EventUnmarshaler, error) {
	erc721ABI, err := iscrollerc721.Iscrollerc721MetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	transferEventQuery := ethereum.FilterQuery{
		FromBlock: new(big.Int).SetUint64(startBlockNumber),
		ToBlock:   new(big.Int).SetUint64(endBlockNumber),
		Topics:    [][]common.Hash{{erc721ABI.Events["Transfer"].ID}},
	}

	logs, err := l.l1Contracts.client.FilterLogs(ctx, transferEventQuery)
	if err != nil {
		return nil, err
	}

	var transferEvents []events.EventUnmarshaler
	for _, vLog := range logs {
		event := struct {
			From    common.Address
			To      common.Address
			TokenID *big.Int
		}{}
		if unpackErr := erc721ABI.UnpackIntoInterface(&event, "Transfer", vLog.Data); unpackErr != nil {
			return nil, err
		}

		if l.l1Contracts.erc721GatewayAddress == event.From {
			transferEvents = append(transferEvents, &events.ERC721GatewayEventUnmarshaler{
				TokenIds:     []*big.Int{event.TokenID},
				Amounts:      []*big.Int{new(big.Int).Neg(big.NewInt(1))},
				TokenAddress: vLog.Address,
				TxHash:       vLog.TxHash,
				Number:       vLog.BlockNumber,
				Layer:        types.Layer1,
			})
		}

		if l.l1Contracts.erc721GatewayAddress == event.To {
			transferEvents = append(transferEvents, &events.ERC721GatewayEventUnmarshaler{
				TokenIds:     []*big.Int{event.TokenID},
				Amounts:      []*big.Int{big.NewInt(1)},
				TokenAddress: vLog.Address,
				TxHash:       vLog.TxHash,
				Number:       vLog.BlockNumber,
				Layer:        types.Layer1,
			})
		}
	}

	return transferEvents, nil
}

func (l *Contracts) getL2Erc721GatewayTransfer(ctx context.Context, startBlockNumber, endBlockNumber uint64) ([]events.EventUnmarshaler, error) {
	erc721ABI, err := iscrollerc721.Iscrollerc721MetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	transferEventQuery := ethereum.FilterQuery{
		FromBlock: new(big.Int).SetUint64(startBlockNumber),
		ToBlock:   new(big.Int).SetUint64(endBlockNumber),
		Topics:    [][]common.Hash{{erc721ABI.Events["Transfer"].ID}},
	}

	logs, err := l.l2Contracts.client.FilterLogs(ctx, transferEventQuery)
	if err != nil {
		return nil, err
	}

	var transferEvents []events.EventUnmarshaler
	for _, vLog := range logs {
		event := struct {
			From    common.Address
			To      common.Address
			TokenID *big.Int
		}{}
		if err := erc721ABI.UnpackIntoInterface(&event, "Transfer", vLog.Data); err != nil {
			return nil, err
		}

		if event.From == l.l2Contracts.erc721GatewayAddress {
			transferEvents = append(transferEvents, &events.ERC721GatewayEventUnmarshaler{
				TokenIds:     []*big.Int{event.TokenID},
				Amounts:      []*big.Int{new(big.Int).Neg(big.NewInt(1))},
				TokenAddress: vLog.Address,
				TxHash:       vLog.TxHash,
				Number:       vLog.BlockNumber,
				Layer:        types.Layer2,
			})
		}

		if event.To == l.l2Contracts.erc721GatewayAddress {
			transferEvents = append(transferEvents, &events.ERC721GatewayEventUnmarshaler{
				TokenIds:     []*big.Int{event.TokenID},
				Amounts:      []*big.Int{big.NewInt(1)},
				TokenAddress: vLog.Address,
				TxHash:       vLog.TxHash,
				Number:       vLog.BlockNumber,
				Layer:        types.Layer2,
			})
		}
	}

	return transferEvents, nil
}
