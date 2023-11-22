package contracts

import (
	"context"
	"math/big"

	"github.com/scroll-tech/go-ethereum"
	"github.com/scroll-tech/go-ethereum/accounts/abi/bind"
	"github.com/scroll-tech/go-ethereum/common"
	"github.com/scroll-tech/go-ethereum/log"

	"github.com/scroll-tech/chain-monitor/internal/logic/contracts/abi/iscrollerc1155"
	"github.com/scroll-tech/chain-monitor/internal/logic/events"
	"github.com/scroll-tech/chain-monitor/internal/types"
)

func (l *Contracts) l1Erc1155Filter(_ context.Context, opts *bind.FilterOpts) ([]types.WrapIterator, error) {
	var iterators []types.WrapIterator

	// deposit
	depositIter, err := l.l1Contracts.ERC1155Gateway.FilterDepositERC1155(opts, nil, nil, nil)
	if err != nil {
		log.Error("get erc1155 gateway deposit iterator failed", "error", err)
		return nil, err
	}

	depositWrapIter := types.WrapIterator{
		Iter:      depositIter,
		EventType: types.L1DepositERC1155,
	}
	iterators = append(iterators, depositWrapIter)

	// batch deposit
	batchDepositIter, err := l.l1Contracts.ERC1155Gateway.FilterBatchDepositERC1155(opts, nil, nil, nil)
	if err != nil {
		log.Error("get erc1155 gateway batch deposit iterator failed", "error", err)
		return nil, err
	}

	batchDepositWrapIter := types.WrapIterator{
		Iter:      batchDepositIter,
		EventType: types.L1BatchDepositERC1155,
	}
	iterators = append(iterators, batchDepositWrapIter)

	// finalizeWithdraw
	finalizeWithdrawIter, err := l.l1Contracts.ERC1155Gateway.FilterFinalizeWithdrawERC1155(opts, nil, nil, nil)
	if err != nil {
		log.Error("get erc1155 gateway finalizeWithdraw iterator failed", "error", err)
		return nil, err
	}

	finalizeWithdrawWrapIter := types.WrapIterator{
		Iter:      finalizeWithdrawIter,
		EventType: types.L1FinalizeWithdrawERC1155,
	}
	iterators = append(iterators, finalizeWithdrawWrapIter)

	// finalize batch Withdraw
	finalizeBatchWithdrawIter, err := l.l1Contracts.ERC1155Gateway.FilterFinalizeBatchWithdrawERC1155(opts, nil, nil, nil)
	if err != nil {
		log.Error("get erc1155 gateway finalizeWithdraw iterator failed", "error", err)
		return nil, err
	}

	finalizeBatchWithdrawWrapIter := types.WrapIterator{
		Iter:      finalizeBatchWithdrawIter,
		EventType: types.L1FinalizeBatchWithdrawERC1155,
	}
	iterators = append(iterators, finalizeBatchWithdrawWrapIter)

	// refund
	refundIter, err := l.l1Contracts.ERC1155Gateway.FilterRefundERC1155(opts, nil, nil)
	if err != nil {
		log.Error("get erc1155 gateway refund iterator failed", "error", err)
		return nil, err
	}

	refundWrapIter := types.WrapIterator{
		Iter:      refundIter,
		EventType: types.L1RefundERC1155,
	}
	iterators = append(iterators, refundWrapIter)

	// batch refund
	batchRefundIter, err := l.l1Contracts.ERC1155Gateway.FilterBatchRefundERC1155(opts, nil, nil)
	if err != nil {
		log.Error("get erc1155 gateway refund iterator failed", "error", err)
		return nil, err
	}

	batchRefundWrapIter := types.WrapIterator{
		Iter:      batchRefundIter,
		EventType: types.L1BatchRefundERC1155,
	}
	iterators = append(iterators, batchRefundWrapIter)

	return iterators, nil
}

func (l *Contracts) l2Erc1155Filter(_ context.Context, opts *bind.FilterOpts) ([]types.WrapIterator, error) {
	var iterators []types.WrapIterator

	// withdraw
	withdrawIter, err := l.l2Contracts.ERC1155Gateway.FilterWithdrawERC1155(opts, nil, nil, nil)
	if err != nil {
		log.Error("get erc1155 gateway withdraw iterator failed", "error", err)
		return nil, err
	}

	withdrawWrapIter := types.WrapIterator{
		Iter:      withdrawIter,
		EventType: types.L2WithdrawERC1155,
	}
	iterators = append(iterators, withdrawWrapIter)

	// batch withdraw
	batchWithdrawIter, err := l.l2Contracts.ERC1155Gateway.FilterBatchWithdrawERC1155(opts, nil, nil, nil)
	if err != nil {
		log.Error("get erc1155 gateway batch withdraw iterator failed", "error", err)
		return nil, err
	}

	batchWithdrawWrapIter := types.WrapIterator{
		Iter:      batchWithdrawIter,
		EventType: types.L2BatchWithdrawERC1155,
	}
	iterators = append(iterators, batchWithdrawWrapIter)

	// finalizeDeposit
	finalizeDepositIter, err := l.l2Contracts.ERC1155Gateway.FilterFinalizeDepositERC1155(opts, nil, nil, nil)
	if err != nil {
		log.Error("get erc1155 gateway finalize deposit iterator failed", "error", err)
		return nil, err
	}

	finalizeWithdrawWrapIter := types.WrapIterator{
		Iter:      finalizeDepositIter,
		EventType: types.L2FinalizeDepositERC1155,
	}
	iterators = append(iterators, finalizeWithdrawWrapIter)

	// batch finalize Deposit
	batchFinalizeDepositIter, err := l.l2Contracts.ERC1155Gateway.FilterFinalizeBatchDepositERC1155(opts, nil, nil, nil)
	if err != nil {
		log.Error("get erc1155 gateway finalize batch deposit iterator failed", "error", err)
		return nil, err
	}

	batchFinalizeWithdrawWrapIter := types.WrapIterator{
		Iter:      batchFinalizeDepositIter,
		EventType: types.L2FinalizeBatchDepositERC1155,
	}
	iterators = append(iterators, batchFinalizeWithdrawWrapIter)

	return iterators, nil
}

func (l *Contracts) getL1Erc1155GatewayTransfer(ctx context.Context, startBlockNumber, endBlockNumber uint64) ([]events.EventUnmarshaler, error) {
	var eventUnmarshalerList []events.EventUnmarshaler
	singleEventUnmarshalerList, err := l.getL1Erc1155GatewayTransferSingle(ctx, startBlockNumber, endBlockNumber)
	if err != nil {
		return nil, err
	}
	batchEventUnmarshalerList, err := l.getL1Erc1155GatewayTransferBatch(ctx, startBlockNumber, endBlockNumber)
	if err != nil {
		return nil, err
	}

	eventUnmarshalerList = append(eventUnmarshalerList, singleEventUnmarshalerList...)
	eventUnmarshalerList = append(eventUnmarshalerList, batchEventUnmarshalerList...)
	return eventUnmarshalerList, nil
}

func (l *Contracts) getL1Erc1155GatewayTransferSingle(ctx context.Context, startBlockNumber, endBlockNumber uint64) ([]events.EventUnmarshaler, error) {
	erc1155ABI, err := iscrollerc1155.Iscrollerc1155MetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	transferEventQuery := ethereum.FilterQuery{
		FromBlock: new(big.Int).SetUint64(startBlockNumber),
		ToBlock:   new(big.Int).SetUint64(endBlockNumber),
		Topics:    [][]common.Hash{{erc1155ABI.Events["TransferSingle"].ID}},
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
			Value   *big.Int
		}{}
		if unpackErr := erc1155ABI.UnpackIntoInterface(&event, "TransferSingle", vLog.Data); unpackErr != nil {
			return nil, err
		}

		if l.l1Contracts.ERC1155GatewayAddress == event.From {
			transferEvents = append(transferEvents, &events.ERC1155GatewayEventUnmarshaler{
				TokenIds:     []*big.Int{event.TokenID},
				Amounts:      []*big.Int{new(big.Int).Neg(event.Value)},
				TokenAddress: vLog.Address,
				TxHash:       vLog.TxHash,
				Number:       vLog.BlockNumber,
			})
		}

		if l.l1Contracts.ERC1155GatewayAddress == event.To {
			transferEvents = append(transferEvents, &events.ERC1155GatewayEventUnmarshaler{
				TokenIds:     []*big.Int{event.TokenID},
				Amounts:      []*big.Int{event.Value},
				TokenAddress: vLog.Address,
				TxHash:       vLog.TxHash,
				Number:       vLog.BlockNumber,
			})
		}
	}

	return transferEvents, nil
}

func (l *Contracts) getL1Erc1155GatewayTransferBatch(ctx context.Context, startBlockNumber, endBlockNumber uint64) ([]events.EventUnmarshaler, error) {
	erc1155ABI, err := iscrollerc1155.Iscrollerc1155MetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	transferEventQuery := ethereum.FilterQuery{
		FromBlock: new(big.Int).SetUint64(startBlockNumber),
		ToBlock:   new(big.Int).SetUint64(endBlockNumber),
		Topics:    [][]common.Hash{{erc1155ABI.Events["TransferBatch"].ID}},
	}

	logs, err := l.l1Contracts.client.FilterLogs(ctx, transferEventQuery)
	if err != nil {
		return nil, err
	}

	var transferEvents []events.EventUnmarshaler
	for _, vLog := range logs {
		event := struct {
			From   common.Address
			To     common.Address
			Ids    []*big.Int
			Values []*big.Int
		}{}
		if unpackErr := erc1155ABI.UnpackIntoInterface(&event, "TransferBatch", vLog.Data); unpackErr != nil {
			return nil, err
		}

		if l.l1Contracts.ERC1155GatewayAddress == event.From {
			var tmpValues []*big.Int
			for _, v := range event.Values {
				tmpValues = append(tmpValues, new(big.Int).Neg(v))
			}
			transferEvents = append(transferEvents, &events.ERC1155GatewayEventUnmarshaler{
				TokenIds:     event.Ids,
				Amounts:      tmpValues,
				TokenAddress: vLog.Address,
				TxHash:       vLog.TxHash,
				Number:       vLog.BlockNumber,
			})
		}

		if l.l1Contracts.ERC1155GatewayAddress == event.To {
			transferEvents = append(transferEvents, &events.ERC1155GatewayEventUnmarshaler{
				TokenIds:     event.Ids,
				Amounts:      event.Values,
				TokenAddress: vLog.Address,
				TxHash:       vLog.TxHash,
				Number:       vLog.BlockNumber,
			})
		}
	}

	return transferEvents, nil
}

func (l *Contracts) getL2Erc1155GatewayTransfer(ctx context.Context, startBlockNumber, endBlockNumber uint64) ([]events.EventUnmarshaler, error) {
	var eventUnmarshalerList []events.EventUnmarshaler
	singleEventUnmarshalerList, err := l.getL2Erc1155GatewayTransferSingle(ctx, startBlockNumber, endBlockNumber)
	if err != nil {
		return nil, err
	}
	batchEventUnmarshalerList, err := l.getL2Erc1155GatewayTransferBatch(ctx, startBlockNumber, endBlockNumber)
	if err != nil {
		return nil, err
	}

	eventUnmarshalerList = append(eventUnmarshalerList, singleEventUnmarshalerList...)
	eventUnmarshalerList = append(eventUnmarshalerList, batchEventUnmarshalerList...)
	return eventUnmarshalerList, nil
}

func (l *Contracts) getL2Erc1155GatewayTransferSingle(ctx context.Context, startBlockNumber, endBlockNumber uint64) ([]events.EventUnmarshaler, error) {
	erc1155ABI, err := iscrollerc1155.Iscrollerc1155MetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	transferEventQuery := ethereum.FilterQuery{
		FromBlock: new(big.Int).SetUint64(startBlockNumber),
		ToBlock:   new(big.Int).SetUint64(endBlockNumber),
		Topics:    [][]common.Hash{{erc1155ABI.Events["TransferSingle"].ID}},
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
			Value   *big.Int
		}{}
		if err := erc1155ABI.UnpackIntoInterface(&event, "TransferSingle", vLog.Data); err != nil {
			return nil, err
		}

		if event.From == l.l2Contracts.ERC1155GatewayAddress {
			transferEvents = append(transferEvents, &events.ERC1155GatewayEventUnmarshaler{
				TokenIds:     []*big.Int{event.TokenID},
				Amounts:      []*big.Int{new(big.Int).Neg(event.Value)},
				TokenAddress: vLog.Address,
				TxHash:       vLog.TxHash,
				Number:       vLog.BlockNumber,
				Layer:        types.Layer1,
			})
		}

		if event.To == l.l2Contracts.ERC1155GatewayAddress {
			transferEvents = append(transferEvents, &events.ERC1155GatewayEventUnmarshaler{
				TokenIds:     []*big.Int{event.Value},
				Amounts:      []*big.Int{event.Value},
				TokenAddress: vLog.Address,
				TxHash:       vLog.TxHash,
				Number:       vLog.BlockNumber,
				Layer:        types.Layer1,
			})
		}
	}

	return transferEvents, nil
}

func (l *Contracts) getL2Erc1155GatewayTransferBatch(ctx context.Context, startBlockNumber, endBlockNumber uint64) ([]events.EventUnmarshaler, error) {
	erc1155ABI, err := iscrollerc1155.Iscrollerc1155MetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	transferEventQuery := ethereum.FilterQuery{
		FromBlock: new(big.Int).SetUint64(startBlockNumber),
		ToBlock:   new(big.Int).SetUint64(endBlockNumber),
		Topics:    [][]common.Hash{{erc1155ABI.Events["TransferBatch"].ID}},
	}

	logs, err := l.l2Contracts.client.FilterLogs(ctx, transferEventQuery)
	if err != nil {
		return nil, err
	}

	var transferEvents []events.EventUnmarshaler
	for _, vLog := range logs {
		event := struct {
			From   common.Address
			To     common.Address
			Ids    []*big.Int
			Values []*big.Int
		}{}
		if err := erc1155ABI.UnpackIntoInterface(&event, "TransferBatch", vLog.Data); err != nil {
			return nil, err
		}

		if event.From == l.l2Contracts.ERC1155GatewayAddress {
			var tmpValues []*big.Int
			for _, v := range event.Values {
				tmpValues = append(tmpValues, new(big.Int).Neg(v))
			}
			transferEvents = append(transferEvents, &events.ERC1155GatewayEventUnmarshaler{
				TokenIds:     event.Ids,
				Amounts:      tmpValues,
				TokenAddress: vLog.Address,
				TxHash:       vLog.TxHash,
				Number:       vLog.BlockNumber,
				Layer:        types.Layer2,
			})
		}

		if event.To == l.l2Contracts.ERC1155GatewayAddress {
			transferEvents = append(transferEvents, &events.ERC1155GatewayEventUnmarshaler{
				TokenIds:     event.Ids,
				Amounts:      event.Values,
				TokenAddress: vLog.Address,
				TxHash:       vLog.TxHash,
				Number:       vLog.BlockNumber,
				Layer:        types.Layer2,
			})
		}
	}

	return transferEvents, nil
}
