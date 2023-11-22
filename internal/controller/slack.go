package controller

import (
	"context"

	"github.com/scroll-tech/go-ethereum/log"

	"github.com/scroll-tech/chain-monitor/internal/config"
	"github.com/scroll-tech/chain-monitor/internal/logic/slack"
)

// SlackAlertController the controller of slack
type SlackAlertController struct {
	slackLogic *slack.AlertSlack
}

// NewSlackAlertController create SlackAlertController
func NewSlackAlertController(ctx context.Context, conf *config.SlackWebhookConfig) *SlackAlertController {
	return &SlackAlertController{
		slackLogic: slack.NewAlertSlack(ctx, conf),
	}
}

// Start the slack alert logic
func (s *SlackAlertController) Start() {
	log.Info("slack alert controller start successful")

	s.slackLogic.Start()
}

// Stop the slack alert logic
func (s *SlackAlertController) Stop() {
	s.slackLogic.Stop()
}
