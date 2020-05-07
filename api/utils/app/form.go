package app

import (
	"net/http"
	C "github.com/fajarsep12dev/go-api/api/utils/constant"
	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
)

// BindAndValid binds and validates data
func BindAndValid(c *gin.Context, form interface{}) (int, int) {
	err := c.BindJSON(form)
	if err != nil {
		return http.StatusBadRequest, C.Success
	}

	valid := validation.Validation{}
	check, err := valid.Valid(form)
	if err != nil {
		return http.StatusInternalServerError, C.Error
	}
	if !check {
		MarkErrors(valid.Errors)
		
		return http.StatusBadRequest, C.InvalidParam
	}

	return http.StatusOK, C.Success
}
