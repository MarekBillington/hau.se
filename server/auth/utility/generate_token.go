package utility

import (
	"time"

	"github.com/golang-jwt/jwt"
)

func GenerateToken(signature string, d int64) (string, error) {

	claims := jwt.MapClaims{}
	claims["authorised"] = true
	claims["signature"] = signature
	claims["exp"] = time.Now().Add(time.Minute * time.Duration(d)).Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// @todo env replace
	return token.SignedString([]byte(Secret))
}
