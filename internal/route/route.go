package route

import (
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/scroll-tech/chain-monitor/internal/controller"
	"github.com/scroll-tech/chain-monitor/internal/utils/observability"
)

// Route register route for coordinator
func Route(router *gin.Engine) {
	router.Use(gin.Recovery())

	observability.Use(router, "chain_monitor", prometheus.DefaultRegisterer)

	r := router.Group("/v1")

	v1(r)
}

func v1(router *gin.RouterGroup) {
	router.GET("/batch_status", controller.FinalizeBatchCtl.BatchStatus)
}
