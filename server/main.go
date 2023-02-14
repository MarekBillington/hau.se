package main

import (
	"flag"
	"hause/house"
	"hause/portfolio"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var db *gorm.DB

func main() {

	// For purpose of testing api before building to docker
	devRun := flag.Bool("isDev", false, "true if local, false for docker")
	flag.Parse()

	r := gin.Default()
	db = Init(*devRun)

	root := r.Group("/api")

	house.Setup(root, db)
	portfolio.Setup(root, db)

	r.Run(":8001")
}
