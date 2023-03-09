package service

import (
	"hause/entity"
	"hause/house/dto"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) CreateHouse(ctx *gin.Context) {
	var new dto.NewHouse

	if err := ctx.ShouldBindJSON(&new); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var house = entity.House{
		PortfolioID: new.PortfolioID,
		Bedrooms:    new.Bedrooms,
		Bathrooms:   new.Bathrooms,
		Garage:      new.Garage,
		Floorspace:  new.Floorspace,
		Landarea:    new.Landarea,
	}

	if new.PropertyID != 0 {
		house.PropertyID = new.PropertyID
	}

	h.DB.Create(&house)

	// @todo move to some mapper???
	var address = entity.Address{
		Street1:   new.NewAddress.Street1,
		Street2:   new.NewAddress.Street2,
		Postcode:  new.NewAddress.Postcode,
		Town:      new.NewAddress.Town,
		CountryID: new.NewAddress.CountryID,
		HouseID:   house.ID,
	}

	h.DB.Create(&address)
	ctx.JSON(http.StatusOK, house)
}
