package controller

import (
	"chain-monitor/internal/config"
	"github.com/scroll-tech/go-ethereum/rpc"
)

// WatcherAPI watcher api, relate to l1watcher and l2watcher.
type WatcherAPI interface {
	IsReady() bool
	CurrentNumber() uint64
	RPCClient() *rpc.Client
}

// L1WatcherAPI watcher api, relate to l1watcher and l2watcher.
type L1WatcherAPI interface {
	WatcherAPI
	L1StartNumber() uint64
	L1Contracts() *config.L1Contracts
}

type L2WatcherAPI interface {
	WatcherAPI
	L2Contracts() *config.L2Contracts
}
