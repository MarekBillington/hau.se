package auth

import (
	"hause/auth/service"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Setup(rg *gin.RouterGroup, conn *gorm.DB) {

	h := service.Handler{
		DB: conn,
	}

	// Register new user
	rg.POST("/register", h.RegisterUser)

	// Refresh user token
	rg.POST("/refresh", h.RefreshUser)

	// Refresh user token
	rg.POST("/login", h.LoginUser)

	// Logout
	rg.POST("/logout", service.LogoutUser)
}
