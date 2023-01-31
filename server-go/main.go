package main

import (
	"github.com/gin-gonic/gin"
)

func main() {

	r := gin.Default()

	home := r.Group("/")
	addHouseRoutes(home)

	r.Run(":8000")
}
