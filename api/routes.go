package api

import (
	"github.com/gin-gonic/gin"
	"github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
	"github.com/jinzhu/gorm"

	"github.com/fajarsep12dev/go-api/api/modules/auth"
	_ "github.com/fajarsep12dev/go-api/docs" // Generated docs
	log "github.com/fajarsep12dev/go-api/api/utils/logger"

)

// InitRouter initialize routing information
func InitRouter(db *gorm.DB) *gin.Engine {
	r := gin.New()
	r.Use(gin.Recovery())

	// Logger middleware
	log.LoggerFileSetup(gin.IsDebugging(), r)

	authService := InitAuthService(db)

	authApi := r.Group("/api/auth")
	authApi.GET("/ping", auth.Ping)
	authApi.GET("/getusers", authService.GetUsers)

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	return r
}


