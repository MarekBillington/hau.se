package service

import (
	"fmt"
	"hause/auth/repository"
	"hause/auth/utility"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

// Refresh access token with provided refresh token provided on first request
// Returns with new JWT and Refresh tokens
func (h Handler) RefreshUser(ctx *gin.Context) {

	// get cookie from request, check it is valid
	refToken, err := ctx.Cookie("refresh")
	if err != nil {
		ctx.String(http.StatusUnauthorized, "Unauthorized")
		return
	}

	userEmail, err := validateToken(refToken)
	if err != nil {
		ctx.String(http.StatusUnauthorized, "Unauthorized")
		return
	}

	user := repository.GetUserFromEmail(userEmail, h.DB)

	t, exp, err := utility.ServeJWT(user.Email, ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"token": t, "expiry": exp})
}

// If the token can be decrypted, the signature will be returned
func validateToken(token string) (string, error) {

	t, err := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
		}
		// @todo replace with env variable
		return []byte(utility.Secret), nil
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
