package main

import (
	"github.com/gin-gonic/gin"
	"github.com/MarekBillington/hau.se/server/house"
	"gorm.io/gorm"
)

var db *gorm.DB

func main() {

	r := gin.Default()
	db = Init()

	root := r.Group("/api")

	house.Setup(root, db)

	r.Run(":8001")
}
