package controller

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/scroll-tech/go-ethereum/common"
	"github.com/scroll-tech/go-ethereum/log"
	"github.com/scroll-tech/go-ethereum/rpc"

	"github.com/scroll-tech/chain-monitor/internal/config"
	"github.com/scroll-tech/chain-monitor/internal/logic/slack"
)

// NodeType represents the type of Ethereum client
type NodeType string

const (
	NodeTypeReth NodeType = "reth"
	NodeTypeGeth NodeType = "geth"
)

// NodeSyncStatus represents the sync status of a node
type NodeSyncStatus struct {
	Height    uint64
	BlockHash common.Hash
	Timestamp time.Time
}

// NodeSyncController monitors and compares sync status of reth and geth nodes
type NodeSyncController struct {
	rethClient *rpc.Client
	gethClient *rpc.Client
	config     *config.NodeSyncConfig

	// Consensus status
	mu sync.RWMutex
	// consensusStatus represents the highest block height where both reth and geth agree
	// Only updated when both nodes have matching block hash at the same height
	consensusStatus *NodeSyncStatus

	stopChan chan struct{}

	// Prometheus metrics
	nodeSyncHeight       *prometheus.GaugeVec
	nodeSyncHeightDiff   prometheus.Gauge
	nodeSyncHashMismatch prometheus.Counter
	nodeSyncCheckTotal   prometheus.Counter
	nodeSyncCheckFailure *prometheus.CounterVec
	nodeSyncAlertTotal   *prometheus.CounterVec
}

// NewNodeSyncController creates a new NodeSyncController
func NewNodeSyncController(ctx context.Context, cfg *config.NodeSyncConfig) (*NodeSyncController, error) {
	if cfg == nil {
		return nil, fmt.Errorf("node sync config is nil")
	}

	rethClient, err := rpc.DialContext(ctx, cfg.RethURL)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to reth node: %w", err)
	}

	gethClient, err := rpc.DialContext(ctx, cfg.GethURL)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to geth node: %w", err)
	}

	reg := prometheus.DefaultRegisterer
	controller := &NodeSyncController{
		rethClient: rethClient,
		gethClient: gethClient,
		config:     cfg,
		stopChan:   make(chan struct{}),

		nodeSyncHeight: promauto.With(reg).NewGaugeVec(prometheus.GaugeOpts{
			Name: "node_sync_height",
			Help: "Current block height of each node type",
		}, []string{"node_type"}),

		nodeSyncHeightDiff: promauto.With(reg).NewGauge(prometheus.GaugeOpts{
			Name: "node_sync_height_diff",
			Help: "Height difference between reth and geth nodes",
		}),

		nodeSyncHashMismatch: promauto.With(reg).NewCounter(prometheus.CounterOpts{
			Name: "node_sync_hash_mismatch_total",
			Help: "Total number of block hash mismatches between reth and geth",
		}),

		nodeSyncCheckTotal: promauto.With(reg).NewCounter(prometheus.CounterOpts{
			Name: "node_sync_check_total",
			Help: "Total number of node sync checks performed",
		}),

		nodeSyncCheckFailure: promauto.With(reg).NewCounterVec(prometheus.CounterOpts{
			Name: "node_sync_check_failure_total",
			Help: "Total number of node sync check failures",
		}, []string{"node_type"}),

		nodeSyncAlertTotal: promauto.With(reg).NewCounterVec(prometheus.CounterOpts{
			Name: "node_sync_alert_total",
			Help: "Total number of node sync alerts sent",
		}, []string{"alert_type"}),
	}

	return controller, nil
}

// Start begins monitoring node sync status
func (n *NodeSyncController) Start(ctx context.Context) {
	interval := time.Duration(n.config.CheckInterval) * time.Second
	if interval == 0 {
		interval = 10 * time.Second // default 10 seconds
	}

	log.Info("Node sync controller started", "interval", interval)

	ticker := time.NewTicker(interval)
	defer ticker.Stop()

	// Do initial check immediately
	n.checkNodeSync(ctx)

	for {
		select {
		case <-ctx.Done():
			if ctx.Err() != nil {
				log.Error("Node sync controller canceled with error", "error", ctx.Err())
			}
			return
		case <-n.stopChan:
			log.Info("Node sync controller stopped")
			return
		case <-ticker.C:
			n.checkNodeSync(ctx)
		}
	}
}

// Stop stops the controller
func (n *NodeSyncController) Stop() {
	close(n.stopChan)
}

// checkNodeSync checks sync status of both nodes and updates consensus status
func (n *NodeSyncController) checkNodeSync(ctx context.Context) {
	n.nodeSyncCheckTotal.Inc()

	// Get reth status
	rethStatus, err := n.getNodeStatus(ctx, n.rethClient, NodeTypeReth)
	if err != nil {
		log.Error("Failed to get reth status", "error", err)
		n.nodeSyncCheckFailure.WithLabelValues(string(NodeTypeReth)).Inc()
		return
	}

	// Get geth status
	gethStatus, err := n.getNodeStatus(ctx, n.gethClient, NodeTypeGeth)
	if err != nil {
		log.Error("Failed to get geth status", "error", err)
		n.nodeSyncCheckFailure.WithLabelValues(string(NodeTypeGeth)).Inc()
		return
	}

	// Update metrics (using local variables, no need to store)
	n.nodeSyncHeight.WithLabelValues(string(NodeTypeReth)).Set(float64(rethStatus.Height))
	n.nodeSyncHeight.WithLabelValues(string(NodeTypeGeth)).Set(float64(gethStatus.Height))

	// Calculate height difference for monitoring
	var heightDiff uint64
	if rethStatus.Height > gethStatus.Height {
		heightDiff = rethStatus.Height - gethStatus.Height
	} else {
		heightDiff = gethStatus.Height - rethStatus.Height
	}
	n.nodeSyncHeightDiff.Set(float64(heightDiff))

	log.Info("Node sync status",
		"reth_height", rethStatus.Height,
		"geth_height", gethStatus.Height,
		"height_diff", heightDiff,
	)

	// Alert if height difference exceeds threshold
	if heightDiff > n.config.HeightDiffThreshold {
		n.nodeSyncAlertTotal.WithLabelValues("height_diff").Inc()
		alertMsg := fmt.Sprintf(
			"⚠️ *Node Height Difference Alert*\n"+
				"Reth Height: `%d`\n"+
				"Geth Height: `%d`\n"+
				"Difference: `%d` (Threshold: `%d`)",
			rethStatus.Height,
			gethStatus.Height,
			heightDiff,
			n.config.HeightDiffThreshold,
		)
		slack.Notify(alertMsg)
		log.Warn("Node height difference exceeds threshold",
			"reth_height", rethStatus.Height,
			"geth_height", gethStatus.Height,
			"diff", heightDiff,
			"threshold", n.config.HeightDiffThreshold,
		)
	}

	// Find consensus height: use the lower height and verify hash consistency
	consensusHeight := rethStatus.Height
	if gethStatus.Height < rethStatus.Height {
		consensusHeight = gethStatus.Height
	}

	// Verify both nodes agree on the block hash at consensus height
	if err := n.updateConsensusStatus(ctx, consensusHeight); err != nil {
		log.Error("Failed to update consensus status", "height", consensusHeight, "error", err)
	}
}

// getNodeStatus retrieves the current sync status of a node
func (n *NodeSyncController) getNodeStatus(ctx context.Context, client *rpc.Client, nodeType NodeType) (*NodeSyncStatus, error) {
	var blockNumber string
	if err := client.CallContext(ctx, &blockNumber, "eth_blockNumber"); err != nil {
		return nil, fmt.Errorf("failed to get block number from %s: %w", nodeType, err)
	}

	var height uint64
	_, err := fmt.Sscanf(blockNumber, "0x%x", &height)
	if err != nil {
		return nil, fmt.Errorf("failed to parse block number from %s: %s, error: %w", nodeType, blockNumber, err)
	}

	// Get block hash
	var block struct {
		Hash common.Hash `json:"hash"`
	}
	if err := client.CallContext(ctx, &block, "eth_getBlockByNumber", blockNumber, false); err != nil {
		return nil, fmt.Errorf("failed to get block from %s: %w", nodeType, err)
	}

	return &NodeSyncStatus{
		Height:    height,
		BlockHash: block.Hash,
		Timestamp: time.Now(),
	}, nil
}

// updateConsensusStatus verifies both nodes agree on block hash at given height
// and updates the consensus status only if they match
func (n *NodeSyncController) updateConsensusStatus(ctx context.Context, height uint64) error {
	blockNumberHex := fmt.Sprintf("0x%x", height)

	// Get block from reth
	var rethBlock struct {
		Hash common.Hash `json:"hash"`
	}
	if err := n.rethClient.CallContext(ctx, &rethBlock, "eth_getBlockByNumber", blockNumberHex, false); err != nil {
		return fmt.Errorf("failed to get reth block: %w", err)
	}

	// Get block from geth
	var gethBlock struct {
		Hash common.Hash `json:"hash"`
	}
	if err := n.gethClient.CallContext(ctx, &gethBlock, "eth_getBlockByNumber", blockNumberHex, false); err != nil {
		return fmt.Errorf("failed to get geth block: %w", err)
	}

	// Check if hashes match
	if rethBlock.Hash != gethBlock.Hash {
		// Hash mismatch - do NOT update consensus status
		n.nodeSyncHashMismatch.Inc()
		n.nodeSyncAlertTotal.WithLabelValues("hash_mismatch").Inc()

		alertMsg := fmt.Sprintf(
			"🚨 *Block Hash Mismatch Alert*\n"+
				"Block Height: `%d`\n"+
				"Reth Hash: `%s`\n"+
				"Geth Hash: `%s`\n"+
				"⚠️ Consensus status NOT updated",
			height,
			rethBlock.Hash.Hex(),
			gethBlock.Hash.Hex(),
		)
		slack.Notify(alertMsg)

		log.Error("Block hash mismatch detected, consensus status not updated",
			"height", height,
			"reth_hash", rethBlock.Hash.Hex(),
			"geth_hash", gethBlock.Hash.Hex(),
		)

		return fmt.Errorf("hash mismatch at height %d", height)
	}

	// Hashes match - update consensus status
	n.mu.Lock()
	n.consensusStatus = &NodeSyncStatus{
		Height:    height,
		BlockHash: rethBlock.Hash, // Both have same hash
		Timestamp: time.Now(),
	}
	n.mu.Unlock()

	log.Info("Consensus status updated",
		"height", height,
		"block_hash", rethBlock.Hash.Hex(),
	)

	return nil
}

// GetMinHeight returns the consensus height where both reth and geth agree
// This is the highest block height where both nodes have matching block hash
// This is the PRIMARY method for external queries (e.g., batch status API)
func (n *NodeSyncController) GetMinHeight() (uint64, error) {
	n.mu.RLock()
	defer n.mu.RUnlock()

	if n.consensusStatus == nil {
		return 0, fmt.Errorf("consensus status not available, nodes may not be in sync yet")
	}

	return n.consensusStatus.Height, nil
}

// GetConsensusStatus returns the full consensus status
func (n *NodeSyncController) GetConsensusStatus() (*NodeSyncStatus, error) {
	n.mu.RLock()
	defer n.mu.RUnlock()

	if n.consensusStatus == nil {
		return nil, fmt.Errorf("consensus status not available")
	}

	// Return a copy to avoid external modification
	statusCopy := *n.consensusStatus
	return &statusCopy, nil
}
