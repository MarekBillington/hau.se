package service

import (
	"hause/entity"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Get User by Id
// No password provided on response
// move to some admin controller
func (h Handler) GetUser(ctx *gin.Context) {
	id := ctx.Param("id")

	var user entity.User

	h.DB.Find(&user, id)

	user.Password = ""

	ctx.JSON(http.StatusOK, user)
}
