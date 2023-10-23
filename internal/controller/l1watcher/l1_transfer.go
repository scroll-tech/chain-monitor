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

// This struct is used to check scrollMessenger eth balance.
type ethEvent struct {
	Number uint64
	TxHash string
	Type   orm.EventType
	Amount *big.Int
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
