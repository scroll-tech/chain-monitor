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
	l1.iERC20.RegisterTransfer(func(vLog *types.Log, data *token.IERC20TransferEvent) error {
		l1.transferEvents[vLog.TxHash.String()] = data
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

func (l1 *l1Contracts) checkETHBalance(ctx context.Context, start, end uint64) (uint64, error) {
	if len(l1.ethEvents) == 0 {
		return 0, nil
	}

	// Get balance at start number.
	sBalance, err := l1.client.BalanceAt(ctx, l1.cfg.ScrollMessenger, big.NewInt(0).SetUint64(start-1))
	if err != nil {
		return 0, err
	}

	// Get latest eth balance.
	eBalance, err := l1.client.BalanceAt(ctx, l1.cfg.ScrollMessenger, big.NewInt(0).SetUint64(end))
	if err != nil {
		return 0, err
	}

	var total = big.NewInt(0).Set(sBalance)

	for _, event := range l1.ethEvents {
		if event.Type == orm.L1FinalizeWithdrawETH {
			total.Sub(total, event.Amount)
		}
		if event.Type == orm.L1DepositETH {
			total.Add(total, event.Amount)
		}
	}

	if total.Cmp(eBalance) != 0 {
		var (
			//ethBalances []*ethBalance
			amount = big.NewInt(0).Set(sBalance)
			height = l1.ethEvents[0].Number
		)
		for i, event := range append(l1.ethEvents, &orm.L1ETHEvent{TxHead: &orm.TxHead{}}) {
			if height != event.Number {
				height = event.Number
				preEvent := l1.ethEvents[i-1]

				// Get eth balance by height.
				balance, err := l1.client.BalanceAt(ctx, l1.cfg.ScrollMessenger, big.NewInt(0).SetUint64(preEvent.Number))
				if err != nil {
					return 0, err
				}
				if amount.Cmp(balance) != 0 {
					controller.ETHBalanceFailedTotal.WithLabelValues(l1.chainName, event.Type.String()).Inc()
					go controller.SlackNotify(
						fmt.Sprintf("the l1scrollMessenger eth balance mismatch, tx_hash: %s, event_type: %s, expect_balance: %s, actual_balance: %s",
							preEvent.TxHash, preEvent.Type.String(), balance.String(), amount.String()))
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

	return 0, nil
}

func (l1 *l1Contracts) checkL1Balance(ctx context.Context, start, end uint64) error {
	var failedNumbers = map[uint64]bool{}
	for _, event := range l1.erc20Events {
		if !l1.transferNormalCheck(event.Type, event.TxHash, event.Amount) {
			failedNumbers[event.Number] = true
		}
	}
	for _, event := range l1.erc721Events {
		if !l1.transferNormalCheck(event.Type, event.TxHash, event.TokenID) {
			failedNumbers[event.Number] = true
		}
	}
	// check the remain log events.
	for _, failedNumber := range l1.transferAbnormalCheck() {
		failedNumbers[failedNumber] = true
	}

	failedNumber, err := l1.checkETHBalance(ctx, start, end)
	if err != nil {
		return err
	}
	failedNumbers[failedNumber] = true

	for _, cfm := range l1.l1Confirms {
		cfm.BalanceStatus = !failedNumbers[cfm.Number]
	}

	return nil
}

func (l1 *l1Contracts) transferNormalCheck(tp orm.EventType, txHash string, amount *big.Int) bool {
	event, exist := l1.transferEvents[txHash]
	if !exist {
		controller.ERC20BalanceFailedTotal.WithLabelValues(l1.chainName, tp.String()).Inc()
		go controller.SlackNotify(fmt.Sprintf("can't find %s relate transfer event, tx_hash: %s", tp.String(), txHash))
		return false
	} else if event.Value.Cmp(amount) != 0 {
		controller.ERC20BalanceFailedTotal.WithLabelValues(l1.chainName, tp.String()).Inc()
		data, _ := json.Marshal(event)
		go controller.SlackNotify(fmt.Sprintf("the %s transfer value doesn't match, tx_hash: %s, expect_value: %s, actual_value: %s, content: %s",
			tp.String(), txHash, amount.String(), event.Value.String(), string(data)))
		return false
	}
	delete(l1.transferEvents, txHash)
	return true
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
