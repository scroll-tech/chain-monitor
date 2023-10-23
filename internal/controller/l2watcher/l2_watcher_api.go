package l2watcher

import (
	"chain-monitor/internal/config"
	"github.com/scroll-tech/go-ethereum/rpc"
	"sync/atomic"
)

// CurrentNumber retrieves the current starting block number
// that the L2Watcher is tracking.
func (l2 *L2Watcher) CurrentNumber() uint64 {
	return atomic.LoadUint64(&l2.currNumber)
}

func (l2 *L2Watcher) setCurrentNumber(number uint64) {
	atomic.StoreUint64(&l2.currNumber, number)
}

// SafeNumber retrieves the current safe block number
// up to which it's considered secure to watch.
func (l2 *L2Watcher) SafeNumber() uint64 {
	return atomic.LoadUint64(&l2.safeNumber)
}

func (l2 *L2Watcher) setSafeNumber(number uint64) {
	atomic.StoreUint64(&l2.safeNumber, number)
}

// IsReady checks whether the L2Watcher is ready. It's considered ready
// when the starting block number matches the safe block number.
func (l2 *L2Watcher) IsReady() bool {
	return l2.CurrentNumber() == l2.SafeNumber()
}

// GetGatewayConfig return l2 gateway config.
func (l2 *L2Watcher) GetGatewayConfig() *config.Gateway {
	return l2.cfg.L2Contracts.Gateway
}

// RPCClient return l2chain client.
func (l2 *L2Watcher) RPCClient() *rpc.Client {
	return l2.filter.rpcCli
}

// L2Contracts return l2chain contracts config.
func (l2 *L2Watcher) L2Contracts() *config.L2Contracts {
	return l2.cfg.L2Contracts
}
