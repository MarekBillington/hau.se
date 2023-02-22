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
	house.GET("", getHouses)

	// get house by id
	house.GET(":id", getHouse)

	// create new house
	house.POST("", postHouse)

	// update house
	house.PATCH(":id", patchHouse)

	// enable house
	house.PATCH(":id/enable", toggleHouseActive)

	// disable house
	house.PATCH(":id/disable", toggleHouseActive)
}

func getHouses(ctx *gin.Context) {
	var houses []House

	query, _ := ctx.GetQuery("active")

	db.Order("id asc").Find(&houses, "active = ?", query)

	ctx.JSON(http.StatusOK, houses)
}

func getHouse(ctx *gin.Context) {
	id := ctx.Param("id")
	var house House

	db.First(&house, id)

	if house.ID == 0 {
		ctx.JSON(http.StatusOK, gin.H{})
	}

	ctx.JSON(http.StatusOK, house)
}

func postHouse(ctx *gin.Context) {
	var house House

	if err := ctx.ShouldBindJSON(&house); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	db.Create(&house)
	ctx.JSON(http.StatusOK, house)
}

func patchHouse(ctx *gin.Context) {
	id := ctx.Param("id")
	var newHouse House
	var house House

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

	var house House
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
