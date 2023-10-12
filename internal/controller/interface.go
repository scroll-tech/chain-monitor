package controller

import "chain-monitor/internal/config"

// WatcherAPI watcher api, relate to l1watcher and l2watcher.
type WatcherAPI interface {
	IsReady() bool
	CurrentNumber() uint64
	GetGatewayConfig() *config.Gateway
}

// L1WatcherAPI watcher api, relate to l1watcher and l2watcher.
type L1WatcherAPI interface {
	WatcherAPI
	L1StartNumber() uint64
}
