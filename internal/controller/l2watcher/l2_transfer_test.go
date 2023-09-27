package l2watcher

import (
	"context"
	"encoding/json"
	"math/big"
	"testing"

	"github.com/scroll-tech/go-ethereum/common"
	"github.com/scroll-tech/go-ethereum/ethclient"
	"github.com/stretchr/testify/assert"

	"chain-monitor/internal/config"
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

func TestL1ETHBalance(t *testing.T) {
	cfg, err := config.NewConfig("../../../config.json")
	assert.NoError(t, err)
	client, err := ethclient.Dial(cfg.L1Config.L1URL)
	assert.NoError(t, err)
	ctx := context.Background()

	balance0, err := client.BalanceAt(ctx, cfg.L1Config.L1Contracts.ScrollMessenger, big.NewInt(4368319))
	assert.NoError(t, err)
	balance1, err := client.BalanceAt(ctx, cfg.L1Config.L1Contracts.ScrollMessenger, big.NewInt(4368320))
	assert.NoError(t, err)
	t.Log(balance0.Cmp(balance1))

	tx, _, err := client.TransactionByHash(ctx, common.HexToHash("0x191e4f409fa5f5e67ded656765d2d2d23d021b8bfa71d90bde5f0c8c107bd4c1"))
	assert.NoError(t, err)
	data, _ := json.MarshalIndent(tx, " ", "	")
	t.Log(string(data))
	var bb = balance1.Sub(balance1, balance0)
	t.Log(tx.Value().String(), bb.String())
}

func TestL2ETHBalance(t *testing.T) {
	cfg, err := config.NewConfig("../../../config.json")
	assert.NoError(t, err)
	client, err := ethclient.Dial(cfg.L2Config.L2URL)
	assert.NoError(t, err)
	ctx := context.Background()

	balance0, err := client.BalanceAt(ctx, cfg.L2Config.L2Contracts.ScrollMessenger, big.NewInt(1083713))
	assert.NoError(t, err)
	balance1, err := client.BalanceAt(ctx, cfg.L2Config.L2Contracts.ScrollMessenger, big.NewInt(1083714))
	assert.NoError(t, err)
	t.Log(balance0.Cmp(balance1))

	tx, _, err := client.TransactionByHash(ctx, common.HexToHash("0xfa4f52dd77c15cbd647dae0b939ebe95b43c789e3a4dfb12b078d69ea3223497"))
	assert.NoError(t, err)
	data, _ := json.MarshalIndent(tx, " ", "	")
	t.Log(string(data))
	var bb = balance1.Sub(balance1, balance0)
	t.Log(tx.Value().String(), bb.String())
}
