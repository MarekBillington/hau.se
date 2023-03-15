package service

import (
	"hause/admin/user/helper"
	"hause/entity"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (h *Handler) GetUserById(ctx *gin.Context) {
	id := ctx.Param("id")
	uid, _ := strconv.ParseUint(id, 10, 64)
	u32id := uint(uid)

	var user entity.User

	// If the user is not an admin and they are not getting their own data, we provide nothing
	if h.user.UserRole.Name != "Admin" && h.user.ID != u32id {
		ctx.JSON(http.StatusOK, gin.H{})
		return
	}

	tx := helper.JoinPortfolio(h.DB)

	tx.First(&user, id)

	ctx.JSON(http.StatusOK, user)
}
