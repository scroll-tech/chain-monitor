package controller

import (
	"context"
	"sync/atomic"

	"github.com/scroll-tech/go-ethereum/log"
	"gorm.io/gorm"

	"github.com/scroll-tech/chain-monitor/internal/config"
)

var l2CurrentMaxBlockNumber atomic.Uint64

// FinalizeBatchCtl the Finalize batch handler
var FinalizeBatchCtl *FinalizeBatchCheckController

// InitAPI init the api controller
func InitAPI(ctx context.Context, conf *config.Config, db *gorm.DB) error {
	finalizeBatchCtl, err := NewFinalizeBatchCheckController(ctx, conf, db)
	if err != nil {
		return err
	}
	FinalizeBatchCtl = finalizeBatchCtl
	if FinalizeBatchCtl.nodeSyncLogic != nil {
		go FinalizeBatchCtl.nodeSyncLogic.Start(ctx)
		log.Info("node sync logic started successfully")
	}

	return nil
}
