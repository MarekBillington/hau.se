package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func addHouseRoutes(rg *gin.RouterGroup) {

	house := rg.Group("/house")

	// get all houses
	house.GET("", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, "houses")
	})

	// get house by id
	house.GET(":id", func(ctx *gin.Context) {
		id := ctx.Param("id")
		ctx.JSON(http.StatusOK, "house"+id)
	})
}
