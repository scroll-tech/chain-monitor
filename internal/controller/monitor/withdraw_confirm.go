package monitor

import (
	"context"
	"fmt"
	"time"

	"github.com/scroll-tech/go-ethereum/log"
	"gorm.io/gorm"

	"chain-monitor/internal/orm"
)

var (
	l1ethSQL = `select 
    l1ee.tx_hash as l1_tx_hash, l1ee.amount as l1_amount, 
    l2ee.tx_hash as l2_tx_hash, l2ee.number as l2_number, l2ee.amount as l2_amount 
from l2_eth_events as l2ee full join l1_eth_events as l1ee 
    on l1ee.msg_hash = l2ee.msg_hash  
where l1ee.number BETWEEN ? AND ? and l1ee.type = ?;`
	l1erc20SQL = `select 
    l1ee.tx_hash as l1_tx_hash, l1ee.amount as l1_amount, 
    l2ee.tx_hash as l2_tx_hash, l2ee.number as l2_number, l2ee.amount as l2_amount 
from l2_erc20_events as l2ee full join l1_erc20_events as l1ee 
    on l1ee.msg_hash = l2ee.msg_hash  
where l1ee.number BETWEEN ? AND ? and l1ee.type in (?, ?, ?, ?);`
	l1erc721SQL = `select 
    l1ee.tx_hash as l1_tx_hash, l1ee.token_id as l1_token_id, 
    l2ee.tx_hash as l2_tx_hash, l2ee.number as l2_number, l2ee.token_id as l2_token_id
from l2_erc721_events as l2ee full join l1_erc721_events as l1ee 
    on l1ee.msg_hash = l2ee.msg_hash 
where l1ee.number BETWEEN ? AND ? and l1ee.type = ?;`
	l1erc1155SQL = `select 
    l1ee.tx_hash as l1_tx_hash, l1ee.amount as l1_amount, l1ee.token_id as l1_token_id, 
    l2ee.tx_hash as l2_tx_hash, l2ee.number as l2_number, l2ee.amount as l2_amount, l2ee.token_id as l2_token_id
from l2_erc1155_events as l2ee full join l1_erc1155_events as l1ee 
    on l1ee.msg_hash = l2ee.msg_hash 
where l1ee.number BETWEEN ? AND ? and l1ee.type = ?;`
)

// WithdrawConfirm the loop in order to confirm withdraw events.
func (ch *ChainMonitor) WithdrawConfirm(ctx context.Context) {
	// Make sure the l2Watcher is ready to use.
	if !ch.l2watcher.IsReady() {
		log.Debug("l2watcher is not ready, sleep 3 seconds")
		time.Sleep(time.Second * 3)
		return
	}
	start, end := ch.getWithdrawStartAndEndNumber()
	if end > ch.withdrawSafeNumber {
		log.Debug("l1watcher is not ready", "l1_start_number", ch.withdrawSafeNumber)
		time.Sleep(time.Second * 3)
		return
	}

	// Get unmatched withdraw event numbers.
	failedNumbers, err := ch.confirmWithdrawEvents(ctx, start, end)
	if err != nil {
		log.Error("failed to get unmatched withdraw events", "start", start, "end", end, "err", err)
		return
	}
	err = ch.db.Transaction(func(tx *gorm.DB) error {
		// Update withdraw records.
		sTx := tx.Model(&orm.L1ChainConfirm{}).Select("withdraw_status", "confirm").
			Where("number BETWEEN ? AND ?", start, end)
		sTx = sTx.Update("withdraw_status", true).Update("confirm", true)
		if sTx.Error != nil {
			return sTx.Error
		}

		// Update failed withdraw records.
		if len(failedNumbers) > 0 {
			fTx := tx.Model(&orm.L1ChainConfirm{}).Select("withdraw_status").
				Where("number in ?", failedNumbers)
			fTx = fTx.Update("withdraw_status", false)
			if fTx.Error != nil {
				return fTx.Error
			}
		}

		return nil
	})
	if err != nil {
		log.Error("failed to check withdraw events", "start", start, "end", end, "err", err)
		time.Sleep(time.Second * 5)
		return
	}
	ch.withdrawStartNumber = end

	log.Info("confirm layer1 withdraw transactions", "start", start, "end", end)
}

func (ch *ChainMonitor) confirmWithdrawEvents(ctx context.Context, start, end uint64) ([]uint64, error) {
	var (
		db            = ch.db.WithContext(ctx)
		failedNumbers []uint64
		flagNumbers   = map[uint64]bool{}
	)

	// check eth events.
	var ethEvents []msgEvents
	db = db.Raw(l1ethSQL, start, end, orm.L1FinalizeWithdrawETH)
	if err := db.Scan(&ethEvents).Error; err != nil {
		return nil, err
	}
	for i := 0; i < len(ethEvents); i++ {
		msg := ethEvents[i]
		if msg.L1Amount != msg.L2Amount {
			if !flagNumbers[msg.L2Number] {
				flagNumbers[msg.L2Number] = true
				failedNumbers = append(failedNumbers, msg.L2Number)
			}
			// If eth msg don't match, alert it.
			go ch.SlackNotify(fmt.Sprintf("eth withdraw doesn't match, message: %v", msg))
			log.Error("the eth withdraw count or amount doesn't match", "start", start, "end", end, "event_type", orm.L1FinalizeWithdrawETH, "l1_tx_hash", msg.L1TxHash, "l2_tx_hash", msg.L2TxHash)
		}
	}

	var erc20Events []msgEvents
	// check erc20 events.
	db = db.Raw(l1erc20SQL,
		start, end,
		orm.L1FinalizeWithdrawDAI,
		orm.L1FinalizeWithdrawWETH,
		orm.L1FinalizeWithdrawStandardERC20,
		orm.L1FinalizeWithdrawCustomERC20)
	if err := db.Scan(&erc20Events).Error; err != nil {
		return nil, err
	}
	for i := 0; i < len(erc20Events); i++ {
		msg := erc20Events[i]
		if msg.L1Amount != msg.L2Amount {
			if !flagNumbers[msg.L2Number] {
				flagNumbers[msg.L2Number] = true
				failedNumbers = append(failedNumbers, msg.L2Number)
			}
			// If erc20 msg don't match, alert it.
			go ch.SlackNotify(fmt.Sprintf("erc20 withdraw doesn't match, message: %v", msg))
			log.Error(
				"the erc20 withdraw count or amount doesn't match",
				"start", start,
				"end", end,
				"event_type", []orm.EventType{orm.L1FinalizeWithdrawDAI, orm.L1FinalizeWithdrawWETH, orm.L1FinalizeWithdrawStandardERC20, orm.L1FinalizeWithdrawCustomERC20},
				"l1_tx_hash", msg.L1TxHash,
				"l2_tx_hash", msg.L2TxHash,
			)
		}
	}

	// check erc721 events.
	var erc721Events []msgEvents
	db = db.Raw(l1erc721SQL, start, end, orm.L1FinalizeWithdrawERC721)
	if err := db.Scan(&erc721Events).Error; err != nil {
		return nil, err
	}
	for i := 0; i < len(erc721Events); i++ {
		msg := erc721Events[i]
		if msg.L1TokenID != msg.L2TokenID {
			if !flagNumbers[msg.L2Number] {
				flagNumbers[msg.L2Number] = true
				failedNumbers = append(failedNumbers, msg.L2Number)
			}
			// If erc721 event don't match, alert it.
			go ch.SlackNotify(fmt.Sprintf("erc721 withdraw doesn't match, message: %v", msg))
			log.Error("the erc721 withdraw count or amount doesn't match", "start", start, "end", end, "event_type", orm.L1FinalizeWithdrawERC721, "l1_tx_hash", msg.L1TxHash, "l2_tx_hash", msg.L2TxHash)
		}
	}

	// check erc1155 events.
	var erc1155Events []msgEvents
	db = db.Raw(l1erc1155SQL, start, end, orm.L1FinalizeWithdrawERC1155)
	if err := db.Scan(&erc1155Events).Error; err != nil {
		return nil, err
	}
	for i := 0; i < len(erc1155Events); i++ {
		msg := erc1155Events[i]
		if msg.L1TokenID != msg.L2TokenID || msg.L1Amount != msg.L2Amount {
			if !flagNumbers[msg.L2Number] {
				flagNumbers[msg.L2Number] = true
				failedNumbers = append(failedNumbers, msg.L2Number)
			}
			// If erc1155 event don't match, alert it.
			go ch.SlackNotify(fmt.Sprintf("erc1155 withdraw doesn't match, message: %v", msg))
			log.Error("the erc1155 withdraw count or amount doesn't match", "start", start, "end", end, "event_type", orm.L1FinalizeWithdrawERC1155, "l1_tx_hash", msg.L1TxHash, "l2_tx_hash", msg.L2TxHash)
		}
	}

	return failedNumbers, nil
}

func (ch *ChainMonitor) getWithdrawStartAndEndNumber() (uint64, uint64) {
	var (
		start = ch.withdrawStartNumber + 1
		end   = start + batchSize - 1
	)
	ch.withdrawSafeNumber = ch.l1watcher.CurrentNumber()
	if end < ch.withdrawSafeNumber {
		return start, end
	}
	if start < ch.withdrawSafeNumber {
		return start, ch.withdrawSafeNumber - 1
	}
	return start, start
}
