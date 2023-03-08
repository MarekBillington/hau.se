package service

import (
	"hause/admin/user/dto"
	"hause/auth/utility"
	"hause/entity"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Get user when auth is available
// For userSession setup on client only
func (h Handler) GetUserForSetup(ctx *gin.Context) {

	user, err := utility.GetRequestingUser(ctx, h.DB)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	user.Password = ""

	var portfolio entity.Portfolio
	h.DB.
		Joins("JOIN user_links ON portfolio.id = user_links.portfolio_id AND user_links.user_id = ?", user.ID).
		First(&portfolio)

	res := dto.SetupResponse{
		User:      user,
		Portfolio: portfolio,
	}

	ctx.JSON(http.StatusOK, res)
}
