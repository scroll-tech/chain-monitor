package route

import (
	"net/http"

	// enable the pprof
	_ "net/http/pprof"

	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"gorm.io/gorm"

	"chain-monitor/common/observability"
	"chain-monitor/internal/controller"
)

// APIHandler routes the APIs
func APIHandler(db *gorm.DB) http.Handler {
	router := gin.New()
	router.Use(gin.Recovery())

	// Register metrics.
	observability.Use(router, "chain_monitor", prometheus.DefaultRegisterer)

	// use v1 version
	v1(router.Group("/v1"), db)
	return router
}

func v1(router *gin.RouterGroup, db *gorm.DB) {
	monitorCtrler := controller.NewMetricsController(db)
	router.GET("/blocks_status", monitorCtrler.ConfirmBlocksStatus)
	router.GET("/batch_status", monitorCtrler.ConfirmBatchStatus)
}

// MetricsHandler metrics Handler.
func MetricsHandler(db *gorm.DB) http.Handler {
	router := gin.New()
	router.Use(gin.Recovery())
	pprof.Register(router)
	router.GET("/metrics", func(context *gin.Context) {
		promhttp.Handler().ServeHTTP(context.Writer, context.Request)
	})

	// Add health api for k8s.
	probeController := observability.NewProbesController(db)
	router.GET("/health", probeController.HealthCheck)
	router.GET("/ready", probeController.Ready)

	return router
}
