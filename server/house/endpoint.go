package house

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func addRoutes(rg *gin.RouterGroup) {

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

	// enable house
	house.PATCH(":id/enable", func(ctx *gin.Context) {
		toggleHouseActive(ctx)
	})

	// disable house
	house.PATCH(":id/disable", func(ctx *gin.Context) {
		toggleHouseActive(ctx)
	})
}

func getHouses(ctx *gin.Context) {
	var houses []house

	query, _ := ctx.GetQuery("active")

	db.Order("id asc").Find(&houses, "active = ?", query)

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

// Using gorm, we are required to be deliberate when a field is made empty
// therefore easiest option is to have specific enable/disable endpoints
// @link https://gorm.io/docs/update.html#Updates-multiple-columns
func toggleHouseActive(ctx *gin.Context) {
	id := ctx.Param("id")

	var house house
	fmt.Fprintf(os.Stdout, "%+v", id)
	db.First(&house, id)
	if house.ID == 0 {
		// @todo move to validation?
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Could not find object by id"})
		return
	}
	house.Active = !house.Active
	db.Model(&house).Update("active", house.Active)

	ctx.JSON(http.StatusOK, house)
}
