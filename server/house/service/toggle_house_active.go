package service

import (
	"hause/entity"
	"hause/utility/database/helper"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Using gorm, we are required to be deliberate when a field is made empty
// therefore easiest option is to have specific enable/disable endpoints
// @link https://gorm.io/docs/update.html#Updates-multiple-columns
func (h Handler) ToggleHouseActive(ctx *gin.Context) {
	id := ctx.Param("id")

	var house entity.House

	helper.JoinPortfolio(h.DB, house, h.user.ID, "Address").
		First(&house, id)

	if house.ID == 0 {
		// @todo move to validation?
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Could not find object by id"})
		return
	}

	house.Active = !house.Active
	h.DB.Model(&house).Update("active", house.Active)

	ctx.JSON(http.StatusOK, house)
}
