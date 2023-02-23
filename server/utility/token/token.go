package token

import (
	"fmt"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

func GenerateToken(signature string, d int64) (string, error) {

	claims := jwt.MapClaims{}
	claims["authorised"] = true
	claims["signature"] = signature
	claims["exp"] = time.Now().Add(time.Minute * time.Duration(d)).Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// @todo env replace
	return token.SignedString([]byte("shhh it's a secret"))
}

// expected that token will be in request cookie
func ExtractBearerToken(c *gin.Context) (string, error) {

	bearerToken := c.Request.Header.Get("Authorization")
	if len(strings.Split(bearerToken, " ")) == 2 {
		return strings.Split(bearerToken, " ")[1], nil
	}

	return "", fmt.Errorf("%v", "Could not find JWT token in header")
}

func TokenValid(token string) (string, error) {

	t, err := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", t.Header["alg"])
		}
		// @todo replace with env variable
		return []byte("shhh it's a secret"), nil
	})

	if err != nil {
		return "", err
	}

	claims, ok := t.Claims.(jwt.MapClaims)
	if ok && t.Valid {
		sig := fmt.Sprintf("%v", claims["signature"])
		return sig, nil
	}
	return "", nil
}
