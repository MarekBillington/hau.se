package service

import (
	"hause/admin/user/helper"
	"hause/entity"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Get all users where suitable
// If the user is not the admin then only their user will be returned
// If the admin they can get all users associated to the portfolio
func (h *Handler) GetAllUsers(ctx *gin.Context) {

	var users []entity.User

	tx := h.DB.Where("users.id = ?", h.user.ID)

	if h.user.UserRole.Name == "Admin" {
		tx = helper.JoinPortfolio(tx)
	}

	tx.Find(&users)

	ctx.JSON(http.StatusOK, users)
}
