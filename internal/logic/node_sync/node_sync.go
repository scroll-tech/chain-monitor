package nodesync

import (
	"context"
	"fmt"
	"strconv"
	"strings"
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

// NodeType represents the type of Ethereum client.
type NodeType string

// Supported node types.
const (
	// NodeTypeReth represents Reth Ethereum client.
	NodeTypeReth NodeType = "reth"
	// NodeTypeGeth represents Geth Ethereum client.
	NodeTypeGeth NodeType = "geth"
)

// Status represents the sync status of a node.
type Status struct {
	Height    uint64
	BlockHash common.Hash
}

// LogicNodeSync monitors and compares sync status of reth and geth nodes.
type LogicNodeSync struct {
	rethClient *rpc.Client
	gethClient *rpc.Client
	config     *config.NodeSyncConfig

	// consensusStatus represents the highest block height where both nodes agree.
	mu              sync.RWMutex
	consensusStatus *Status

	// Alert state tracking to avoid alert storms.
	heightDiffAlerted   bool
	hashMismatchAlerted bool
	lastMismatchHeight  uint64

	stopChan chan struct{}

	// Prometheus metrics.
	nodeSyncHeight       *prometheus.GaugeVec
	nodeSyncHeightDiff   prometheus.Gauge
	nodeSyncHashMismatch prometheus.Counter
	nodeSyncCheckTotal   prometheus.Counter
	nodeSyncCheckFailure *prometheus.CounterVec
	nodeSyncAlertTotal   *prometheus.CounterVec
}

// NewNodeSyncLogic creates a new LogicNodeSync.
func NewNodeSyncLogic(ctx context.Context, cfg *config.NodeSyncConfig) (*LogicNodeSync, error) {
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
	logic := &LogicNodeSync{
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

	return logic, nil
}

// Start begins monitoring node sync status.
func (n *LogicNodeSync) Start(ctx context.Context) {
	interval := time.Duration(n.config.CheckInterval) * time.Second
	if interval == 0 {
		interval = 10 * time.Second
	}

	log.Info("Node sync logic started", "interval", interval)

	ticker := time.NewTicker(interval)
	defer ticker.Stop()

	n.checkNodeSync(ctx)

	for {
		select {
		case <-ctx.Done():
			if ctx.Err() != nil {
				log.Error("Node sync logic canceled with error", "error", ctx.Err())
			}
			return
		case <-n.stopChan:
			log.Info("Node sync logic stopped")
			return
		case <-ticker.C:
			n.checkNodeSync(ctx)
		}
	}
}

// Stop stops the logic.
func (n *LogicNodeSync) Stop() {
	close(n.stopChan)
}

// checkNodeSync checks sync status of both nodes and updates consensus status.
func (n *LogicNodeSync) checkNodeSync(ctx context.Context) {
	n.nodeSyncCheckTotal.Inc()

	rethHeight, err := n.getNodeHeight(ctx, n.rethClient, NodeTypeReth)
	if err != nil {
		log.Error("Failed to get reth status", "error", err)
		n.nodeSyncCheckFailure.WithLabelValues(string(NodeTypeReth)).Inc()
		return
	}

	gethHeight, err := n.getNodeHeight(ctx, n.gethClient, NodeTypeGeth)
	if err != nil {
		log.Error("Failed to get geth status", "error", err)
		n.nodeSyncCheckFailure.WithLabelValues(string(NodeTypeGeth)).Inc()
		return
	}

	n.nodeSyncHeight.WithLabelValues(string(NodeTypeReth)).Set(float64(rethHeight))
	n.nodeSyncHeight.WithLabelValues(string(NodeTypeGeth)).Set(float64(gethHeight))

	var heightDiff uint64
	if rethHeight > gethHeight {
		heightDiff = rethHeight - gethHeight
	} else {
		heightDiff = gethHeight - rethHeight
	}
	n.nodeSyncHeightDiff.Set(float64(heightDiff))

	log.Info("Node sync status",
		"reth_height", rethHeight,
		"geth_height", gethHeight,
		"height_diff", heightDiff,
	)

	n.mu.Lock()
	currentlyAlerting := heightDiff > n.config.HeightDiffThreshold
	wasAlerting := n.heightDiffAlerted
	n.heightDiffAlerted = currentlyAlerting
	n.mu.Unlock()

	if currentlyAlerting && !wasAlerting {
		n.nodeSyncAlertTotal.WithLabelValues("height_diff").Inc()
		info := slack.NodeSyncHeightDiffInfo{
			RethHeight: rethHeight,
			GethHeight: gethHeight,
			Difference: heightDiff,
			Threshold:  n.config.HeightDiffThreshold,
		}
		slack.Notify(slack.MrkDwnNodeSyncHeightDiffAlert(info))
		log.Warn("Node height difference exceeds threshold",
			"reth_height", rethHeight,
			"geth_height", gethHeight,
			"diff", heightDiff,
			"threshold", n.config.HeightDiffThreshold,
		)
	} else if !currentlyAlerting && wasAlerting {
		info := slack.NodeSyncHeightDiffInfo{
			RethHeight: rethHeight,
			GethHeight: gethHeight,
			Difference: heightDiff,
			Threshold:  n.config.HeightDiffThreshold,
		}
		slack.Notify(slack.MrkDwnNodeSyncHeightDiffRecovered(info))
		log.Info("Node height difference recovered",
			"reth_height", rethHeight,
			"geth_height", gethHeight,
			"diff", heightDiff,
			"threshold", n.config.HeightDiffThreshold,
		)
	}

	consensusHeight := rethHeight
	if gethHeight < rethHeight {
		consensusHeight = gethHeight
	}

	if err := n.updateConsensusStatus(ctx, consensusHeight); err != nil {
		log.Error("Failed to update consensus status", "height", consensusHeight, "error", err)
	}
}

// getNodeHeight retrieves the latest block height of a node.
func (n *LogicNodeSync) getNodeHeight(ctx context.Context, client *rpc.Client, nodeType NodeType) (uint64, error) {
	var blockNumber string
	if err := client.CallContext(ctx, &blockNumber, "eth_blockNumber"); err != nil {
		return 0, fmt.Errorf("failed to get block number from %s: %w", nodeType, err)
	}

	height, err := strconv.ParseUint(strings.TrimPrefix(blockNumber, "0x"), 16, 64)
	if err != nil {
		return 0, fmt.Errorf("failed to parse block number from %s: %s, error: %w", nodeType, blockNumber, err)
	}

	return height, nil
}

// updateConsensusStatus verifies both nodes agree on block hash at the given height.
func (n *LogicNodeSync) updateConsensusStatus(ctx context.Context, height uint64) error {
	blockNumberHex := fmt.Sprintf("0x%x", height)

	var rethBlock struct {
		Hash common.Hash `json:"hash"`
	}
	if err := n.rethClient.CallContext(ctx, &rethBlock, "eth_getBlockByNumber", blockNumberHex, false); err != nil {
		return fmt.Errorf("failed to get reth block: %w", err)
	}

	var gethBlock struct {
		Hash common.Hash `json:"hash"`
	}
	if err := n.gethClient.CallContext(ctx, &gethBlock, "eth_getBlockByNumber", blockNumberHex, false); err != nil {
		return fmt.Errorf("failed to get geth block: %w", err)
	}

	if rethBlock.Hash != gethBlock.Hash {
		n.nodeSyncHashMismatch.Inc()

		n.mu.Lock()
		shouldAlert := !n.hashMismatchAlerted || n.lastMismatchHeight != height
		n.hashMismatchAlerted = true
		n.lastMismatchHeight = height
		n.mu.Unlock()

		if shouldAlert {
			n.nodeSyncAlertTotal.WithLabelValues("hash_mismatch").Inc()
			info := slack.NodeSyncHashMismatchInfo{
				Height:   height,
				RethHash: rethBlock.Hash,
				GethHash: gethBlock.Hash,
			}
			slack.Notify(slack.MrkDwnNodeSyncHashMismatchAlert(info))
		}

		log.Error("Block hash mismatch detected, consensus status not updated",
			"height", height,
			"reth_hash", rethBlock.Hash.Hex(),
			"geth_hash", gethBlock.Hash.Hex(),
		)

		return fmt.Errorf("hash mismatch at height %d", height)
	}

	n.mu.Lock()
	wasHashMismatchAlerting := n.hashMismatchAlerted
	n.hashMismatchAlerted = false
	n.consensusStatus = &Status{
		Height:    height,
		BlockHash: rethBlock.Hash,
	}
	n.mu.Unlock()

	if wasHashMismatchAlerting {
		slack.Notify(slack.MrkDwnNodeSyncHashMismatchRecovered(height, rethBlock.Hash))
		log.Info("Block hash mismatch recovered",
			"height", height,
			"block_hash", rethBlock.Hash.Hex(),
		)
	} else {
		log.Info("Consensus status updated",
			"height", height,
			"block_hash", rethBlock.Hash.Hex(),
		)
	}

	return nil
}

// GetMinHeight returns the consensus height where both reth and geth agree.
func (n *LogicNodeSync) GetMinHeight() (uint64, error) {
	n.mu.RLock()
	defer n.mu.RUnlock()

	if n.consensusStatus == nil {
		return 0, fmt.Errorf("consensus status not available")
	}

	return n.consensusStatus.Height, nil
}

// GetConsensusStatus returns the full consensus status.
func (n *LogicNodeSync) GetConsensusStatus() (*Status, error) {
	n.mu.RLock()
	defer n.mu.RUnlock()

	if n.consensusStatus == nil {
		return nil, fmt.Errorf("consensus status not available")
	}

	statusCopy := *n.consensusStatus
	return &statusCopy, nil
}
