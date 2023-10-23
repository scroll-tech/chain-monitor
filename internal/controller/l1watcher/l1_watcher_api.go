package l1watcher

import (
	"chain-monitor/internal/config"
	"github.com/scroll-tech/go-ethereum/rpc"
	"sync/atomic"
)

// L1StartNumber returns l1watcher start number.
func (l1 *L1Watcher) L1StartNumber() uint64 {
	return l1.cfg.StartNumber
}

// CurrentNumber return l1watcher start number.
func (l1 *L1Watcher) CurrentNumber() uint64 {
	return atomic.LoadUint64(&l1.currNumber)
}

// GetGatewayConfig  return l1 gateway config.
func (l1 *L1Watcher) GetGatewayConfig() *config.Gateway {
	return l1.cfg.L1Contracts.Gateway
}

func (l1 *L1Watcher) setCurrentNumber(number uint64) {
	atomic.StoreUint64(&l1.currNumber, number)
}

// SafeNumber return safe number.
func (l1 *L1Watcher) SafeNumber() uint64 {
	return atomic.LoadUint64(&l1.safeNumber)
}

func (l1 *L1Watcher) setSafeNumber(number uint64) {
	atomic.StoreUint64(&l1.safeNumber, number)
}

// IsReady if l1watcher is ready return true.
func (l1 *L1Watcher) IsReady() bool {
	return l1.CurrentNumber() == l1.SafeNumber()
}

// RPCClient return l1chain client.
func (l1 *L1Watcher) RPCClient() *rpc.Client {
	return l1.filter.rpcCli
}

// L1Contracts return l1chain contracts config.
func (l1 *L1Watcher) L1Contracts() *config.L1Contracts {
	return l1.cfg.L1Contracts
}
