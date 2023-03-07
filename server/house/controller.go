package house

import (
	"hause/house/service"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Setup(rg *gin.RouterGroup, conn *gorm.DB) {

	h := service.Handler{
		DB: conn,
	}

	house := rg.Group("/house")

	// create new house
	house.POST("", h.Setup, h.CreateHouse)

	// get all houses
	house.GET("", h.Setup, h.GetAllHouses)

	// get house by id
	house.GET(":id", h.Setup, h.GetHouseById)

	// update house
	house.PATCH(":id", h.UpdateHouse)

	// toggle active flag of house
	house.PATCH(":id/active", h.ToggleHouseActive)
}
