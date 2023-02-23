package main

import (
	"flag"
	"hause/house"
	"hause/portfolio"
	"hause/user"
	"hause/utility/auth"

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

	// accessible endpoints that wont be validated
	authRoute := r.Group("/api/auth")
	auth.AddRoutes(authRoute)

	// functional endpoints that use JWT middleware
	root := r.Group("/api")
	root.Use(auth.JWTAuthMiddleware())
	house.Setup(root, db)
	portfolio.Setup(root, db)
	user.Setup(root, db)

	port := ":8001"
	if *devRun {
		port = ":8002"
	}

	r.Run(port)
}
