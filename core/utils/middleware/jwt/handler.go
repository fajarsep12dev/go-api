package jwt

import (
	"fmt"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"

	"github.com/fajarsep12dev/go-api/core/utils/app"
	C "github.com/fajarsep12dev/go-api/core/utils/constant"
)

// JWT is jwt handler gin middlerware
func JWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		var code int
		var data interface{}

		code = C.SUCCESS
		tokenString := c.Request.Header.Get("Authorization")
		fmt.Println(tokenString)

		if tokenString == "" {
			code = C.INVALID_PARAM
		} else {
			//nolint:gosimple
			token := tokenString[7:len(tokenString)]
			fmt.Println(token)
			_, err := ParseToken(token)
			if err != nil {
				switch err.(*jwt.ValidationError).Errors {
				case jwt.ValidationErrorExpired:
					code = C.ERROR_AUTH_CHECK_TOKEN_TIMEOUT
				default:
					code = C.ERROR_AUTH_CHECK_TOKEN_FAIL
				}
			}
		}

		if code != C.SUCCESS {
			c.JSON(http.StatusUnauthorized, gin.H{
				"code": code,
				"msg":  app.GetMsg(code),
				"data": data,
			})

			c.Abort()
			return
		}

		c.Next()
	}
}