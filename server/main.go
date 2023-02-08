package main

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var db *gorm.DB

func main() {

	r := gin.Default()
	db = Init()

	home := r.Group("/")
	addHouseRoutes(home)

	r.Run(":8001")
}
