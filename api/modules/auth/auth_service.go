package auth

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/fajarsep12dev/go-api/api/modules/auth/dto"
	"github.com/fajarsep12dev/go-api/core/utils/app"
	C "github.com/fajarsep12dev/go-api/core/utils/constant"
	L "github.com/fajarsep12dev/go-api/core/utils/logger"
)

type AuthService struct {
	AuthRepository AuthRepository
}

func ProvideAuthService(authRepo AuthRepository) AuthService {
	return AuthService{
		AuthRepository: authRepo,
	}
} 

// @Summary Get multiple article tags
// @Produce json
// @Success 200 {string} app.Response
// @Failure 500 {string} app.Response
// @Router /auth/ping [get]
// @tags Ping
func Ping(c *gin.Context) {
	appG := app.Gin{C: c}
	L.Debug("TEST DEBUGGING")
	appG.Response(http.StatusOK, C.SUCCESS, map[string]string{
		"message": "ping",
	})
}

// @Summary Get all users
// @Produce json
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
// @Router /auth/getusers [get]
// @tags Users
func (s *AuthService) GetUsers(c *gin.Context) {
	appG := app.Gin{C: c}
	users := s.AuthRepository.GetAll()

	appG.Response(http.StatusOK, C.SUCCESS, ToUserDTOs(users))
}

// @Summary Save data users
// @Produce json
// @Param users body dto.UserDto true "User"
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
// @Router /auth/createusers [post]
// @tags Users
func (s *AuthService) CreateUser(c *gin.Context) {
	var (
		appG = app.Gin{C: c}
		form dto.UserDto
	)

	httpCode, errCode := app.BindAndValid(c, &form)

	if errCode != C.SUCCESS {
		appG.Response(httpCode, errCode, nil)
		return
	}

	createdAuth := s.AuthRepository.Save(ToUser(form))
	appG.Response(http.StatusOK, C.SUCCESS, ToUserDTO(createdAuth))

}