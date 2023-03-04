package utility

import (
	"time"

	"github.com/gin-gonic/gin"
)

// Creates the refresh token and sends to set as cookie in response
// @todo utilise local dto instad of entity ref
func ServeRefCookie(signature string, ctx *gin.Context) error {

	refexp := time.Now().Add(time.Hour * time.Duration(24)).Unix()
	r, err := GenerateToken(signature, refexp)
	if err != nil {
		return err
	}

	// @todo tidy up url and for prod set to secure
	ctx.SetCookie("refresh", r, int(refexp), "/", "dev.hau.se", false, true)
	return nil
}
