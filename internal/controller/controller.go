package controller

import (
	"github.com/scroll-tech/chain-monitor/internal/config"
	"gorm.io/gorm"
)

// FinalizeBatchCtl the Finalize batch handler
var FinalizeBatchCtl *FinalizeBatchCheckController

// InitApi init the api controller
func InitApi(conf *config.Config, db *gorm.DB) {
	FinalizeBatchCtl = NewFinalizeBatchCheckController(conf, db)
}
