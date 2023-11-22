package slack

import (
	"context"
	"testing"
	"time"

	"github.com/scroll-tech/chain-monitor/internal/config"
	"github.com/scroll-tech/chain-monitor/internal/orm"
)

func TestAlertSlack(t *testing.T) {
	conf := config.SlackWebhookConfig{
		Channel:          "#test_chain_monitor",
		UserName:         "chain_monitor_balance",
		WebhookURL:       "https://hooks.slack.com/services/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",
		WorkerCount:      5,
		WorkerBufferSize: 100,
	}
	NewAlertSlack(context.Background(), &conf)
	for i := 0; i <= 2; i++ {
		messageMatch := &orm.MessageMatch{
			ID:            100000,
			MessageHash:   "0x23520a393d8171b19d71e41aa5953bcd927606434eadd6786c1e59f8cc001fe9",
			TokenType:     1,
			L1EventType:   1,
			L1BlockNumber: 100,
			L1TxHash:      "0x23520a393d8171b19d71e41aa5953bcd927606434eadd6786c1e59f8cc001fe9",
			L1TokenIds:    "100",
			L1Amounts:     "1000",
			L2EventType:   2,
			L2BlockNumber: 1000,
			L2TxHash:      "0x23520a393d8171b19d71e41aa5953bcd927606434eadd6786c1e59f8cc001fe9",
			L2TokenIds:    "100",
			L2Amounts:     "1000",
		}
		msg := MrkDwnGatewayCrossChainMessage(messageMatch, 1)
		Notify(msg)
	}
	time.Sleep(time.Second * 2)
}
