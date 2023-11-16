package slack

import (
	"testing"
	"time"

	"github.com/scroll-tech/chain-monitor/internal/config"
)

func TestAlertSlack(t *testing.T) {
	conf := config.SlackWebhookConfig{
		Channel:          "#test_chain_monitor",
		UserName:         "chain_monitor_balance",
		WebhookURL:       "https://hooks.slack.com/services/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",
		WorkerCount:      5,
		WorkerBufferSize: 100,
	}
	InitAlertSlack(&conf)
	for i := 0; i <= 10; i++ {
		Notify("test alert slack")
	}
	time.Sleep(time.Second * 2)
}
