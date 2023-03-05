package house

import (
	"fmt"
	"hause/entity"
	"hause/house/service"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var h service.Handler

func Setup(rg *gin.RouterGroup, conn *gorm.DB) {

	h := service.Handler{
		DB: conn,
	}

	house := rg.Group("/house")
	
	// create new house
	house.POST("", h.CreateHouse)

	// get all houses
	house.GET("", getHouses)

	// get house by id
	house.GET(":id", getHouse)

	// update house
	house.PATCH(":id", patchHouse)

	// enable house
	house.PATCH(":id/enable", toggleHouseActive)

	// disable house
	house.PATCH(":id/disable", toggleHouseActive)
}

func getHouses(ctx *gin.Context) {
	var houses []entity.House

	query, _ := ctx.GetQuery("active")

	h.DB.Order("id asc").Find(&houses, "active = ?", query)

	ctx.JSON(http.StatusOK, houses)
}

func getHouse(ctx *gin.Context) {
	id := ctx.Param("id")
	var house entity.House

	h.DB.First(&house, id)

	if house.ID == 0 {
		ctx.JSON(http.StatusOK, gin.H{})
	}

	ctx.JSON(http.StatusOK, house)
}

func patchHouse(ctx *gin.Context) {
	id := ctx.Param("id")
	var newHouse entity.House
	var house entity.House

	h.DB.First(&house, id)
	if house.ID == 0 {
		// @todo move to validation?
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Could not find object by id"})
		return
	}

	if err := ctx.ShouldBindJSON(&newHouse); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	h.DB.Model(&house).Updates(&newHouse)

	ctx.JSON(http.StatusOK, house)
}

// Using gorm, we are required to be deliberate when a field is made empty
// therefore easiest option is to have specific enable/disable endpoints
// @link https://gorm.io/docs/update.html#Updates-multiple-columns
func toggleHouseActive(ctx *gin.Context) {
	id := ctx.Param("id")

	var house entity.House
	fmt.Fprintf(os.Stdout, "%+v", id)
	h.DB.First(&house, id)
	if house.ID == 0 {
		// @todo move to validation?
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Could not find object by id"})
		return
	}
	house.Active = !house.Active
	h.DB.Model(&house).Update("active", house.Active)

	ctx.JSON(http.StatusOK, house)
}
