package main

import (
	"flag"
	"hause/auth"
	"hause/auth/middleware"
	"hause/house"
	"hause/portfolio"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var DB *gorm.DB

func main() {

	// For purpose of testing api before building to docker
	devRun := flag.Bool("isDev", false, "true if local, false for docker")
	flag.Parse()

	r := gin.Default()
	DB = Init(*devRun)

	// accessible endpoints that wont be validated
	authRoute := r.Group("/api/auth")
	auth.Setup(authRoute, DB)

	// functional endpoints that use JWT middleware
	root := r.Group("/api")
	root.Use(middleware.JWTAuth())

	house.Setup(root, DB)
	portfolio.Setup(root, DB)
	//user.Setup(root, db)

	port := ":8001"
	if *devRun {
		port = ":8002"
	}

	r.Run(port)
}
