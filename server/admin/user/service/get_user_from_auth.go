package service

import (
	"hause/auth/utility"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Get user when auth is available
// For userSession setup on client only
func (h Handler) GetUserFromAuth(ctx *gin.Context) {

	user, err := utility.GetRequestingUser(ctx, h.DB)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	user.Password = ""

	ctx.JSON(http.StatusOK, user)
}
