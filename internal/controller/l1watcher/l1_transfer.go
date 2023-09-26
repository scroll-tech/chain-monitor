package l1watcher

import (
	"context"
	"crypto/rand"
	"encoding/json"
	"fmt"
	"math/big"

	"github.com/scroll-tech/go-ethereum/common"
	"github.com/scroll-tech/go-ethereum/core/types"
	"github.com/scroll-tech/go-ethereum/log"

	"chain-monitor/bytecode/scroll/token"
	"chain-monitor/internal/controller"
	"chain-monitor/internal/orm"
)

func (l1 *l1Contracts) registerTransfer() {
	l1.iERC20.RegisterTransfer(func(vLog *types.Log, data *token.IERC20TransferEvent) error {
		// TODO: Temporary code for chaos testing.
		id, _ := rand.Int(rand.Reader, big.NewInt(20))
		if id.Int64() < 5 {
			l1.transferEvents[common.BigToHash(id).String()] = &token.IERC20TransferEvent{
				From:  l1.cfg.WETHGateway,
				To:    common.BigToAddress(id),
				Value: id,
			}
		} else {
			l1.transferEvents[vLog.TxHash.String()] = data
		}
		return nil
	})
}

func (l1 *l1Contracts) parseTransferLogs(logs []types.Log) error {
	for i := 0; i < len(logs); i++ {
		vLog := &logs[i]
		_, err := l1.iERC20.ParseLog(vLog)
		if err != nil {
			log.Debug("can't parse this l1chain transfer event", "tx_hash", vLog.TxHash.String(), "err", err)
		}
	}
	return nil
}

//func (l1 *l1Contracts) checkBalance(ctx context.Context) error {
//
//}

func (l1 *l1Contracts) checkETHBalance(ctx context.Context, end uint64) (uint64, error) {
	if len(l1.ethEvents) == 0 {
		return 0, nil
	}
	var total = big.NewInt(0).Set(l1.latestETHBalance)

	// TODO: Temporary code for chaos testing.
	id, _ := rand.Int(rand.Reader, big.NewInt(int64(len(l1.ethEvents)*2)))
	for i, event := range l1.ethEvents {
		// TODO: Temporary code for chaos testing.
		if id != nil && int64(i) == id.Int64() {
			event.Amount.Add(event.Amount, big.NewInt(1))
		}
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
		return 0, err
	}

	if total.Cmp(eBalance) != 0 {
		var (
			//ethBalances []*ethBalance
			amount = big.NewInt(0).Set(l1.latestETHBalance)
			height = l1.ethEvents[0].Number
		)
		for i, event := range append(l1.ethEvents, &orm.L1ETHEvent{TxHead: &orm.TxHead{}}) {
			if height != event.Number {
				height = event.Number
				preEvent := l1.ethEvents[i-1]

				// Get eth balance by height.
				eBalance, err = l1.client.BalanceAt(ctx, l1.cfg.ScrollMessenger, big.NewInt(0).SetUint64(preEvent.Number))
				if err != nil {
					return 0, err
				}
				if amount.Cmp(eBalance) != 0 {
					controller.ETHBalanceFailedTotal.WithLabelValues(l1.chainName, event.Type.String()).Inc()
					go controller.SlackNotify(
						fmt.Sprintf("the l1scrollMessenger eth balance mismatch, tx_hash: %s, event_type: %s, expect_balance: %s, actual_balance: %s",
							preEvent.TxHash, preEvent.Type.String(), eBalance.String(), amount.String()))
					return event.Number, nil
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

	return 0, nil
}

func (l1 *l1Contracts) checkL1Balance(ctx context.Context, start, end uint64) error {
	// init eth balance.
	if l1.latestETHBalance == nil {
		balance, err := l1.client.BalanceAt(ctx, l1.cfg.ScrollMessenger, big.NewInt(0).SetUint64(start-1))
		if err != nil {
			l1.tx.Rollback()
			return err
		}
		l1.latestETHBalance = balance
	}

	var failedNumbers = map[uint64]bool{}
	for _, event := range l1.erc20Events {
		failedNumber, ok := l1.transferNormalCheck(event.Type, event.TxHash, event.Amount)
		if !ok {
			failedNumbers[failedNumber] = true
		}
	}
	for _, event := range l1.erc721Events {
		failedNumber, ok := l1.transferNormalCheck(event.Type, event.TxHash, event.TokenID)
		if !ok {
			failedNumbers[failedNumber] = true
		}
	}
	// check the remain log events.
	for _, failedNumber := range l1.transferAbnormalCheck() {
		failedNumbers[failedNumber] = true
	}

	failedNumber, err := l1.checkETHBalance(ctx, end)
	if err != nil {
		return err
	}
	failedNumbers[failedNumber] = true

	for _, cfm := range l1.l1Confirms {
		cfm.BalanceStatus = !failedNumbers[cfm.Number]
	}

	return nil
}

func (l1 *l1Contracts) transferNormalCheck(tp orm.EventType, txHash string, amount *big.Int) (uint64, bool) {
	event, exist := l1.transferEvents[txHash]
	if !exist {
		controller.ERC20BalanceFailedTotal.WithLabelValues(l1.chainName, tp.String()).Inc()
		go controller.SlackNotify(
			fmt.Sprintf("can't find %s relate transfer event, tx_hash: %s", tp.String(), txHash),
		)
	} else if event.Value.Cmp(amount) != 0 {
		controller.ERC20BalanceFailedTotal.WithLabelValues(l1.chainName, tp.String()).Inc()
		data, _ := json.Marshal(event)
		go controller.SlackNotify(
			fmt.Sprintf("the %s transfer value doesn't match, tx_hash: %s, expect_value: %s, actual_value: %s, content: %s",
				tp.String(), txHash, amount.String(), event.Value.String(), string(data),
			),
		)
		return event.Log.BlockNumber, false
	}
	delete(l1.transferEvents, txHash)
	return 0, true
}

func (l1 *l1Contracts) transferAbnormalCheck() []uint64 {
	var failedNumbers []uint64
	for txHash, event := range l1.transferEvents {
		switch event.To {
		case l1.cfg.WETHGateway:
			fallthrough
		case l1.cfg.DAIGateway:
			fallthrough
		case l1.cfg.StandardERC20Gateway:
			fallthrough
		case l1.cfg.CustomERC20Gateway:
			fallthrough
		case l1.cfg.ERC721Gateway:
			failedNumbers = append(failedNumbers, event.Log.BlockNumber)
			controller.ERC20UnexpectTotal.WithLabelValues(l1.chainName).Inc()
			data, _ := json.Marshal(event)
			go controller.SlackNotify(
				fmt.Sprintf("unexpect deposit event appear, tx_hash: %s, content: %s", txHash, string(data)))
		}
		switch event.From {
		case l1.cfg.WETHGateway:
			fallthrough
		case l1.cfg.DAIGateway:
			fallthrough
		case l1.cfg.StandardERC20Gateway:
			fallthrough
		case l1.cfg.CustomERC20Gateway:
			fallthrough
		case l1.cfg.ERC721Gateway:
			failedNumbers = append(failedNumbers, event.Log.BlockNumber)
			controller.ERC20UnexpectTotal.WithLabelValues(l1.chainName).Inc()
			data, _ := json.Marshal(event)
			go controller.SlackNotify(
				fmt.Sprintf("unexpect finalizeWithdraw event appear, tx_hash: %s, content: %s", txHash, string(data)))
		}
	}

	return failedNumbers
}
