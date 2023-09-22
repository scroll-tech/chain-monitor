package l2watcher

import (
	"context"
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
