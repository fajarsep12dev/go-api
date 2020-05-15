
// +build ignore

package api

import (
	"github.com/fajarsep12dev/go-api/api/modules/auth"
	"github.com/google/wire"
	"github.com/jinzhu/gorm"
)

func InitAuthService(db *gorm.DB) auth.AuthService {
	wire.Build(auth.ProvideAuthRepository, auth.ProvideAuthService)

	return auth.AuthService{}
}