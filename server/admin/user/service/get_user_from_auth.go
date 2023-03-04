package service

import (
	"hause/admin/user/repository"
	"hause/auth/utility"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Get user when auth is available
func (h Handler) GetUserFromAuth(ctx *gin.Context) {

	// @todo just user token valid and rename
	signature, err := utility.ValidateToken(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	user := repository.GetUserFromEmail(signature, h.DB)

	user.Password = ""

	ctx.JSON(http.StatusOK, user)
}
