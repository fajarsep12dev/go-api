package jwt

import (
	"fmt"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"

	"github.com/fajarsep12dev/go-api/api/utils/app"
	C "github.com/fajarsep12dev/go-api/api/utils/constant"
)

// JWT is jwt handler gin middlerware
func JWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		var code int
		var data interface{}

		code = C.Success
		tokenString := c.Request.Header.Get("Authorization")
		fmt.Println(tokenString)

		if tokenString == "" {
			code = C.InvalidParam
		} else {
			//nolint:gosimple
			token := tokenString[7:len(tokenString)]
			fmt.Println(token)
			_, err := ParseToken(token)
			if err != nil {
				switch err.(*jwt.ValidationError).Errors {
				case jwt.ValidationErrorExpired:
					code = C.ErrorAuthCheckTokenTimeOut
				default:
					code = C.ErrorAuthCheckTokenFail
				}
			}
		}

		if code != C.Success {
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