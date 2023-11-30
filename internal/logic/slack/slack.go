package slack

import (
	"context"
	"encoding/json"
	"time"

	"github.com/go-resty/resty/v2"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/scroll-tech/go-ethereum/log"

	"github.com/scroll-tech/chain-monitor/internal/config"
	"github.com/scroll-tech/chain-monitor/internal/utils/fanout"
)

var alertSlack *AlertSlack

// AlertSlack send slack message
type AlertSlack struct {
	cfg       *config.SlackWebhookConfig
	notifyCli *resty.Client

	ctx             context.Context
	senderQueue     chan string
	sendWorker      *fanout.Fanout
	stopTimeoutChan chan struct{}

	alertSlackRunningTotal prometheus.Counter
}

// NewAlertSlack init the alert slack
func NewAlertSlack(ctx context.Context, cfg *config.SlackWebhookConfig) *AlertSlack {
	as := &AlertSlack{
		ctx:             ctx,
		cfg:             cfg,
		senderQueue:     make(chan string, cfg.WorkerBufferSize),
		stopTimeoutChan: make(chan struct{}),
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

	as.alertSlackRunningTotal = promauto.With(prometheus.DefaultRegisterer).NewCounter(prometheus.CounterOpts{
		Name: "alert_slack_running_total",
		Help: "The total number of alert slack running.",
	})

	return as
}

// Start the slack alert
func (as *AlertSlack) Start() {
	go as.run()
}

// Stop the slack alert
func (as *AlertSlack) Stop() {
	as.stopTimeoutChan <- struct{}{}
}

// Notify a alert message to AlertSlack
func Notify(msg string) {
	alertSlack.senderQueue <- msg
}

func (as *AlertSlack) send(msg string) {
	doSendSlack := func(ctx context.Context) {
		hookContent := map[string]string{
			"types": "mrkdwn",
			"text":  msg,
		}

		data, err := json.Marshal(hookContent)
		if err != nil {
			log.Error("failed to marshal hook content", "err", err)
			return
		}

		request := as.notifyCli.R().SetHeader("Content-Type", "application/json")
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
	for {
		as.alertSlackRunningTotal.Inc()

		select {
		case senderMessage := <-as.senderQueue:
			as.send(senderMessage)
		case <-as.ctx.Done():
			if as.ctx.Err() != nil {
				log.Error("alert slack canceled with error", "error", as.ctx.Err())
			}
			return
		case <-as.stopTimeoutChan:
			log.Info("alert slack the run loop exit")
			return
		}
	}
}
