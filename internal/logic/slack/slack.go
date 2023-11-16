package slack

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/go-resty/resty/v2"
	"github.com/scroll-tech/go-ethereum/log"

	"github.com/scroll-tech/chain-monitor/internal/config"
	"github.com/scroll-tech/chain-monitor/internal/utils/fanout"
)

var alertSlack *AlertSlack

// AlertSlack send slack message
type AlertSlack struct {
	cfg       *config.SlackWebhookConfig
	notifyCli *resty.Client

	senderQueue chan string
	sendWorker  *fanout.Fanout
}

// InitAlertSlack init the alert slack
func InitAlertSlack(cfg *config.SlackWebhookConfig) {
	as := &AlertSlack{
		cfg:         cfg,
		senderQueue: make(chan string, cfg.WorkerBufferSize),
	}

	cli := resty.New()
	cli.SetRetryCount(5)
	cli.SetTimeout(time.Second * 3)
	as.notifyCli = cli

	as.sendWorker = fanout.New("sendWorker",
		fanout.WithWorker(cfg.WorkerCount),
		fanout.WithBuffer(cfg.WorkerBufferSize),
	)

	alertSlack = as

	go as.run()
}

// Notify a alert message to AlertSlack
func Notify(msg string) {
	alertSlack.senderQueue <- msg
}

func (as *AlertSlack) send(msg string) {
	doSendSlack := func(ctx context.Context) {
		hookContent := map[string]string{
			"channel":  as.cfg.Channel,
			"username": as.cfg.UserName,
			"text":     msg,
		}

		data, err := json.Marshal(hookContent)
		if err != nil {
			log.Error("failed to marshal hook content", "err", err)
			return
		}

		request := as.notifyCli.R().SetHeader("Content-Type", "application/x-www-form-urlencoded")
		request = request.SetFormData(map[string]string{"payload": string(data)})
		_, err = request.Post(as.cfg.WebhookURL)
		if err != nil {
			log.Error("appear error when send slack message", "err", err)
		}
	}

	if err := as.sendWorker.Do(context.Background(), doSendSlack); err != nil {
		log.Error("do send notify failed", "error", err, "msg", msg)
	}
}

func (as *AlertSlack) run() {
	defer func() {
		if err := recover(); err != nil {
			nerr := fmt.Errorf("alert slack panic error: %v", err)
			log.Warn(nerr.Error())
		}
	}()

	for senderMessage := range as.senderQueue {
		as.send(senderMessage)
	}
}
