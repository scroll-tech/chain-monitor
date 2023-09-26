package monitor

import (
	"time"

	"github.com/go-resty/resty/v2"
	"gorm.io/gorm"

	"chain-monitor/internal/controller"
	"chain-monitor/internal/orm"
)

var batchSize uint64 = 500

// ChainMonitor struct represents a monitoring structure for blockchain operations.
type ChainMonitor struct {
	db *gorm.DB

	l1watcher controller.L1WatcherAPI
	l2watcher controller.WatcherAPI

	// Used for deposit confirm loop.
	depositStartNumber uint64
	depositSafeNumber  uint64

	// Used for withdraw confirm loop.
	withdrawStartNumber uint64
	withdrawSafeNumber  uint64
}

// NewChainMonitor initializes a new instance of the ChainMonitor.
func NewChainMonitor(db *gorm.DB, l1Watcher controller.L1WatcherAPI, l2Watcher controller.WatcherAPI) (*ChainMonitor, error) {
	depositStartNumber, err := orm.GetL2ConfirmNumber(db)
	if err != nil {
		return nil, err
	}
	withdrawStartNumber, err := orm.GetL1ConfirmNumber(db)
	if err != nil {
		return nil, err
	}
	if withdrawStartNumber == 0 {
		withdrawStartNumber = l1Watcher.L1StartNumber()
	}

	// Use resty and init it.
	cli := resty.New()
	cli.SetRetryCount(5)
	cli.SetTimeout(time.Second * 3)

	monitor := &ChainMonitor{
		db:                  db,
		depositStartNumber:  depositStartNumber,
		withdrawStartNumber: withdrawStartNumber,
		l1watcher:           l1Watcher,
		l2watcher:           l2Watcher,
	}
	return monitor, nil
}
