package service

import (
	"fmt"
	"hause/entity"
	"hause/house/dto"
	mapper "hause/house/helper"
	"hause/utility/database/helper"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func (h *Handler) UpdateHouse(ctx *gin.Context) {
	id := ctx.Param("id")
	var input dto.UpdateHouse
	var house entity.House

	fmt.Fprintf(os.Stdout, "%+v", h.user)
	helper.JoinPortfolio(h.DB, house, h.user.ID, "").
		First(&house, id)

	if house.ID == 0 {
		// @todo move to validation?
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Could not find object by id"})
		return
	}

	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	update := mapper.MapHouseFromUpdateDto(input)

	h.DB.Model(&house).Updates(&update)

	ctx.JSON(http.StatusOK, house)
}
