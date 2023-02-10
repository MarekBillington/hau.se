package house

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type House struct {
	ID         int    `json:"id" gorm:"primaryKey"`
	Address    string `json:"address"`
	Bedrooms   string `json:"bedrooms"`
	Bathrooms  string `json:"bathrooms"`
	Garage     string `json:"garage"`
	Floorspace string `json:"floorspace"`
	Landarea   string `json:"Landarea"`
}

func addRoutes(rg *gin.RouterGroup) {

	house := rg.Group("/house")

	// get all houses
	house.GET("", func(ctx *gin.Context) {
		houses := getHouses()
		ctx.JSON(http.StatusOK, houses)
	})

	// get house by id
	house.GET(":id", func(ctx *gin.Context) {
		id := ctx.Param("id")
		house := getHouse(id)
		ctx.JSON(http.StatusOK, house)
	})

	// create new house
	// house.POST("", func(ctx *gin.Context) {
	// 	house := postHouse(ctx.Params)
	// 	ctx.JSON(http.StatusOK, house)
	// })
}

func initDb() {
	db.AutoMigrate(&House{})
}

func getHouses() []House {
	var houses []House
	db.Find(&houses)
	return houses
}

func getHouse(id string) House {
	var house House
	db.First(&house, id)
	return house
}

func postHouse(params gin.Params) House {
	// @todo add a data validator
	fmt.Println(params);
	house := House{
		Address:    params.ByName("address"),
		Bedrooms:   params.ByName("bedrooms"),
		Bathrooms:  params.ByName("bathrooms"),
		Garage:     params.ByName("garage"),
		Floorspace: params.ByName("floorspace"),
		Landarea:   params.ByName("landarea"),
	}
	db.Create(&house)
	return house
}
