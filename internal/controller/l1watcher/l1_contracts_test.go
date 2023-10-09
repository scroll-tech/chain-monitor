package l1watcher

import (
	"context"
	"math/big"
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

func TestPP(t *testing.T) {
	cfg, err := config.NewConfig("../../../config.json")
	assert.NoError(t, err)
	client, err := ethclient.Dial(cfg.L1Config.L1URL)
	assert.NoError(t, err)
	ctx := context.Background()

	latestNumber, err := client.BlockNumber(ctx)
	assert.NoError(t, err)
	t.Log(latestNumber)

	number := latestNumber
	for ; true; number-- {
		bls, err := client.BalanceAt(ctx, cfg.L1Config.L1Contracts.ScrollMessenger, big.NewInt(int64(number)))
		if err != nil {
			break
		}
		t.Log(bls.String())
	}
	t.Log(number, latestNumber-number)
}
