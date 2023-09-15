package route

import (
	"net/http"
	"time"

	// enable the pprof
	_ "net/http/pprof"

	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/urfave/cli/v2"
	"gorm.io/gorm"

	"chain-monitor/common/observability"
	"chain-monitor/internal/controller"
	"chain-monitor/internal/utils"
)

// APIHandler routes the APIs
func APIHandler(ctx *cli.Context, db *gorm.DB, reg prometheus.Registerer) http.Handler {
	router := gin.New()
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	// Open metrics.
	if ctx.Bool(utils.MetricsEnabled.Name) {
		pprof.Register(router)
		// Add api metrics.
		observability.Use(router, "chain_monitor", reg)
		// Add metrics api.
		router.GET("/metrics", func(context *gin.Context) {
			promhttp.Handler().ServeHTTP(context.Writer, context.Request)
		})
		// Add health and ready api for k8s.
		probes := observability.NewProbesController(db)
		router.GET("/health", probes.HealthCheck)
		router.GET("/ready", probes.Ready)
	}

	// use v1 version
	v1(router.Group("/v1"), db)
	return router
}

func v1(router *gin.RouterGroup, db *gorm.DB) {
	monitorCtrler := controller.NewMetricsController(db)
	router.GET("/blocks_status", monitorCtrler.ConfirmBlocksStatus)
	router.GET("/batch_status", monitorCtrler.ConfirmBatchStatus)
}
