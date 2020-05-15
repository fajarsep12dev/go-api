package api

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"

	"github.com/fajarsep12dev/go-api/api/auth"
	log "github.com/fajarsep12dev/go-api/core/utils/logger"
	_ "github.com/fajarsep12dev/go-api/docs" // Generated docs
)

// InitRouter initialize routing information
func InitRouter(db *gorm.DB) *gin.Engine {
	r := gin.New()
	r.Use(gin.Recovery())

	// Logger middleware
	log.LoggerFileSetup(gin.IsDebugging(), r)

	authService := initAuthService(db)

	authApi := r.Group("/api/auth")
	authApi.GET("/ping", auth.Ping)
	authApi.GET("/getusers", authService.GetUsers)
	authApi.POST("/createusers", authService.CreateUser)

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	return r
}


