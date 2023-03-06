package service

import (
	"hause/entity"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) GetAllHouses(ctx *gin.Context) {
	var houses []entity.House

	query, exists := ctx.GetQuery("active")

	tx := h.DB.Preload("Address").
		Joins("JOIN portfolios ON portfolios.id = houses.portfolio_id").
		Joins("JOIN user_links ON user_links.portfolio_id = portfolios.id AND user_links.user_id = ?", h.user.ID).
		Order("id asc")

	if exists {
		tx.Where("houses.active = ?", query)
	}

	tx.Find(&houses)

	ctx.JSON(http.StatusOK, houses)
}
