package l1watcher

import (
	"context"
	"testing"

	"github.com/scroll-tech/go-ethereum/ethclient"
	"github.com/stretchr/testify/assert"

	"chain-monitor/internal/config"
	"chain-monitor/internal/utils"
)

func TestL1ScrollMessenger(t *testing.T) {
	cfg, err := config.NewConfig("../../../config.json")
	assert.NoError(t, err)
	client, err := ethclient.Dial(cfg.L1Config.L1URL)
	assert.NoError(t, err)
	ctx := context.Background()

	// Create db instance.
	db, err := utils.InitDB(cfg.DBConfig)
	assert.NoError(t, err)
	defer utils.CloseDB(db)

	filter, err := newL1Contracts(cfg.L1Config.L1URL, cfg.L1Config.L1Contracts)
	assert.NoError(t, err)
	filter.registerMessengerHandlers()

	filter.gatewayFilter.GetLogs(ctx, client, 4373817, 4373818, filter.gatewayFilter.ParseLogs)

}
