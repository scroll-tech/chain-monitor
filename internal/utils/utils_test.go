package utils

import (
	"context"
	"testing"

	"github.com/scroll-tech/go-ethereum/common"
	"github.com/scroll-tech/go-ethereum/rpc"
	"github.com/stretchr/testify/assert"
)

func TestGetBatchBalances(t *testing.T) {
	ctx := context.Background()
	cli, err := rpc.DialContext(ctx, "http://127.0.0.1:8545")
	assert.NoError(t, err)
	balances, err := GetBatchBalances(ctx, cli, common.HexToAddress("0xBa50f5340FB9F3Bd074bD638c9BE13eCB36E603d"), []uint64{11318, 11322})
	assert.NoError(t, err)
	t.Log(balances)
}
