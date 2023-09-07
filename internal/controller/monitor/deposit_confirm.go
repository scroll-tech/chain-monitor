package monitor

import (
	"context"
	"fmt"
	"time"

	"github.com/scroll-tech/go-ethereum/log"
	"gorm.io/gorm"

	"chain-monitor/internal/orm"
)

// DepositConfirm monitors the blockchain and confirms the deposit events.
func (ch *ChainMonitor) DepositConfirm(ctx context.Context) {
	// Make sure the l1Watcher is ready to use.
	if !ch.l1watcher.IsReady() {
		log.Debug("l1watcher is not ready, sleep 3 seconds")
		time.Sleep(time.Second * 3)
		return
	}
	start, end := ch.getDepositStartAndEndNumber()
	if end > ch.depositSafeNumber {
		log.Debug("l2watcher is not ready", "l2_start_number", ch.depositSafeNumber)
		time.Sleep(time.Second * 3)
		return
	}

	err := ch.db.Transaction(func(tx *gorm.DB) error {
		// confirm deposit events.
		failedNumbers, err := ch.confirmDepositEvents(ctx, tx, start, end)
		if err != nil {
			return err
		}
		// Update deposit records.
		sTx := tx.Model(&orm.ChainConfirm{}).Select("deposit_status", "deposit_confirm").
			Where("number BETWEEN ? AND ?", start, end)
		sTx = sTx.Update("deposit_status", true).Update("deposit_confirm", true)
		if sTx.Error != nil {
			return sTx.Error
		}

		if len(failedNumbers) > 0 {
			fTx := tx.Model(&orm.ChainConfirm{}).Select("deposit_status", "deposit_confirm").
				Where("number in ?", failedNumbers)
			fTx = fTx.Update("deposit_status", false)
			if fTx.Error != nil {
				return fTx.Error
			}
		}

		return nil
	})
	if err != nil {
		log.Error("failed to check deposit events", "start", start, "end", end, "err", err)
		time.Sleep(time.Second * 5)
		return
	}
	ch.depositStartNumber = end

	log.Info("confirm layer2 deposit transactions", "start", start, "end", end)
}

func (ch *ChainMonitor) confirmDepositEvents(ctx context.Context, db *gorm.DB, start, end uint64) ([]uint64, error) {
	db = db.WithContext(ctx)
	var (
		failedNumbers []uint64
		flagNumbers   = map[uint64]bool{}
	)

	// check eth events.
	var ethEvents []msgEvents
	db = db.Raw(ethSQL, start, end, orm.L2FinalizeDepositETH)
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
			go ch.SlackNotify(fmt.Sprintf("deposit eth don't match, message: %v", msg))
			log.Error("the eth deposit count or amount don't match", "start", start, "end", end, "event_type", orm.L2FinalizeDepositETH, "l1_tx_hash", msg.L1TxHash, "l2_tx_hash", msg.L2TxHash)
		}
	}

	var erc20Events []msgEvents
	// check erc20 events.
	db = db.Raw(erc20SQL,
		start, end,
		orm.L2FinalizeDepositDAI,
		orm.L2FinalizeDepositWETH,
		orm.L2FinalizeDepositStandardERC20,
		orm.L2FinalizeDepositCustomERC20)
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
			go ch.SlackNotify(fmt.Sprintf("erc20 deposit don't match, message: %v", msg))
			log.Error(
				"the erc20 deposit count or amount doesn't match",
				"start", start,
				"end", end,
				"event_type", []orm.EventType{orm.L2FinalizeDepositDAI, orm.L2FinalizeDepositWETH, orm.L2FinalizeDepositStandardERC20, orm.L2FinalizeDepositCustomERC20},
				"l1_tx_hash", msg.L1TxHash,
				"l2_tx_hash", msg.L2TxHash,
			)
		}
	}

	// check erc721 events.
	var erc721Events []msgEvents
	db = db.Raw(erc721SQL, start, end, orm.L2FinalizeDepositERC721)
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
			go ch.SlackNotify(fmt.Sprintf("erc721 event don't match, message: %v", msg))
			log.Error("the erc721 deposit count or amount doesn't match", "start", start, "end", end, "event_type", orm.L2FinalizeDepositERC721, "l1_tx_hash", msg.L1TxHash, "l2_tx_hash", msg.L2TxHash)
		}
	}

	// check erc1155 events.
	var erc1155Events []msgEvents
	db = db.Raw(erc1155SQL, start, end, orm.L2FinalizeDepositERC1155)
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
			go ch.SlackNotify(fmt.Sprintf("erc1155 event don't match, message: %v", msg))
			log.Error("the erc1155 deposit count or amount doesn't match", "start", start, "end", end, "event_type", orm.L2FinalizeDepositERC1155, "l1_tx_hash", msg.L1TxHash, "l2_tx_hash", msg.L2TxHash)
		}
	}

	return failedNumbers, nil
}

func (ch *ChainMonitor) getDepositStartAndEndNumber() (uint64, uint64) {
	var (
		start = ch.depositStartNumber + 1
		end   = start + batchSize - 1
	)
	ch.depositSafeNumber = ch.l2watcher.StartNumber()
	if end < ch.depositSafeNumber {
		return start, end
	}
	if start < ch.depositSafeNumber {
		return start, ch.depositSafeNumber - 1
	}
	return start, start
}
