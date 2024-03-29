package jwt

import (
	"os"
	"time"

	"github.com/fajarsep12dev/go-api/core/utils/app"

	"github.com/dgrijalva/jwt-go"
)

var jwtSecret = []byte(os.Getenv("JWT_SECRET_KEY"))

// Claims struct
type Claims struct {
	Username string `json:"username"`
	Password string `json:"password"`
	jwt.StandardClaims
}


// GenerateToken generate tokens used for auth
func GenerateToken(username, password string) (string, error) {
	now := time.Now()
	expireTime := now.Add(3 * time.Hour)

	claims := Claims{
		app.EncodeMD5(username),
		app.EncodeMD5(password),
		jwt.StandardClaims{
			ExpiresAt: 	expireTime.Unix(),
			Issuer:    	os.Getenv("JWT_ISSUER"),
			Audience:	os.Getenv("JWT_AUDIENCE"),
		},
	}

	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tokenClaims.SignedString(jwtSecret)

	return token, err
}

// ParseToken parsing token
func ParseToken(token string) (*Claims, error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})

	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
			return claims, nil
		}
	}

	return nil, err
}