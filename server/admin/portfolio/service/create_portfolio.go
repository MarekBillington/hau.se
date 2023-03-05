package service

import (
	"hause/admin/portfolio/dto"
	"hause/entity"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h Handler) CreatePortfolio(ctx *gin.Context) {
	var np dto.NewPortfolio

	if err := ctx.ShouldBindJSON(&np); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	portfolio := entity.Portfolio{
		Name:      np.Name,
		CountryID: np.CountryID,
	}

	h.DB.Create(&portfolio)

	link := entity.UserLink{
		PortfolioID: portfolio.ID,
		UserID:      np.UserID,
	}

	h.DB.Create(&link)

	ctx.JSON(http.StatusOK, portfolio)
}
