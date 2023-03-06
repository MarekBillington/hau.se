package portfolio

import (
	"hause/admin/portfolio/service"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Setup(rg *gin.RouterGroup, conn *gorm.DB) {
	h := service.Handler{
		DB: conn,
	}

	portfolio := rg.Group("/portfolio")

	portfolio.POST("/create", h.CreatePortfolio)

	portfolio.GET("", h.Setup, h.GetRelatedPortfolio)
}
