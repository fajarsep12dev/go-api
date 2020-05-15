//+build wireinject

package api

import (
	"github.com/fajarsep12dev/go-api/api/auth"
	"github.com/google/wire"
	"github.com/jinzhu/gorm"
)

func initAuthService(db *gorm.DB) auth.AuthService {
	wire.Build(auth.ProvideAuthRepository, auth.ProvideAuthService)

	return auth.AuthService{}
}