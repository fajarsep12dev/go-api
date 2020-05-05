package api

import (
	"github.com/gin-gonic/gin"
	"github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"

	"github.com/fajarsep12dev/go-api/api/modules/auth"
	_ "github.com/fajarsep12dev/go-api/docs" // Generated docs
	log "github.com/fajarsep12dev/go-api/api/utils/logger"

)

// InitRouter initialize routing information
func InitRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Recovery())

	// Logger middleware
	log.LoggerFileSetup(gin.IsDebugging(), r)

	auth.Routes(r)

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	return r
}


