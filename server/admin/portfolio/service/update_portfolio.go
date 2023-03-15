package service

import (
	"hause/admin/portfolio/dto"
	"hause/entity"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) UpdatePortfolio(ctx *gin.Context) {
	id := ctx.Param("id")
	var input dto.UpdatePortfolio
	var portfolio entity.Portfolio

	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	h.DB.Joins("JOIN user_links ON portfolio_id = portfolios.id AND user_id = ?", h.user.ID).
		First(&portfolio, id)

	if portfolio.ID == 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Could not find object by id"})
		return
	}

	update := entity.Portfolio{
		Name:      input.Name,
		CountryID: input.CountryID,
	}

	h.DB.Model(&portfolio).Updates(&update)

	ctx.JSON(http.StatusOK, portfolio)
}
