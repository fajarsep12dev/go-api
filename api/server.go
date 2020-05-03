package api

import (
	"fmt"
	"net/http"
	"os"
	"time"
	"github.com/rs/zerolog/log"
	"github.com/gin-gonic/gin"
	"github.com/fajarsep12dev/go-api/api/utils/setting"
)


func init() {
	setting.Setup()
}

// Run web server
func Run() {
	gin.SetMode(os.Getenv("RUN_MODE"))

	router := InitRouter()
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
