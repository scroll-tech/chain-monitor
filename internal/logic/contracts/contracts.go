package contracts

import (
	"context"
	"fmt"
	"math/big"

	"github.com/scroll-tech/go-ethereum"
	"github.com/scroll-tech/go-ethereum/accounts/abi/bind"
	"github.com/scroll-tech/go-ethereum/common"
	"github.com/scroll-tech/go-ethereum/log"
	"github.com/scroll-tech/go-ethereum/rpc"

	"github.com/scroll-tech/chain-monitor/internal/config"
	"github.com/scroll-tech/chain-monitor/internal/logic/contracts/abi/il2scrollmessenger"
	"github.com/scroll-tech/chain-monitor/internal/logic/contracts/abi/iscrollerc20"
	"github.com/scroll-tech/chain-monitor/internal/logic/events"
	"github.com/scroll-tech/chain-monitor/internal/types"
	"github.com/scroll-tech/chain-monitor/internal/utils"
)

// Contracts fetch/watch the logs from l1/l2
type Contracts struct {
	l1Contracts *l1Contracts
	l2Contracts *l2Contracts
}

// NewContracts create contracts filter logs fetcher
func NewContracts(l1Client, l2Client *rpc.Client) *Contracts {
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

	return nil, fmt.Errorf("invalid type, layerType: %v, txEventCategory: %v", layerType, txEventCategory)
}

func (l *Contracts) GetGatewayTransfer(ctx context.Context, startBlockNumber, endBlockNumber uint64, layerType types.LayerType, txEventCategory types.TxEventCategory) ([]events.EventUnmarshaler, error) {
	if layerType == types.Layer1 {
		switch txEventCategory {
		case types.ERC20EventCategory:
			return l.getl1Erc20GatewayTransfer(ctx, startBlockNumber, endBlockNumber)
		case types.ERC721EventCategory:
		case types.ERC1155EventCategory:
		}
	}

	if layerType == types.Layer2 {
		switch txEventCategory {
		case types.ERC20EventCategory:
			return l.getl2Erc20GatewayTransfer(ctx, startBlockNumber, endBlockNumber)
		case types.ERC721EventCategory:
		case types.ERC1155EventCategory:
		}
	}

	return nil, fmt.Errorf("invalid type, layerType: %v, txEventCategory: %v", layerType, txEventCategory)
}

func (l *Contracts) GetL2SentMessageEventsAndWithdrawRoots(ctx context.Context, startBlockNumber, endBlockNumber uint64) (map[uint64][]events.SentMessageEvent, map[uint64]common.Hash, error) {
	l2MessengerABI, err := il2scrollmessenger.Il2scrollmessengerMetaData.GetAbi()
	if err != nil {
		return nil, nil, err
	}
	sentMessageEventQuery := ethereum.FilterQuery{
		FromBlock: new(big.Int).SetUint64(startBlockNumber),
		ToBlock:   new(big.Int).SetUint64(endBlockNumber),
		Topics:    [][]common.Hash{{l2MessengerABI.Events["SentMessage"].ID}},
	}

	logs, err := l.l1Contracts.client.FilterLogs(ctx, sentMessageEventQuery)
	if err != nil {
		return nil, nil, err
	}

	sentMessageEventsMap := make(map[uint64][]events.SentMessageEvent)
	for _, vLog := range logs {
		var event events.SentMessageEvent
		err := l2MessengerABI.UnpackIntoInterface(&event, "SentMessage", vLog.Data)
		if err != nil {
			return nil, nil, err
		}
		blockNumber := vLog.BlockNumber
		event.MessageHash = utils.ComputeMessageHash(event.Sender, event.Target, event.Value, event.MessageNonce, event.Message)
		sentMessageEventsMap[blockNumber] = append(sentMessageEventsMap[blockNumber], event)
	}

	withdrawRootsMap, err := utils.GetBatchWithdrawRootsInRange(ctx, l.l2Contracts.rpcClient, l.l2Contracts.l2Config.L2Contracts.MessageQueue, startBlockNumber, endBlockNumber)
	if err != nil {
		return nil, nil, err
	}

	return sentMessageEventsMap, withdrawRootsMap, nil
}

func (l *Contracts) getl1Erc20GatewayTransfer(ctx context.Context, startBlockNumber, endBlockNumber uint64) ([]events.EventUnmarshaler, error) {
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

func (l *Contracts) getl2Erc20GatewayTransfer(ctx context.Context, startBlockNumber, endBlockNumber uint64) ([]events.EventUnmarshaler, error) {
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

		sentMessageIter, err := l.l1Contracts.Messenger.FilterSentMessage(opts, []common.Address{erc20Token.Address}, nil)
		if err != nil {
			log.Error("get messenger sentMessage iterator failed", "token type", erc20Token.TokenType, "address", erc20Token.Address, "error", err)
			return nil, err
		}

		sentMessageWrapIter := types.WrapIterator{
			Iter:      sentMessageIter,
			EventType: types.L1SentMessage,
		}
		iterators = append(iterators, sentMessageWrapIter)

		// Note: can only select all relayed message within the block range.
		relayedMessageIter, err := l.l1Contracts.Messenger.FilterRelayedMessage(opts, nil)
		if err != nil {
			log.Error("get messenger relayedMessage iterator failed", "token type", erc20Token.TokenType, "address", erc20Token.Address, "error", err)
			return nil, err
		}

		relayedMessageWrapIter := types.WrapIterator{
			Iter:      relayedMessageIter,
			EventType: types.L1RelayedMessage,
		}
		iterators = append(iterators, relayedMessageWrapIter)
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

		sentMessageIter, err := l.l2Contracts.Messenger.FilterSentMessage(opts, []common.Address{erc20Token.Address}, nil)
		if err != nil {
			log.Error("get messenger sentMessage iterator failed", "token type", erc20Token.TokenType, "address", erc20Token.Address, "error", err)
			return nil, err
		}

		sentMessageWrapIter := types.WrapIterator{
			Iter:      sentMessageIter,
			EventType: types.L2SentMessage,
		}
		iterators = append(iterators, sentMessageWrapIter)

		// Note: can only select all relayed message within the block range.
		relayedMessageIter, err := l.l2Contracts.Messenger.FilterRelayedMessage(opts, nil)
		if err != nil {
			log.Error("get messenger relayedMessage iterator failed", "token type", erc20Token.TokenType, "address", erc20Token.Address, "error", err)
			return nil, err
		}

		relayedMessageWrapIter := types.WrapIterator{
			Iter:      relayedMessageIter,
			EventType: types.L2RelayedMessage,
		}
		iterators = append(iterators, relayedMessageWrapIter)
	}
	return iterators, nil
}
