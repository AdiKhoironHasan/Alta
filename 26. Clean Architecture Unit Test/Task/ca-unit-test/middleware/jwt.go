package middleware

import (
	"belajar-go-echo/config"
	conf "belajar-go-echo/config"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func JWTMiddleware() echo.MiddlewareFunc {
	return middleware.JWTWithConfig(middleware.JWTConfig{
		SigningMethod: "HS256",
		SigningKey:    conf.SECRET_JWT,
	})
}

func CreateTokenJWT() (string, error) {
	claims := jwt.MapClaims{}
	claims["exp"] = time.Now().Add(time.Hour * 1)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	jwtToken, err := token.SignedString([]byte(config.SECRET_JWT))

	if err != nil {
		return "", err
	}

	return jwtToken, nil
}
