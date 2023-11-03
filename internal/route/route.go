package route

import (
	"net/http"

	// enable the pprof
	_ "net/http/pprof"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// APIHandler routes the APIs
func APIHandler(db *gorm.DB) http.Handler {
	router := gin.New()
	router.Use(gin.Recovery())

	// use v1 version
	v1(router.Group("/v1"), db)
	return router
}

func v1(router *gin.RouterGroup, db *gorm.DB) {
}
