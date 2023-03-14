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
	var input dto.UpdateHouse
	var house entity.House

	helper.JoinPortfolio(h.DB, house, h.user.ID, "Address").
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

	update := entity.House{
		PortfolioID: input.PortfolioID,
		Street1:     input.Street1,
		Street2:     input.Street2,
		Postcode:    input.Postcode,
		Town:        input.Town,
		CountryID:   input.CountryID,
		Bedrooms:    input.Bedrooms,
		Bathrooms:   input.Bathrooms,
		Garage:      input.Garage,
		Floorspace:  input.Floorspace,
		Landarea:    input.Landarea,
	}

	h.DB.Model(&house).Updates(&update)

	ctx.JSON(http.StatusOK, house)
}
