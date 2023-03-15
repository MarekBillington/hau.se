package service

import (
	"hause/entity"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) GetRelatedPortfolio(ctx *gin.Context) {
	var portfolio []entity.Portfolio

	h.DB.Joins("JOIN user_links ON portfolio_id = portfolios.id AND user_id = ?", h.user.ID).
		Order("portfolios.id asc").
		Find(&portfolio)

	ctx.JSON(http.StatusOK, portfolio)
}
