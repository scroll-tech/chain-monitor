package l2watcher

import (
	"sync/atomic"

	"chain-monitor/internal/controller"
)

func (l2 *L2Watcher) StartNumber() uint64 {
	return atomic.LoadUint64(&l2.startNumber)
}

func (l2 *L2Watcher) setStartNumber(number uint64) {
	atomic.StoreUint64(&l2.startNumber, number)
}

func (l2 *L2Watcher) SafeNumber() uint64 {
	return atomic.LoadUint64(&l2.safeNumber)
}

func (l2 *L2Watcher) setSafeNumber(number uint64) {
	atomic.StoreUint64(&l2.safeNumber, number)
}

func (l2 *L2Watcher) IsReady() bool {
	return l2.StartNumber() == l2.SafeNumber()
}

func (l2 *L2Watcher) SetMonitor(monitor controller.MonitorAPI) {
	l2.filter.setMonitorAPI(monitor)
}
