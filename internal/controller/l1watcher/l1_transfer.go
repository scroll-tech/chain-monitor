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
	"chain-monitor/internal/utils"
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

// This struct is used to check scrollMessenger eth balance.
type ethEvent struct {
	Number uint64
	TxHash string
	Type   orm.EventType
	Amount *big.Int
}

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

	var (
		total    = big.NewInt(0).Set(sBalance)
		events   = make(map[uint64][]*ethEvent)
		txHashes = make(map[string]bool)
	)
	for _, event := range l1.ethEvents {
		var amount = big.NewInt(0).Set(event.Amount)
		// L1DepositETH: +amount, L1FinalizeWithdrawETH: -amount, L1RefundETH: -amount
		if event.Type == orm.L1FinalizeWithdrawETH || event.Type == orm.L1RefundETH {
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
	for _, event := range l1.erc20Events {
		if !(event.Type == orm.L1DepositWETH ||
			event.Type == orm.L1FinalizeWithdrawWETH ||
			event.Type == orm.L1RefundWETH) {
			continue
		}
		var amount = big.NewInt(0).Set(event.Amount)
		// L1DepositWETH: +amount, L1FinalizeWithdrawWETH: -amount
		if event.Type == orm.L1FinalizeWithdrawWETH || event.Type == orm.L1RefundWETH {
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

	for _, msgList := range l1.msgSentEvents {
		for _, msg := range msgList {
			txHash := msg.Data.Log.TxHash.String()
			if !txHashes[txHash] {
				txHashes[txHash] = true
				total.Add(total, msg.Data.Value)
				events[msg.Data.Log.BlockNumber] = append(events[msg.Data.Log.BlockNumber], &ethEvent{
					Number: msg.Data.Log.BlockNumber,
					TxHash: txHash,
					Type:   msg.Type,
					Amount: big.NewInt(0).Set(msg.Data.Value),
				})
			}
		}
	}
	if total.Cmp(eBalance) == 0 {
		return 0, nil
	}

	// Get eth batch balances.
	numbers := make([]uint64, 0, end-start+1)
	for number := start; number <= end; number++ {
		numbers = append(numbers, number)
	}
	balances, err := utils.GetBatchBalances(ctx, l1.rpcCli, l1.cfg.ScrollMessenger, numbers)
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
			controller.ETHBalanceFailedTotal.WithLabelValues(l1.chainName).Inc()
			go controller.SlackNotify(fmt.Sprintf("l1ScrollMessenger eth balance mismatch appeared, number: %d, expect_balance: %s, actual_balance: %s", number, balance.String(), amount.String()))
			return number, nil
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
		for _, api := range l1.gatewayAPIs {
			addr := api.GetAddress()
			if event.To == addr || event.From == addr {
				failedNumbers = append(failedNumbers, event.Log.BlockNumber)
				controller.ERC20UnexpectTotal.WithLabelValues(l1.chainName).Inc()
				data, _ := json.Marshal(event)
				go controller.SlackNotify(
					fmt.Sprintf("l1chain unexpect tx.From or tx.To address used gateway, tx_hash: %s, content: %s", txHash, string(data)))
			}
		}
	}

	return failedNumbers
}
