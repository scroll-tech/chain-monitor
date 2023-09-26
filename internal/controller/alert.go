package controller

import (
	"encoding/json"
	"sync"
	"time"

	"github.com/go-resty/resty/v2"
	"github.com/scroll-tech/go-ethereum/log"

	"chain-monitor/internal/config"
)

var (
	slackAlert *webhookAlert
	once       sync.Once
)

type webhookAlert struct {
	cfg *config.SlackWebhookConfig

	notifyCli *resty.Client
	alertTime time.Time
}

// SlackNotify sends an alert message to a Slack channel.
func SlackNotify(msg string) {
	ch := slackAlert
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

// InitWebhookAlert init the independent webhook alert instance.
func InitWebhookAlert(cfg *config.SlackWebhookConfig) {
	once.Do(func() {
		// Use resty and init it.
		cli := resty.New()
		cli.SetRetryCount(5)
		cli.SetTimeout(time.Second * 3)
		slackAlert = &webhookAlert{
			cfg:       cfg,
			notifyCli: cli,
			alertTime: time.Now(),
		}
	})
}
