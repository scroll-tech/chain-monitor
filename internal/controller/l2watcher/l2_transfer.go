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
	l2.iERC20.RegisterTransfer(func(vLog *types.Log, data *token.IERC20TransferEvent) error {
		l2.transferEvents[vLog.TxHash.String()] = data
		return nil
	})
}

func (l2 *l2Contracts) parseTransferLogs(logs []types.Log) error {
	for i := 0; i < len(logs); i++ {
		vLog := &logs[i]
		_, err := l2.iERC20.ParseLog(vLog)
		if err != nil {
			log.Debug("can't parse this l2chain transfer event", "tx_hash", vLog.TxHash.String(), "err", err)
		}
	}
	return nil
}

// This struct is used to check scrollMessenger eth balance.
type ethEvent struct {
	Number uint64
	TxHash string
	Type   orm.EventType
	Amount *big.Int
}

func (l2 *l2Contracts) checkETHBalance(ctx context.Context, end uint64) (uint64, error) {
	if len(l2.ethEvents) == 0 {
		return 0, nil
	}
	var (
		total  = big.NewInt(0).Set(l2.latestETHBalance)
		events []*ethEvent
	)

	for i, event := range l2.ethEvents {
		if event.Type == orm.L2FinalizeDepositETH {
			total.Sub(total, event.Amount)
		}
		if event.Type == orm.L2WithdrawETH {
			total.Add(total, event.Amount)
		}
		events = append(events, &ethEvent{
			Number: event.Number,
			TxHash: event.TxHash,
			Type:   event.Type,
			Amount: event.Amount,
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
			Number: event.Number,
			TxHash: event.TxHash,
			Type:   event.Type,
			Amount: event.Amount,
		})
	}

	// Get latest eth balance.
	eBalance, err := l2.client.BalanceAt(ctx, l2.cfg.ScrollMessenger, big.NewInt(0).SetUint64(end))
	if err != nil {
		return 0, err
	}

	if total.Cmp(eBalance) != 0 {
		sort.Slice(events, func(i, j int) bool {
			return events[i].Number < events[j].Number
		})
		var (
			//ethBalances []*ethBalance
			amount = big.NewInt(0).Set(l2.latestETHBalance)
			height = events[0].Number
		)
		for i, event := range append(events, &ethEvent{Amount: big.NewInt(0)}) {
			if height != event.Number {
				height = event.Number
				preEvent := events[i-1]

				// Get eth balance by height.
				eBalance, err = l2.client.BalanceAt(ctx, l2.cfg.ScrollMessenger, big.NewInt(0).SetUint64(preEvent.Number))
				if err != nil {
					return 0, err
				}
				if amount.Cmp(eBalance) != 0 {
					controller.ETHBalanceFailedTotal.WithLabelValues(l2.chainName, event.Type.String()).Inc()
					go controller.SlackNotify(fmt.Sprintf("the l2scrollMessenger eth balance mismatch, tx_hash: %s, event_type: %s, expect_balance: %s, actual_balance: %s",
						preEvent.TxHash, preEvent.Type.String(), eBalance.String(), amount.String()))
					return event.Number, nil
				}
			}
			if event.Type == orm.L2FinalizeDepositETH {
				total.Sub(total, event.Amount)
			}
			if event.Type == orm.L2WithdrawETH {
				total.Add(total, event.Amount)
			}
			if event.Type == orm.L2FinalizeDepositWETH {
				total.Sub(total, event.Amount)
			}
			if event.Type == orm.L2WithdrawWETH {
				total.Add(total, event.Amount)
			}
		}
	}
	l2.latestETHBalance = eBalance

	return 0, nil
}

func (l2 *l2Contracts) checkL2Balance(ctx context.Context, start, end uint64) error {
	// init eth balance.
	if l2.latestETHBalance == nil {
		balance, err := l2.client.BalanceAt(ctx, l2.cfg.ScrollMessenger, big.NewInt(0).SetUint64(start-1))
		if err != nil {
			return err
		}
		l2.latestETHBalance = balance
	}

	var failedNumbers = map[uint64]bool{}
	for _, event := range l2.erc20Events {
		if !l2.transferNormalCheck(event.Type, event.TxHash, event.Amount) {
			failedNumbers[event.Number] = true
		}
	}
	for _, event := range l2.erc721Events {
		if !l2.transferNormalCheck(event.Type, event.TxHash, event.TokenID) {
			failedNumbers[event.Number] = true
		}
	}
	// check the remain log events.
	for _, failedNumber := range l2.transferAbnormalCheck() {
		failedNumbers[failedNumber] = true
	}

	// Check scroll messenger eth balance.
	failedNumber, err := l2.checkETHBalance(ctx, end)
	if err != nil {
		return err
	}
	failedNumbers[failedNumber] = true

	// Update balance_status
	for number, cfm := range l2.l2Confirms {
		cfm.BalanceStatus = !failedNumbers[number]
	}

	return nil
}

func (l2 *l2Contracts) transferNormalCheck(tp orm.EventType, txHash string, amount *big.Int) bool {
	event, exist := l2.transferEvents[txHash]
	if !exist {
		controller.ERC20BalanceFailedTotal.WithLabelValues(l2.chainName, tp.String()).Inc()
		go controller.SlackNotify(fmt.Sprintf("can't find %s relate transfer event, tx_hash: %s", tp.String(), txHash))
		return false
	} else if event.Value.Cmp(amount) != 0 {
		controller.ERC20BalanceFailedTotal.WithLabelValues(l2.chainName, tp.String()).Inc()
		go controller.SlackNotify(fmt.Sprintf("the %s transfer value doesn't match, tx_hash: %s, expect_value: %s, actual_value: %s", tp.String(), txHash, amount.String(), event.Value.String()))
		return false
	}
	delete(l2.transferEvents, txHash)

	return true
}

func (l2 *l2Contracts) transferAbnormalCheck() []uint64 {
	var failedNumbers []uint64
	// unexpect mint or burn operation.
	for txHash, event := range l2.transferEvents {
		switch event.To {
		case l2.cfg.WETHGateway:
			fallthrough
		case l2.cfg.DAIGateway:
			fallthrough
		case l2.cfg.StandardERC20Gateway:
			fallthrough
		case l2.cfg.CustomERC20Gateway:
			fallthrough
		case l2.cfg.ERC721Gateway:
			failedNumbers = append(failedNumbers, event.Log.BlockNumber)
			controller.ERC20UnexpectTotal.WithLabelValues(l2.chainName).Inc()
			data, _ := json.Marshal(event)
			go controller.SlackNotify(
				fmt.Sprintf("l2chain unexpect transfer used to(gateway) address, tx_hash: %x, content: %s",
					txHash, string(data)),
			)
		}
		switch event.From {
		case l2.cfg.WETHGateway:
			fallthrough
		case l2.cfg.DAIGateway:
			fallthrough
		case l2.cfg.StandardERC20Gateway:
			fallthrough
		case l2.cfg.CustomERC20Gateway:
			fallthrough
		case l2.cfg.ERC721Gateway:
			failedNumbers = append(failedNumbers, event.Log.BlockNumber)
			controller.ERC20UnexpectTotal.WithLabelValues(l2.chainName).Inc()
			data, _ := json.Marshal(event)
			go controller.SlackNotify(
				fmt.Sprintf("l2chain unexpect transfer used from(gateway) address, tx_hash: %x, content: %s",
					txHash, string(data)),
			)
		}
	}
	return failedNumbers
}
