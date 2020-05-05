package auth

import (
	"github.com/gin-gonic/gin"
)

// Routes auth group
func Routes(route *gin.Engine) {
	r := route.Group("/api/auth")
		r.GET("/ping", Ping)
}