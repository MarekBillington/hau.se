package service

import (
	"hause/auth/utility"
	"hause/entity"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Handler struct {
	DB   *gorm.DB
	user entity.User
}

func (h *Handler) Setup(ctx *gin.Context) {
	user, err := utility.GetRequestingUser(ctx, h.DB)
	if err != nil {
		return
	}
	h.user = user
}
