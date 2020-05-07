package auth

import (
	"net/http"
	
	"github.com/gin-gonic/gin"

	"github.com/fajarsep12dev/go-api/api/utils/app"
	C "github.com/fajarsep12dev/go-api/api/utils/constant"
	L "github.com/fajarsep12dev/go-api/api/utils/logger"
)

// AuthController struct
type AuthService struct {
	AuthRepository AuthRepository
}

// Ping handler example
// @Summary Get multiple article tags
// @Produce json
// @Success 200 {string} app.Response
// @Failure 500 {string} app.Response
// @Router /auth/ping [get]
// @tags Ping
func Ping(c *gin.Context) {
	appG := app.Gin{C: c}
	L.Debug("TEST DEBUGGING")
	appG.Response(http.StatusOK, C.Success, map[string]string{
		"message": "ping",
	})
}