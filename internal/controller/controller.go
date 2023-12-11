package controller

import (
	"sync/atomic"

	"gorm.io/gorm"

	"github.com/scroll-tech/chain-monitor/internal/config"
)

var l2CurrentMaxBlockNumber atomic.Uint64

// FinalizeBatchCtl the Finalize batch handler
var FinalizeBatchCtl *FinalizeBatchCheckController

// InitAPI init the api controller
func InitAPI(conf *config.Config, db *gorm.DB) {
	FinalizeBatchCtl = NewFinalizeBatchCheckController(conf, db)
}
