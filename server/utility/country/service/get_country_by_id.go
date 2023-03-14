package service

import (
	"hause/entity"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h Handler) GetCountryById(ctx *gin.Context) {
	id := ctx.Param("id")
	var country entity.Country

	h.DB.Find(&country, id)

	if country.ID == 0 {
		ctx.JSON(http.StatusOK, gin.H{})
		return
	}

	ctx.JSON(http.StatusOK, country)
}
