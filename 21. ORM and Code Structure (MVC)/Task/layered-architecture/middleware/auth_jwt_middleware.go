package middleware

import (
	con "layered-architecture/constants"
	"time"

	"github.com/dgrijalva/jwt-go"
)

func CreateToken(userId int, name string) (string, error) {
	claims := jwt.MapClaims{
		"userId": userId,
		"name":   name,
		"exp":    time.Now().Add(1 * time.Hour).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte(con.SECRET_JWT))
}
