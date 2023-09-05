package scroll

import (
	"context"
	"encoding/json"
	"math/big"
	"testing"

	"github.com/scroll-tech/go-ethereum/ethclient"
	"github.com/stretchr/testify/assert"
)

func TestPP(t *testing.T) {
	client, err := ethclient.Dial("http://localhost:8545")
	assert.NoError(t, err)
	client.SetHeader("Content-Encoding", "deflate")
	trace1, err := client.GetBlockTraceByNumber(context.Background(), big.NewInt(11427))
	assert.NoError(t, err)

	data1, err := json.Marshal(trace1)
	assert.NoError(t, err)
	t.Log(trace1.ChainID, len(data1))

	//
}

func TestCC(t *testing.T) {
	client2, err := ethclient.Dial("http://localhost:8545")
	assert.NoError(t, err)
	trace2, err := client2.GetBlockTraceByNumber(context.Background(), big.NewInt(11427))
	assert.NoError(t, err)

	data2, err := json.Marshal(trace2)
	assert.NoError(t, err)

	t.Log(trace2.ChainID, len(data2))
}
