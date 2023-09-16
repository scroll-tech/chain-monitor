package route

import (
	"net/http"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

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

	// use v1 version
	v1(router.Group("/v1"), db)
	return router
}

func v1(router *gin.RouterGroup, db *gorm.DB) {
	monitorCtrler := controller.NewMetricsController(db)
	router.GET("/blocks_status", monitorCtrler.ConfirmBlocksStatus)
	router.GET("/batch_status", monitorCtrler.ConfirmBatchStatus)
}
