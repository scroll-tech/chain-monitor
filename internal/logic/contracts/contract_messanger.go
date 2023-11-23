package contracts

import (
	"context"

	"github.com/scroll-tech/go-ethereum/accounts/abi/bind"
	"github.com/scroll-tech/go-ethereum/log"

	"github.com/scroll-tech/chain-monitor/internal/types"
)

func (l *Contracts) l1MessengerFilter(_ context.Context, opts *bind.FilterOpts) ([]types.WrapIterator, error) {
	var iterators []types.WrapIterator
	sentMessageIter, err := l.l1Contracts.messenger.FilterSentMessage(opts, nil, nil)
	if err != nil {
		log.Error("get messenger sentMessage iterator failed", "error", err)
		return nil, err
	}

	sentMessageWrapIter := types.WrapIterator{
		Iter:      sentMessageIter,
		EventType: types.L1SentMessage,
	}
	iterators = append(iterators, sentMessageWrapIter)

	relayedMessageIter, err := l.l1Contracts.messenger.FilterRelayedMessage(opts, nil)
	if err != nil {
		log.Error("get messenger relayedMessage iterator failed", "error", err)
		return nil, err
	}

	relayedMessageWrapIter := types.WrapIterator{
		Iter:      relayedMessageIter,
		EventType: types.L1RelayedMessage,
	}
	iterators = append(iterators, relayedMessageWrapIter)
	return iterators, nil
}

func (l *Contracts) l2MessengerFilter(_ context.Context, opts *bind.FilterOpts) ([]types.WrapIterator, error) {
	var iterators []types.WrapIterator
	sentMessageIter, err := l.l2Contracts.messenger.FilterSentMessage(opts, nil, nil)
	if err != nil {
		log.Error("get messenger sentMessage iterator failed", "error", err)
		return nil, err
	}

	sentMessageWrapIter := types.WrapIterator{
		Iter:      sentMessageIter,
		EventType: types.L2SentMessage,
	}
	iterators = append(iterators, sentMessageWrapIter)

	relayedMessageIter, err := l.l2Contracts.messenger.FilterRelayedMessage(opts, nil)
	if err != nil {
		log.Error("get messenger relayedMessage iterator failed", "error", err)
		return nil, err
	}

	relayedMessageWrapIter := types.WrapIterator{
		Iter:      relayedMessageIter,
		EventType: types.L2RelayedMessage,
	}
	iterators = append(iterators, relayedMessageWrapIter)
	return iterators, nil
}
