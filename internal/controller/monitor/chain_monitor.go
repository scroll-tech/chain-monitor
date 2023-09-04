package monitor

import (
	"context"
	"time"

	"github.com/scroll-tech/go-ethereum/log"
	"gorm.io/gorm"

	"chain-monitor/orm"
)

var (
	batchSize uint64 = 100
)

type WatcherAPI interface {
	IsReady() bool
	StartNumber() uint64
}

type ChainMonitor struct {
	db *gorm.DB

	l1watcher WatcherAPI
	l2watcher WatcherAPI

	startNumber uint64
	safeNumber  uint64
}

func NewChainMonitor(db *gorm.DB, l1Watcher, l2Watcher WatcherAPI) (*ChainMonitor, error) {
	startNumber, err := orm.GetLatestConfirmedNumber(db)
	if err != nil {
		return nil, err
	}
	monitor := &ChainMonitor{
		db:          db,
		startNumber: startNumber,
		l1watcher:   l1Watcher,
		l2watcher:   l2Watcher,
	}
	return monitor, nil
}

func (ch *ChainMonitor) ChainMonitor(ctx context.Context) {
	// Make sure the l1Watcher is ready to use.
	if !ch.l1watcher.IsReady() {
		log.Debug("l1watcher is not ready, sleep 3 seconds")
		time.Sleep(time.Second * 5)
		return
	}
	start, end := ch.getStartAndEndNumber()
	if end > ch.safeNumber {
		log.Debug("l2watcher is not ready", "l2_start_number", ch.safeNumber)
		time.Sleep(time.Second * 3)
		return
	}

	// Make sure scan number is ready.
	l2Number := ch.l2watcher.StartNumber()
	if l2Number <= ch.startNumber {
		log.Debug("l2watcher is not ready", "l2_start_number", l2Number)
		time.Sleep(time.Second * 3)
		return
	}

	err := ch.db.Transaction(func(tx *gorm.DB) error {
		// confirm deposit events.
		isOk, err := ch.confirmDepositEvents(ctx, tx, start, end)
		if err != nil {
			return err
		}
		tx = tx.Model(&orm.ChainConfirm{}).Select("deposit_status").Where("number BETWEEN ? AND ?", start, end)
		tx = tx.Update("deposit_status", isOk)
		tx = tx.Update("confirm", true)
		return tx.Error
	})
	if err != nil {
		log.Error("failed to check deposit events", "start", start, "end", end, "err", err)
		time.Sleep(time.Second * 10)
		return
	}
	ch.startNumber = end
}

func (ch *ChainMonitor) getStartAndEndNumber() (uint64, uint64) {
	var (
		start = ch.startNumber + 1
		end   = start + batchSize - 1
	)
	ch.safeNumber = ch.l2watcher.StartNumber()
	if end < ch.safeNumber {
		return start, end
	}
	if start < ch.safeNumber {
		return start, ch.safeNumber - 1
	}
	return start, start
}
