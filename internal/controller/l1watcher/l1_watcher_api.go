package l1watcher

import (
	"sync/atomic"

	"chain-monitor/internal/controller"
)

// L1StartNumber returns l1watcher start number.
func (l1 *L1Watcher) L1StartNumber() uint64 {
	return l1.cfg.StartNumber
}

// CurrentNumber return l1watcher start number.
func (l1 *L1Watcher) CurrentNumber() uint64 {
	return atomic.LoadUint64(&l1.currNumber)
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

// SetMonitor sets monitor api.
func (l1 *L1Watcher) SetMonitor(monitor controller.MonitorAPI) {
	l1.filter.setMonitorAPI(monitor)
}
