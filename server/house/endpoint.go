package house

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func addRoutes(rg *gin.RouterGroup) {

	// @todo should pass
	house := rg.Group("/house")

	// get all houses
	house.GET("", func(ctx *gin.Context) {
		getHouses(ctx)
	})

	// get house by id
	house.GET(":id", func(ctx *gin.Context) {
		getHouse(ctx)
	})

	// create new house
	house.POST("", func(ctx *gin.Context) {
		postHouse(ctx)
	})

	// update house
	house.PATCH(":id", func(ctx *gin.Context) {
		patchHouse(ctx)
	})
}

func getHouses(ctx *gin.Context) {
	var houses []house

	db.Find(&houses)

	ctx.JSON(http.StatusOK, houses)
}

func getHouse(ctx *gin.Context) {
	id := ctx.Param("id")
	var house house

	db.First(&house, id)

	ctx.JSON(http.StatusOK, house)
}

func postHouse(ctx *gin.Context) {
	var house house

	if err := ctx.ShouldBindJSON(&house); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	db.Create(&house)
	ctx.JSON(http.StatusOK, house)
}

func patchHouse(ctx *gin.Context) {
	id := ctx.Param("id")
	var newHouse house
	var house house

	db.First(&house, id)
	if house.ID == 0 {
		// @todo move to validation?
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Could not find object by id"})
		return
	}

	if err := ctx.ShouldBindJSON(&newHouse); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	db.Model(&house).Updates(&newHouse)

	ctx.JSON(http.StatusOK, house)
}
