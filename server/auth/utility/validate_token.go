package utility

import (
	"fmt"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

// @todo move to env
const Secret = "shhh it's a secret"

// If the token can be decrypted, the signature will be returned
func ValidateToken(ctx *gin.Context) (string, error) {

	token, err := extractBearerToken(ctx)
	if err != nil {
		return "", err
	}

	t, err := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
		}
		// @todo replace with env variable
		return []byte(Secret), nil
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

// expected that token will be in request cookie
func extractBearerToken(c *gin.Context) (string, error) {

	bearerToken := c.Request.Header.Get("Authorization")
	if len(strings.Split(bearerToken, " ")) == 2 {
		return strings.Split(bearerToken, " ")[1], nil
	}

	return "", fmt.Errorf("%v", "Could not find JWT token in header")
}
