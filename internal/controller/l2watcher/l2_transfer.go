package l2watcher

import (
	"encoding/json"
	"fmt"
	"math/big"

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

func (l2 *l2Contracts) checkL2Balance() {
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

	// Update balance_status
	for number, cfm := range l2.l2Confirms {
		cfm.BalanceStatus = !failedNumbers[number]
	}
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
