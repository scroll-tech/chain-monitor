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
	cli, err := rpc.DialContext(ctx, "http://10.5.11.195:8545")
	assert.NoError(t, err)
	blocks := []uint64{11318, 11322}
	balances, err := GetBatchBalances(ctx, cli, common.HexToAddress("0xBa50f5340FB9F3Bd074bD638c9BE13eCB36E603d"), blocks)
	assert.NoError(t, err)
	assert.Len(t, balances, len(blocks))
	t.Log(balances)
}
