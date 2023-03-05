package service

import (
	"hause/house/dto"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h Handler) CreateHouse(ctx *gin.Context) {
	var house dto.NewHouse

	if err := ctx.ShouldBindJSON(&house); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	h.DB.Create(&house)
	ctx.JSON(http.StatusOK, house)
}
