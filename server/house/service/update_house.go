package service

import (
	"hause/entity"
	"hause/house/dto"
	mapper "hause/house/helper"
	"hause/utility/database/helper"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) UpdateHouse(ctx *gin.Context) {
	id := ctx.Param("id")
	var input dto.UpdateHouse
	var house entity.House

	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	helper.JoinPortfolio(h.DB, house, h.user.ID, "").
		First(&house, id)

	if house.ID == 0 {
		// @todo move to validation?
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Could not find object by id"})
		return
	}

	update := mapper.MapHouseFromUpdateDto(input)

	h.DB.Model(&house).Updates(&update)

	ctx.JSON(http.StatusOK, house)
}
