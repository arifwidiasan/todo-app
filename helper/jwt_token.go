package helper

import (
	"time"

	"github.com/golang-jwt/jwt"
)

func CreateToken(id int, username, secret string) (string, error) {
	claims := jwt.MapClaims{}
	claims["exp"] = time.Now().Add((24 * time.Hour)).Unix()
	claims["username"] = username
	claims["id"] = id

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(secret))
}
