package helper

import (
	"hause/entity"
	"hause/house/dto"
)

func MapHouseFromNewDto(newHouse dto.NewHouse) entity.House {

	return entity.House{
		PortfolioID: newHouse.PortfolioID,
		Street1:     newHouse.Street1,
		Street2:     newHouse.Street2,
		Postcode:    newHouse.Postcode,
		Town:        newHouse.Town,
		CountryID:   newHouse.CountryID,
		Bedrooms:    newHouse.Bedrooms,
		Bathrooms:   newHouse.Bathrooms,
		Garage:      newHouse.Garage,
		Floorspace:  newHouse.Floorspace,
		Landarea:    newHouse.Landarea,
	}
}

func MapHouseFromUpdateDto(uHouse dto.UpdateHouse) entity.House {

	return entity.House{
		PortfolioID: uHouse.PortfolioID,
		Street1:     uHouse.Street1,
		Street2:     uHouse.Street2,
		Postcode:    uHouse.Postcode,
		Town:        uHouse.Town,
		CountryID:   uHouse.CountryID,
		Bedrooms:    uHouse.Bedrooms,
		Bathrooms:   uHouse.Bathrooms,
		Garage:      uHouse.Garage,
		Floorspace:  uHouse.Floorspace,
		Landarea:    uHouse.Landarea,
	}
}
