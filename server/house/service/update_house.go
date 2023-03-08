package service

import (
	"hause/entity"
	"hause/house/dto"
	"hause/utility/database/helper"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h Handler) UpdateHouse(ctx *gin.Context) {
	id := ctx.Param("id")
	var uh dto.UpdateHouse
	var house entity.House

	helper.JoinPortfolio(h.DB, house, h.user.ID, "Address").
		First(&house, id)

	if house.ID == 0 {
		// @todo move to validation?
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Could not find object by id"})
		return
	}

	if err := ctx.ShouldBindJSON(&uh); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	h.DB.Model(&house).Updates(&uh)

	ctx.JSON(http.StatusOK, house)
}