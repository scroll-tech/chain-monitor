package monitor

import (
	"chain-monitor/internal/config"
	"context"
	"github.com/go-resty/resty/v2"
	"time"

	"github.com/scroll-tech/go-ethereum/log"
	"gorm.io/gorm"

	"chain-monitor/internal/controller"
	"chain-monitor/orm"
)

var (
	batchSize uint64 = 500
)

type ChainMonitor struct {
	cfg *config.MonitorConfig
	db  *gorm.DB

	notifyCli *resty.Client

	l1watcher controller.WatcherAPI
	l2watcher controller.WatcherAPI

	startNumber uint64
	safeNumber  uint64
}

// SlackNotify is used to send alert message.
func (ch *ChainMonitor) SlackNotify(msg string) {
	if ch.cfg.SlackNotify == "" {
		return
	}
	request := ch.notifyCli.R()
	_, err := request.SetBody(msg).Post(ch.cfg.SlackNotify)
	if err != nil {
		log.Error("failed to send slack notify message", "err", err)
	}
}

func NewChainMonitor(cfg *config.MonitorConfig, db *gorm.DB, l1Watcher, l2Watcher controller.WatcherAPI) (*ChainMonitor, error) {
	startNumber, err := orm.GetLatestConfirmedNumber(db)
	if err != nil {
		return nil, err
	}

	// Use resty and init it.
	cli := resty.New()
	cli.SetRetryCount(5)
	cli.SetTimeout(time.Second * 3)
	//cli.SetHeader("Content-Type", "application/json")

	monitor := &ChainMonitor{
		cfg:         cfg,
		db:          db,
		notifyCli:   cli,
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

	err := ch.db.Transaction(func(db *gorm.DB) error {
		// confirm deposit events.
		failedNumbers, err := ch.confirmDepositEvents(ctx, db, start, end)
		if err != nil {
			return err
		}
		// store
		sTx := db.Model(&orm.ChainConfirm{}).Select("deposit_status", "confirm").Where("number BETWEEN ? AND ?", start, end)
		sTx = sTx.Update("deposit_status", true).Update("confirm", true)
		if sTx.Error != nil {
			return sTx.Error
		}

		if len(failedNumbers) > 0 {
			fTx := db.Model(&orm.ChainConfirm{}).Select("deposit_status", "confirm").Where("number in ?", failedNumbers)
			fTx = fTx.Update("deposit_status", false).Update("confirm", true)
			if fTx.Error != nil {
				return fTx.Error
			}
		}

		return nil
	})
	if err != nil {
		log.Error("failed to check deposit events", "start", start, "end", end, "err", err)
		time.Sleep(time.Second * 10)
		return
	}
	ch.startNumber = end

	log.Info("confirm l2 blocks", "start", start, "end", end)
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
