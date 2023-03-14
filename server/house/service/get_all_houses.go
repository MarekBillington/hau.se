package service

import (
	"hause/entity"
	"hause/utility/database/helper"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) GetAllHouses(ctx *gin.Context) {
	var houses []entity.House

	query, exists := ctx.GetQuery("active")

	tx := helper.JoinPortfolio(h.DB, houses, h.user.ID, "").
		Order("id asc")

	if exists {
		tx.Where("houses.active = ?", query)
	}

	tx.Find(&houses)

	ctx.JSON(http.StatusOK, houses)
}
