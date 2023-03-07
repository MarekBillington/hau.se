package service

import (
	"hause/entity"
	"hause/utility/database/helper"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) GetHouseById(ctx *gin.Context) {
	id := ctx.Param("id")
	var house entity.House

	helper.JoinPortfolio(h.DB, house, h.user.ID, "Address").
		First(&house, id)

	if house.ID == 0 {
		ctx.JSON(http.StatusOK, gin.H{})
	}

	ctx.JSON(http.StatusOK, house)
}
