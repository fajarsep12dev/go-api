package api

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"

	"github.com/fajarsep12dev/go-api/core/db"
	"github.com/fajarsep12dev/go-api/core/db/migrations"
	"github.com/fajarsep12dev/go-api/core/db/seed"
	"github.com/fajarsep12dev/go-api/core/utils/setting"
)


var orm = db.ORM{}

func init() {
	setting.Initialize()
	orm.Initialize()
	migrations.Initialize(orm.DB)
	seed.Initialize(orm.DB)
}

// Run web server
func Run() {
	gin.SetMode(os.Getenv("RUN_MODE"))
	
	router := InitRouter(orm.DB)
	endPoint := fmt.Sprintf(":%s", os.Getenv("HTTP_PORT"))
	maxHeaderBytes := 1 << 20

	// Initialize gin instance
	server := &http.Server{
		Addr:           endPoint,
		Handler:        router,
		ReadTimeout:    60 * time.Second,
		WriteTimeout:   60 * time.Second,
		MaxHeaderBytes: maxHeaderBytes,
	}

	log.Info().Msgf("[info] start http server listening %s", endPoint)

	server.ListenAndServe()
}
