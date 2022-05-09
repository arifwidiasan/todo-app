package helper

import (
	"time"

	"github.com/golang-jwt/jwt"
)

func CreateToken(id int, username, secret string) (string, error) {
	claims := jwt.MapClaims{}
	claims["exp"] = time.Now().Add((24 * time.Hour)).Unix() // kapan token gabisa dipake
	// claims["iat"] = time.Now().Unix()                         // kapan token dibuat
	// claims["nbf"] = time.Now().Add((25 * time.Second)).Unix() // kapan token bisa digunakan == iat
	claims["username"] = username
	claims["id"] = id

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(secret))
}
