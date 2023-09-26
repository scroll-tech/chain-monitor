package monitor

import (
	"encoding/json"
	"time"

	"github.com/go-resty/resty/v2"
	"github.com/scroll-tech/go-ethereum/log"
	"gorm.io/gorm"

	"chain-monitor/internal/config"
	"chain-monitor/internal/controller"
	"chain-monitor/internal/orm"
)

var batchSize uint64 = 500

// ChainMonitor struct represents a monitoring structure for blockchain operations.
type ChainMonitor struct {
	cfg *config.SlackWebhookConfig
	db  *gorm.DB

	alertTime time.Time

	notifyCli *resty.Client

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
func NewChainMonitor(cfg *config.SlackWebhookConfig, db *gorm.DB, l1Watcher controller.L1WatcherAPI, l2Watcher controller.WatcherAPI) (*ChainMonitor, error) {
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
		cfg:                 cfg,
		db:                  db,
		alertTime:           time.Now(),
		notifyCli:           cli,
		depositStartNumber:  depositStartNumber,
		withdrawStartNumber: withdrawStartNumber,
		l1watcher:           l1Watcher,
		l2watcher:           l2Watcher,
	}
	return monitor, nil
}

// SlackNotify sends an alert message to a Slack channel.
func (ch *ChainMonitor) SlackNotify(msg string) {
	if ch.cfg.WebhookURL == "" {
		return
	}
	curTime := time.Now()
	if curTime.Sub(ch.alertTime).Milliseconds() < 500 {
		return
	}
	ch.alertTime = curTime

	hookContent := map[string]string{
		"channel":  ch.cfg.Channel,
		"username": ch.cfg.UserName,
		"text":     msg,
	}
	data, err := json.Marshal(hookContent)
	if err != nil {
		log.Error("failed to marshal hook content", "err", err)
		return
	}

	request := ch.notifyCli.R().SetHeader("Content-Type", "application/x-www-form-urlencoded")
	request = request.SetFormData(map[string]string{"payload": string(data)})
	_, err = request.Post(ch.cfg.WebhookURL)
	if err != nil {
		log.Error("appear error when send slack message", "err", err)
	}
}
