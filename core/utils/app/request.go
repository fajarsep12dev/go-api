package app

import (
	"github.com/astaxie/beego/validation"
	log "github.com/fajarsep12dev/go-api/core/utils/logger"
)

// MarkErrors logs error logs
func MarkErrors(errors []*validation.Error) {
	for _, err := range errors {
		// log.Info().
		// 	Str("Error", err.Key).
		// 	Msg(err.Message)
		log.Error("Error : %s %s", err.Key, err.Message)
	}
}
