package route

import (
	"net/http"
	"time"

	// enable the pprof
	_ "net/http/pprof"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus"
	"gorm.io/gorm"

	"chain-monitor/common/observability"
	"chain-monitor/internal/controller"
)

// APIHandler routes the APIs
func APIHandler(db *gorm.DB) http.Handler {
	router := gin.New()
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	// Add api metrics.
	observability.Use(router, "chain_monitor", prometheus.DefaultRegisterer)

	// Add health and ready api for k8s.
	probes := observability.NewProbesController(db)
	router.GET("/health", probes.HealthCheck)
	router.GET("/ready", probes.Ready)

	// use v1 version
	v1(router.Group("/v1"), db)
	return router
}

func v1(router *gin.RouterGroup, db *gorm.DB) {
	monitorCtrler := controller.NewMetricsController(db)
	router.GET("/blocks_status", monitorCtrler.ConfirmBlocksStatus)
	router.GET("/batch_status", monitorCtrler.ConfirmBatchStatus)
}
