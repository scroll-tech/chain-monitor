package controller

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/scroll-tech/go-ethereum/common"
	"github.com/scroll-tech/go-ethereum/rpc"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// TestGetNodeStatusWithScrollRPC tests getNodeStatus with real Scroll RPC
func TestGetNodeStatusWithScrollRPC(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping integration test in short mode")
	}

	ctx := context.Background()
	scrollRPC := "https://rpc.scroll.io"

	// Connect to Scroll RPC
	client, err := rpc.DialContext(ctx, scrollRPC)
	require.NoError(t, err, "Failed to connect to Scroll RPC")
	defer client.Close()

	controller := &NodeSyncController{}

	// Test getting node status
	status, err := controller.getNodeStatus(ctx, client, NodeTypeGeth)
	require.NoError(t, err, "Failed to get node status")

	// Verify status
	assert.Greater(t, status.Height, uint64(0), "Block height should be greater than 0")
	assert.NotEqual(t, common.Hash{}, status.BlockHash, "Block hash should not be empty")
	assert.False(t, status.Timestamp.IsZero(), "Timestamp should be set")

	t.Logf("✅ Scroll RPC Status: Height=%d, Hash=%s", status.Height, status.BlockHash.Hex())
}

// TestGetMinHeight tests the GetMinHeight function with consensus status
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
				Timestamp: time.Now(),
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
				Timestamp: time.Now(),
			},
			wantHeight: 1000000,
			wantErr:    false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			controller := &NodeSyncController{
				consensusStatus: tt.consensusStatus,
			}

			height, err := controller.GetMinHeight()

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

// TestUpdateConsensusStatusWithScrollRPC tests consensus status update with real Scroll RPC
func TestUpdateConsensusStatusWithScrollRPC(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping integration test in short mode")
	}

	ctx := context.Background()
	scrollRPC := "https://rpc.scroll.io"

	// Connect to Scroll RPC (using same RPC for both to ensure hashes match)
	client1, err := rpc.DialContext(ctx, scrollRPC)
	require.NoError(t, err, "Failed to connect to Scroll RPC")
	defer client1.Close()

	client2, err := rpc.DialContext(ctx, scrollRPC)
	require.NoError(t, err, "Failed to connect to Scroll RPC")
	defer client2.Close()

	controller := &NodeSyncController{
		rethClient: client1,
		gethClient: client2,
	}

	// Get a recent block height
	var blockNumber string
	err = client1.CallContext(ctx, &blockNumber, "eth_blockNumber")
	require.NoError(t, err, "Failed to get block number")

	var height uint64
	_, err = fmt.Sscanf(blockNumber, "0x%x", &height)
	require.NoError(t, err, "Failed to parse block number")

	// Update consensus status (should succeed since using same RPC)
	err = controller.updateConsensusStatus(ctx, height-10) // Use a confirmed block
	assert.NoError(t, err, "Consensus status update should succeed for same RPC")

	// Verify status was updated
	consensusStatus, err := controller.GetConsensusStatus()
	require.NoError(t, err, "Should be able to get consensus status")
	assert.Equal(t, height-10, consensusStatus.Height, "Consensus height should match")
	assert.NotEqual(t, common.Hash{}, consensusStatus.BlockHash, "Consensus hash should not be empty")

	t.Logf("✅ Consensus status updated: height=%d, hash=%s",
		consensusStatus.Height, consensusStatus.BlockHash.Hex())
}
