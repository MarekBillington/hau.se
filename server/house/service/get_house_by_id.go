package service

import (
	"hause/entity"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) GetHouseById(ctx *gin.Context) {
	id := ctx.Param("id")
	var house entity.House

	h.DB.Preload("Address").
		First(&house, id).
		Joins("JOIN portfolios ON portfolios.id = houses.portfolio_id").
		Joins("JOIN user_links ON user_links.portfolio_id = portfolios.id AND ul.user_id = ?", h.user.ID)

	if house.ID == 0 {
		ctx.JSON(http.StatusOK, gin.H{})
	}

	ctx.JSON(http.StatusOK, house)
}
