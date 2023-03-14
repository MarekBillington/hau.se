package service

import (
	"hause/house/dto"
	"hause/house/helper"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) CreateHouse(ctx *gin.Context) {
	var new dto.NewHouse

	if err := ctx.ShouldBindJSON(&new); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	house := helper.MapHouseFromNewDto(new)

	if new.PropertyID != 0 {
		house.PropertyID = new.PropertyID
	}

	h.DB.Create(&house)

	ctx.JSON(http.StatusOK, house)
}
