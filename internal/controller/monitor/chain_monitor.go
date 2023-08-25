package monitor

import (
	"chain-monitor/internal/config"
	"chain-monitor/orm"
	"context"
	"github.com/scroll-tech/go-ethereum/common"
	"github.com/scroll-tech/go-ethereum/log"
	"gorm.io/gorm"
	"time"
)

type L2API interface {
	StartNumber() uint64
	WithdrawRoot(ctx context.Context, number uint64) (common.Hash, error)
}

type ChainMonitor struct {
	db *gorm.DB

	bridgeHistoryURL string
	l2API            L2API
}

func NewChainMonitor(cfg *config.ChainMonitor, db *gorm.DB, l2API L2API) *ChainMonitor {
	return &ChainMonitor{
		db:               db,
		bridgeHistoryURL: cfg.BridgeHistoryUrl,
		l2API:            l2API,
	}
}

func (c *ChainMonitor) ChainMonitor(ctx context.Context) {
	tx := c.db.Begin()
	latestIndex, err := orm.GetConfirmedBatchIndex(tx)
	if err != nil {
		log.Error("failed to get latest batchIndex in BatchConfirm loop", "err", err)
		return
	}
	latestIndex += 1

	scrollChain, err := orm.GetScrollChainByIndex(tx, latestIndex)
	if err != nil {
		if err.Error() == gorm.ErrRecordNotFound.Error() {
			time.Sleep(time.Second * 3)
			log.Debug("scrollChain is not ready", "batch_index", latestIndex)
			return
		}
		log.Error("failed to get scrollChain by batchIndex", "batch_index", latestIndex, "err", err)
		return
	}

	if c.l2API.StartNumber() < scrollChain.L2EndNumber {
		time.Sleep(time.Second * 3)
		log.Debug("l2watcher is not ready", "batch_index", latestIndex)
		return
	}

	var (
		start, end = scrollChain.L2StartNumber, scrollChain.L2EndNumber
		batchIndex = scrollChain.BatchIndex
	)
	l1Count, l2Count, err := orm.GetMessengerCount(tx, start, end)
	if err != nil {
		log.Error("failed to get l1_count and l2_count", "batch_index", batchIndex, "err", err)
		return
	}
	eventStatus := l1Count == l2Count

	rootStatus, err := c.confirmWithdrawRoot(batchIndex, end)
	if err != nil {
		log.Error("failed to confirm withdraw root", "batch_index", batchIndex, "err", err)
		return
	}

	res := tx.Save(&orm.BatchConfirm{
		BatchIndex:         batchIndex,
		DepositStatus:      l1Count == l2Count,
		WithdrawRootStatus: rootStatus,
	})
	if err = res.Error; err != nil {
		tx.Rollback()
		log.Error("failed to store BatchConfirm result", "batch_index", batchIndex, "err", err)
		return
	}
	if err = tx.Commit().Error; err != nil {
		tx.Rollback()
		log.Error("failed to commit BatchConfirm status", "batch_index", batchIndex, "err", err)
		return
	}
	if !eventStatus || !rootStatus {
		log.Error("event count or withdraw root is not right", "batch_index", batchIndex, "l1_count", l1Count, "l2_count", l2Count, "withdraw_root_status", rootStatus)
	} else {
		log.Info("scan batch successful", "batch_index", batchIndex, "deposit_status", eventStatus, "withdraw_root_status", rootStatus)
	}
}
