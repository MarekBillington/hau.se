package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type House struct {
	ID      int    `json:"id" gorm:"primaryKey"`
	Address string `json:"address"`
}

func addHouseRoutes(rg *gin.RouterGroup) {

	house := rg.Group("/house")

	// get all houses
	house.GET("", func(ctx *gin.Context) {
		hs := getHouses()

		ctx.JSON(http.StatusOK, hs)
	})

	// get house by id
	house.GET(":id", func(ctx *gin.Context) {
		id := ctx.Param("id")
		ctx.JSON(http.StatusOK, "house"+id)
	})
}

func getHouses() []House {
	// h1 := new(House)
	// h1.ID = 1
	// h1.Address = "4321 test road"
	var houses []House

	db.Find(&houses)

	return houses
}
