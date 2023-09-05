package l1watcher

import (
	"sync/atomic"

	"chain-monitor/internal/controller"
)

func (l1 *L1Watcher) StartNumber() uint64 {
	return atomic.LoadUint64(&l1.startNumber)
}

func (l1 *L1Watcher) setStartNumber(number uint64) {
	atomic.StoreUint64(&l1.startNumber, number)
}

func (l1 *L1Watcher) SafeNumber() uint64 {
	return atomic.LoadUint64(&l1.safeNumber)
}

func (l1 *L1Watcher) setSafeNumber(number uint64) {
	atomic.StoreUint64(&l1.safeNumber, number)
}

func (l1 *L1Watcher) IsReady() bool {
	return l1.StartNumber() == l1.SafeNumber()
}

func (l1 *L1Watcher) SetMonitor(monitor controller.MonitorAPI) {
	l1.filter.setMonitorAPI(monitor)
}
