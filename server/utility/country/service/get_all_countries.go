package service

import (
	"hause/entity"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h Handler) GetAllCountries(ctx *gin.Context) {
	var countries []entity.Country

	h.DB.Order("name asc").Find(&countries)

	ctx.JSON(http.StatusOK, countries)
}
