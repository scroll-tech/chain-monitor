package l2watcher

import (
	"chain-monitor/internal/config"
	"context"
	"encoding/json"
	"math/big"
	"testing"

	"github.com/scroll-tech/go-ethereum/common"
	"github.com/scroll-tech/go-ethereum/ethclient"
	"github.com/stretchr/testify/assert"
)

func TestBalance(t *testing.T) {
	client, err := ethclient.Dial("http://10.5.11.195:8545")
	assert.NoError(t, err)
	var balance *big.Int
	for number := 138; number <= 140; number++ {
		bls, err := client.BalanceAt(context.Background(), common.HexToAddress("0xBa50f5340FB9F3Bd074bD638c9BE13eCB36E603d"), big.NewInt(int64(number)))
		if err != nil {
			break
		}
		t.Log(bls.String())
		if balance == nil {
			balance = big.NewInt(0).Set(bls)
		} else if balance.Cmp(bls) != 0 {
			t.Log(number, balance.String(), bls.String())
			break
		}
	}
}

func TestETHBalance(t *testing.T) {
	cfg, err := config.NewConfig("../../../config.json")
	assert.NoError(t, err)
	client, err := ethclient.Dial(cfg.L2Config.L2URL)
	assert.NoError(t, err)
	ctx := context.Background()

	balance0, err := client.BalanceAt(ctx, cfg.L2Config.L2Contracts.ScrollMessenger, big.NewInt(747))
	assert.NoError(t, err)
	balance1, err := client.BalanceAt(ctx, cfg.L2Config.L2Contracts.ScrollMessenger, big.NewInt(748))
	assert.NoError(t, err)
	t.Log(balance0.Cmp(balance1))

	tx, _, err := client.TransactionByHash(ctx, common.HexToHash("0xbeace18a957ad6a4b587f088b7b66b0829b51b8d3b50b9204258d90af7027b33"))
	assert.NoError(t, err)
	data, _ := json.MarshalIndent(tx, " ", "	")
	t.Log(string(data))
	var bb = balance1.Sub(balance1, balance0)
	t.Log(bb.String())
}
