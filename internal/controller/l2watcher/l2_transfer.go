package l2watcher

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
	"chain-monitor/internal/utils"
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

func (l2 *l2Contracts) checkETHBalance(ctx context.Context, start, end uint64) (uint64, error) {
	if len(l2.msgSentEvents) == 0 {
		return 0, nil
	}

	// Get balance at start, end number.
	balances, err := utils.GetBatchBalances(ctx, l2.rpcCli, l2.cfg.ScrollMessenger, []uint64{start, end})
	if err != nil {
		return 0, err
	}
	sBalance, eBalance := balances[0], balances[1]

	var (
		total    = big.NewInt(0).Set(sBalance)
		events   = make(map[uint64][]*ethEvent)
		txHashes = make(map[string]bool)
	)
	for _, event := range l2.ethEvents {
		var amount = big.NewInt(0).Set(event.Amount)
		// L2WithdrawETH: +amount, L2FinalizeDepositETH: -amount
		if event.Type == orm.L2FinalizeDepositETH {
			amount.Mul(amount, big.NewInt(-1))
		}
		total.Add(total, amount)
		events[event.Number] = append(events[event.Number], &ethEvent{
			Number: event.Number,
			TxHash: event.TxHash,
			Type:   event.Type,
			Amount: amount,
		})
		txHashes[event.TxHash] = true
	}
	for _, event := range l2.erc20Events {
		if !(event.Type == orm.L2FinalizeDepositWETH || event.Type == orm.L2WithdrawWETH) {
			continue
		}
		var amount = big.NewInt(0).Set(event.Amount)
		// L2WithdrawWETH: +amount, L2FinalizeDepositWETH: -amount
		if event.Type == orm.L2FinalizeDepositWETH {
			amount.Mul(amount, big.NewInt(-1))
		}
		total.Add(total, amount)
		events[event.Number] = append(events[event.Number], &ethEvent{
			Number: event.Number,
			TxHash: event.TxHash,
			Type:   event.Type,
			Amount: amount,
		})
		txHashes[event.TxHash] = true
	}

	for _, msgList := range l2.msgSentEvents {
		for _, msg := range msgList {
			txHash := msg.Data.Log.TxHash.String()
			if !txHashes[txHash] {
				txHashes[txHash] = true
				total.Add(total, msg.Data.Value)
				events[msg.Number] = append(events[msg.Number], &ethEvent{
					Number: msg.Number,
					TxHash: txHash,
					Type:   msg.Type,
					Amount: big.NewInt(0).Set(msg.Data.Value),
				})
			}
		}
	}
	if total.Cmp(eBalance) <= 0 {
		return 0, nil
	}

	// Get eth batch balances.
	numbers := make([]uint64, 0, end-start+1)
	for number := start; number <= end; number++ {
		numbers = append(numbers, number)
	}
	balances, err = utils.GetBatchBalances(ctx, l2.rpcCli, l2.cfg.ScrollMessenger, numbers)
	if err != nil {
		return 0, err
	}

	var amount = big.NewInt(0).Set(sBalance)
	for idx, number := range numbers {
		for _, event := range events[number] {
			amount.Add(amount, event.Amount)
		}

		balance := balances[idx]
		if amount.Cmp(balance) != 0 {
			controller.ETHBalanceFailedTotal.WithLabelValues(l2.chainName).Inc()
			go controller.SlackNotify(fmt.Sprintf("l2ScrollMessenger eth balance mismatch appeared, number: %d, expect_balance: %s, actual_balance: %s", number, balance.String(), amount.String()))
			return number, nil
		}
	}

	return 0, nil
}

func (l2 *l2Contracts) checkL2Balance(ctx context.Context, start, end uint64) error {
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
	failedNumber, err := l2.checkETHBalance(ctx, start, end)
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
		// check to address
		for _, api := range l2.gatewayAPIs {
			addr := api.GetAddress()
			if event.To == addr || event.From == addr {
				failedNumbers = append(failedNumbers, event.Log.BlockNumber)
				controller.ERC20UnexpectTotal.WithLabelValues(l2.chainName).Inc()
				data, _ := json.Marshal(event)
				go controller.SlackNotify(
					fmt.Sprintf("l2chain unexpect tx.From or tx.To address used gateway, tx_hash: %x, content: %s",
						txHash, string(data)),
				)
			}
		}
	}
	return failedNumbers
}
