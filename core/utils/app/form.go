package app

import (
	"net/http"

	"github.com/astaxie/beego/validation"
	C "github.com/fajarsep12dev/go-api/core/utils/constant"
	"github.com/gin-gonic/gin"
)

// BindAndValid binds and validates data
func BindAndValid(c *gin.Context, form interface{}) (int, int) {
	err := c.BindJSON(form)
	if err != nil {
		return http.StatusBadRequest, C.SUCCESS
	}

	valid := validation.Validation{}
	check, err := valid.Valid(form)
	if err != nil {
		return http.StatusInternalServerError, C.ERROR
	}
	if !check {
		MarkErrors(valid.Errors)
		
		return http.StatusBadRequest, C.INVALID_PARAM
	}

	return http.StatusOK, C.SUCCESS
}
