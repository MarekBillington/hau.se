package house

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var db *gorm.DB

func Setup(rg *gin.RouterGroup, conn *gorm.DB) {
	
	db = conn
	
	initDb()
	addRoutes(rg)
}
