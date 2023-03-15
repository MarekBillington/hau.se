package admin

import (
	"hause/admin/user/service"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Setup(rg *gin.RouterGroup, conn *gorm.DB) {

	h := service.Handler{
		DB: conn,
	}

	user := rg.Group("/user")

	user.GET("/setup", h.Setup, h.GetUserForSetup)

	user.GET("", h.Setup, h.GetAllUsers)

	user.GET(":id", h.Setup, h.GetUserById)
}
