package service

import (
	"hause/auth/dto"
	"hause/auth/utility"
	"hause/entity"
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

	user, err := h.createUser(reg)
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

// private function to act as repo for service
func (h Handler) createUser(reg dto.Register) (entity.User, error) {

	user := entity.User{
		Email:     reg.Email,
		Password:  reg.Password,
		FirstName: reg.FirstName,
		LastName:  reg.LastName,
	}

	err := h.DB.Create(&user).Error
	if err != nil {
		return user, err
	}

	return user, nil
}
