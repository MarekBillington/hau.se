package service

import (
	"hause/auth/dto"
	"hause/auth/repository"
	"hause/auth/utility"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Register creates the user with a provided email and password
// Returns JWT for use as Bearer token
func (h Handler) RegisterUser(ctx *gin.Context) {
	var reg dto.Register

	if err := ctx.ShouldBindJSON(&reg); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := repository.CreateUser(reg, h.DB)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	t, exp, err := utility.ServeJWT(user.Email, ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = utility.ServeRefCookie(user.Email, ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// set to header the refresh token cookie

	ctx.JSON(http.StatusOK, gin.H{"token": t, "expires": exp})
}
