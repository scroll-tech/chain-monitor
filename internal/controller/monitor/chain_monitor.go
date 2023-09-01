package monitor

import (
	"context"
	"time"

	"github.com/scroll-tech/go-ethereum/common"
	"github.com/scroll-tech/go-ethereum/log"
	"gorm.io/gorm"

	"chain-monitor/internal/config"
	"chain-monitor/orm"
)

type L2API interface {
	StartNumber() uint64
	WithdrawRoot(ctx context.Context, number uint64) (common.Hash, error)
}

type ChainMonitor struct {
	db *gorm.DB

	startNumber      uint64
	bridgeHistoryURL string
	l2API            L2API
}

func NewChainMonitor(cfg *config.ChainMonitor, db *gorm.DB, l2API L2API) (*ChainMonitor, error) {
	startNumber, err := orm.GetLatestConfirmedNumber(db)
	if err != nil {
		return nil, err
	}
	monitor := &ChainMonitor{
		db:               db,
		startNumber:      startNumber,
		bridgeHistoryURL: cfg.BridgeHistoryUrl,
		l2API:            l2API,
	}
	return monitor, nil
}

func (ch *ChainMonitor) ChainMonitor(ctx context.Context) {
	startNumber := ch.startNumber + 1

	// Make sure scan number is ready.
	if ch.l2API.StartNumber() < startNumber {
		time.Sleep(time.Second * 3)
		log.Debug("l2watcher is not ready", "batch_index", startNumber)
		return
	}

	err := ch.db.Transaction(func(tx *gorm.DB) error {
		// confirm deposit events.
		isOk, err := ch.confirmDepositEvents(ctx, tx, startNumber)
		if err != nil {
			return err
		}
		tx = tx.Model(&orm.ChainConfirm{}).Select("deposit_status").Where("number = ?", startNumber)
		tx = tx.Update("deposit_status", isOk)
		tx = tx.Update("confirm", true)
		return tx.Error
	})
	if err != nil {
		log.Error("failed to check deposit events", "number", startNumber, "err", err)
		return
	}

	ch.startNumber = startNumber
}
