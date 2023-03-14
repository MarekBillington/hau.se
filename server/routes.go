package main

import (
	portfolio "hause/admin/portfolio"
	user "hause/admin/user"
	"hause/house"
	"hause/utility/country"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func AddRoutes(rootRoute *gin.RouterGroup, DB *gorm.DB) {

	// api/house/...
	house.Setup(rootRoute, DB)

	// api/user/..
	user.Setup(rootRoute, DB)

	// api/portfolio/..
	portfolio.Setup(rootRoute, DB)

	// api/utility/country/..
	country.Setup(rootRoute, DB)
}
