package app

import (
	"net/http"
	c "github.com/fajarsep12dev/api/go-api/utils/constant"
	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
)

// BindAndValid binds and validates data
func BindAndValid(c *gin.Context, form interface{}) (int, int) {
	err := c.Bind(form)
	if err != nil {
		return http.StatusBadRequest, c.Success
	}

	valid := validation.Validation{}
	check, err := valid.Valid(form)
	if err != nil {
		return http.StatusInternalServerError, c.Error
	}
	if !check {
		MarkErrors(valid.Errors)
		return http.StatusBadRequest, c.InvalidParam
	}

	return http.StatusOK, c.Success
}
