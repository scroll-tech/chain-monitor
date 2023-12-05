package controller

import (
	"github.com/scroll-tech/chain-monitor/internal/config"
	"gorm.io/gorm"
)

// FinalizeBatchCtl the Finalize batch handler
var FinalizeBatchCtl *FinalizeBatchCheckController

// InitAPI init the api controller
func InitAPI(conf *config.Config, db *gorm.DB) {
	FinalizeBatchCtl = NewFinalizeBatchCheckController(conf, db)
}
