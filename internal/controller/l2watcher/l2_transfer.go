package l2watcher

import (
	"context"
	"encoding/json"
	"fmt"
	"math/big"
	"sort"

	"github.com/scroll-tech/go-ethereum/core/types"
	"github.com/scroll-tech/go-ethereum/log"

	"chain-monitor/bytecode/scroll/token"
	"chain-monitor/internal/controller"
	"chain-monitor/internal/orm"
)

func (l2 *l2Contracts) registerTransfer() {
	l2.iERC20.RegisterTransfer(func(vLog *types.Log, data *token.IScrollERC20TransferEvent) error {
		l2.transferEvents[vLog.TxHash.String()] = data
		return nil
	})
}

func (l2 *l2Contracts) parseTransferLogs(logs []types.Log) error {
	for i := 0; i < len(logs); i++ {
		vLog := &logs[i]
		_, err := l2.iERC20.ParseLog(vLog)
		if err != nil {
			log.Debug("can't parse transfer event", "tx_hash", vLog.TxHash.String(), "err", err)
		}
	}
	return nil
}

// This struct is used to check scrollMessenger eth balance.
type ethEvent struct {
	number uint64
	tp     orm.EventType
	amount *big.Int
}

func (l2 *l2Contracts) checkETHBalance(ctx context.Context, start, end uint64) error {
	if len(l2.ethEvents) == 0 {
		return nil
	}
	var (
		total  = big.NewInt(0).Set(l2.latestETHBalance)
		events []*ethEvent
	)
	for _, event := range l2.ethEvents {
		if event.Type == orm.L2FinalizeDepositETH {
			total.Sub(total, event.Amount)
		}
		if event.Type == orm.L2WithdrawETH {
			total.Add(total, event.Amount)
		}
		events = append(events, &ethEvent{
			number: event.Number,
			tp:     event.Type,
			amount: event.Amount,
		})
	}
	for _, event := range l2.erc20Events {
		if event.Type == orm.L2FinalizeDepositWETH {
			total.Sub(total, event.Amount)
		}
		if event.Type == orm.L2WithdrawWETH {
			total.Add(total, event.Amount)
		}
		events = append(events, &ethEvent{
			number: event.Number,
			tp:     event.Type,
			amount: event.Amount,
		})
	}

	// Get latest eth balance.
	eBalance, err := l2.client.BalanceAt(ctx, l2.cfg.ScrollMessenger, big.NewInt(0).SetUint64(end))
	if err != nil {
		return err
	}

	if total.Cmp(eBalance) != 0 {
		sort.Slice(events, func(i, j int) bool {
			return events[i].number < events[j].number
		})
		var (
			//ethBalances []*ethBalance
			amount = big.NewInt(0).Set(l2.latestETHBalance)
			height = events[0].number
		)
		for i, event := range events {
			if height != event.number {
				height = event.number
				preEvent := events[i-1]

				// Get eth balance by height.
				eBalance, err = l2.client.BalanceAt(ctx, l2.cfg.ScrollMessenger, big.NewInt(0).SetUint64(preEvent.number))
				if err != nil {
					return err
				}
				if amount.Cmp(eBalance) != 0 {
					data, _ := json.Marshal(preEvent)
					go controller.SlackNotify(fmt.Sprintf("the l2scrollMessenger eth balance mismatch, event_type: %s, expect_balance: %s, actual_balance: %s, content: %s", preEvent.tp.String(), eBalance.String(), amount.String(), string(data)))
				}
			}
			if event.tp == orm.L2FinalizeDepositETH {
				total.Sub(total, event.amount)
			}
			if event.tp == orm.L2WithdrawETH {
				total.Add(total, event.amount)
			}
			if event.tp == orm.L2FinalizeDepositWETH {
				total.Sub(total, event.amount)
			}
			if event.tp == orm.L2WithdrawWETH {
				total.Add(total, event.amount)
			}
		}
	}
	l2.latestETHBalance = eBalance

	return nil
}
