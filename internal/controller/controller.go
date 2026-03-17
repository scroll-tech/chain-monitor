package controller

import (
	"sync/atomic"

	"gorm.io/gorm"

	"github.com/scroll-tech/chain-monitor/internal/config"
	nodesync "github.com/scroll-tech/chain-monitor/internal/logic/node_sync"
)

var l2CurrentMaxBlockNumber atomic.Uint64

// FinalizeBatchCtl the Finalize batch handler
var FinalizeBatchCtl *FinalizeBatchCheckController

// InitAPI init the api controller
func InitAPI(conf *config.Config, db *gorm.DB, nodeSyncLogic *nodesync.LogicNodeSync) {
	FinalizeBatchCtl = NewFinalizeBatchCheckController(conf, db, nodeSyncLogic)
}
