package nodesync

import (
	"context"
	"strconv"
	"strings"
	"testing"

	"github.com/scroll-tech/go-ethereum/common"
	"github.com/scroll-tech/go-ethereum/rpc"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// TestGetNodeHeightWithScrollRPC tests getNodeHeight with real Scroll RPC.
func TestGetNodeHeightWithScrollRPC(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping integration test in short mode")
	}

	ctx := context.Background()
	scrollRPC := "https://rpc.scroll.io"

	client, err := rpc.DialContext(ctx, scrollRPC)
	require.NoError(t, err, "Failed to connect to Scroll RPC")
	defer client.Close()

	logic := &LogicNodeSync{}

	height, err := logic.getNodeHeight(ctx, client, NodeTypeGeth)
	require.NoError(t, err, "Failed to get node height")

	assert.Greater(t, height, uint64(0), "Block height should be greater than 0")

	t.Logf("Scroll RPC Status: Height=%d", height)
}

// TestGetMinHeight tests the GetMinHeight function with consensus status.
func TestGetMinHeight(t *testing.T) {
	tests := []struct {
		name            string
		consensusStatus *NodeSyncStatus
		wantHeight      uint64
		wantErr         bool
		errMsg          string
	}{
		{
			name: "consensus status available",
			consensusStatus: &NodeSyncStatus{
				Height:    100,
				BlockHash: common.HexToHash("0x123"),
			},
			wantHeight: 100,
			wantErr:    false,
		},
		{
			name:            "consensus status not available",
			consensusStatus: nil,
			wantErr:         true,
			errMsg:          "consensus status not available",
		},
		{
			name: "high consensus height",
			consensusStatus: &NodeSyncStatus{
				Height:    1000000,
				BlockHash: common.HexToHash("0xabc"),
			},
			wantHeight: 1000000,
			wantErr:    false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			logic := &LogicNodeSync{
				consensusStatus: tt.consensusStatus,
			}

			height, err := logic.GetMinHeight()

			if tt.wantErr {
				assert.Error(t, err)
				if tt.errMsg != "" {
					assert.Contains(t, err.Error(), tt.errMsg)
				}
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.wantHeight, height)
			}
		})
	}
}

// TestUpdateConsensusStatusWithScrollRPC tests consensus status update with real Scroll RPC.
func TestUpdateConsensusStatusWithScrollRPC(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping integration test in short mode")
	}

	ctx := context.Background()
	scrollRPC := "https://rpc.scroll.io"

	client1, err := rpc.DialContext(ctx, scrollRPC)
	require.NoError(t, err, "Failed to connect to Scroll RPC")
	defer client1.Close()

	client2, err := rpc.DialContext(ctx, scrollRPC)
	require.NoError(t, err, "Failed to connect to Scroll RPC")
	defer client2.Close()

	logic := &LogicNodeSync{
		rethClient: client1,
		gethClient: client2,
	}

	var blockNumber string
	err = client1.CallContext(ctx, &blockNumber, "eth_blockNumber")
	require.NoError(t, err, "Failed to get block number")

	height, err := strconv.ParseUint(strings.TrimPrefix(blockNumber, "0x"), 16, 64)
	require.NoError(t, err, "Failed to parse block number")

	err = logic.updateConsensusStatus(ctx, height-10)
	assert.NoError(t, err, "Consensus status update should succeed for same RPC")

	consensusStatus, err := logic.GetConsensusStatus()
	require.NoError(t, err, "Should be able to get consensus status")
	assert.Equal(t, height-10, consensusStatus.Height, "Consensus height should match")
	assert.NotEqual(t, common.Hash{}, consensusStatus.BlockHash, "Consensus hash should not be empty")

	t.Logf("Consensus status updated: height=%d, hash=%s",
		consensusStatus.Height, consensusStatus.BlockHash.Hex())
}
