package utility

import (
	"time"

	"github.com/gin-gonic/gin"
)

// Creates the auth token used as bearer for rest of api
func ServeJWT(signature string, ctx *gin.Context) (string, int64, error) {
	exp := time.Now().Add(time.Minute * time.Duration(10)).Unix()

	// on correct create new token which is served as output
	t, err := GenerateToken(signature, exp)
	if err != nil {
		return "", 0, err
	}

	return t, exp, nil
}
