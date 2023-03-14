package country

import (
	"hause/utility/country/service"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Setup(rg *gin.RouterGroup, conn *gorm.DB) {

	h := service.Handler{
		DB: conn,
	}

	uc := rg.Group("/utility/country")

	// Get all countries
	uc.GET("", h.GetAllCountries)

	// Get country by Id
	uc.GET(":id", h.GetCountryById)
}
