package l1watcher

import (
	"context"
	"encoding/json"
	"fmt"
	"math/big"

	"github.com/scroll-tech/go-ethereum/core/types"
	"github.com/scroll-tech/go-ethereum/log"

	"chain-monitor/bytecode/scroll/token"
	"chain-monitor/internal/controller"
	"chain-monitor/internal/orm"
)

func (l1 *l1Contracts) registerTransfer() {
	l1.iERC20.RegisterTransfer(func(vLog *types.Log, data *token.IScrollERC20TransferEvent) error {
		l1.transferEvents[vLog.TxHash.String()] = data
		return nil
	})
}

func (l1 *l1Contracts) parseTransferLogs(logs []types.Log) error {
	for i := 0; i < len(logs); i++ {
		vLog := &logs[i]
		_, err := l1.iERC20.ParseLog(vLog)
		if err != nil {
			log.Debug("can't parse this log to transfer", "tx_hash", vLog.TxHash.String(), "err", err)
		}
	}
	return nil
}

func (l1 *l1Contracts) checkETHBalance(ctx context.Context, start, end uint64) error {
	if len(l1.ethEvents) == 0 {
		return nil
	}
	var total = big.NewInt(0).Set(l1.latestETHBalance)
	for _, event := range l1.ethEvents {
		if event.Type == orm.L1FinalizeWithdrawETH {
			total.Sub(total, event.Amount)
		}
		if event.Type == orm.L1DepositETH {
			total.Add(total, event.Amount)
		}
	}

	// Get latest eth balance.
	eBalance, err := l1.client.BalanceAt(ctx, l1.cfg.ScrollMessenger, big.NewInt(0).SetUint64(end))
	if err != nil {
		return err
	}

	if total.Cmp(eBalance) != 0 {
		var (
			//ethBalances []*ethBalance
			amount = big.NewInt(0).Set(l1.latestETHBalance)
			height = l1.ethEvents[0].Number
		)
		for i, event := range l1.ethEvents {
			if height != event.Number {
				height = event.Number
				preEvent := l1.ethEvents[i-1]

				// Get eth balance by height.
				eBalance, err = l1.client.BalanceAt(ctx, l1.cfg.ScrollMessenger, big.NewInt(0).SetUint64(preEvent.Number))
				if err != nil {
					return err
				}
				if amount.Cmp(eBalance) != 0 {
					data, _ := json.Marshal(preEvent)
					go controller.SlackNotify(fmt.Sprintf("the l1scrollMessenger eth balance mismatch, event_type: %s, expect_balance: %s, actual_balance: %s, content: %s", preEvent.Type.String(), eBalance.String(), amount.String(), string(data)))
				}
			}
			if event.Type == orm.L1FinalizeWithdrawETH {
				amount.Sub(amount, event.Amount)
			}
			if event.Type == orm.L1DepositETH {
				amount.Add(amount, event.Amount)
			}
		}
	}
	l1.latestETHBalance = eBalance

	return nil
}
