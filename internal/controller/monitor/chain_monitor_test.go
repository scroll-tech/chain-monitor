package monitor

import (
	"encoding/json"
	"testing"
	"time"

	"github.com/go-resty/resty/v2"
	"github.com/stretchr/testify/assert"
)

func TestMonitor(t *testing.T) {
	cli := resty.New()
	cli.SetTimeout(time.Second * 3)
	cli.SetRetryCount(5)

	// Slack Webhook URL
	slackWebhookURL := "https://hooks.slack.com/services/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"

	// Create payload data.
	payloadData := map[string]string{
		"channel":    "#chain_monitor_test",
		"username":   "chain_monitor",
		"text":       "This is posted to <https://sepolia.etherscan.io/> and comes from a bot named webhookbot.",
		"icon_emoji": ":ghost:",
	}
	data, err := json.Marshal(payloadData)
	assert.NoError(t, err)

	payload := map[string]string{
		"payload": string(data),
	}

	// Create resty client.
	client := resty.New()

	// Send POST request to Slack Webhook
	request := client.R().SetHeader("Content-Type", "application/x-www-form-urlencoded")
	request = request.SetFormData(payload)
	_, err = request.Post(slackWebhookURL)
	assert.NoError(t, err)
}
