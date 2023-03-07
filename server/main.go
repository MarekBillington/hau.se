package main

import (
	"flag"
	portfolio "hause/admin/portfolio"
	user "hause/admin/user"
	"hause/auth"
	security "hause/auth/middleware"
	"hause/house"
	"hause/utility/database"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var DB *gorm.DB

func main() {

	// For purpose of testing api before building to docker
	devRun := flag.Bool("isDev", false, "true if local, false for docker")
	flag.Parse()

	r := gin.Default()
	// Can this be moved to a entity utility that the repo's can access, allowing for the
	DB = database.Init(*devRun)
	database.SetupDb(DB)

	// accessible endpoints that wont be validated
	authRoute := r.Group("/api/auth")
	auth.Setup(authRoute, DB)

	// functional endpoints that use JWT middleware
	root := r.Group("/api")
	root.Use(security.JWTAuth())

	house.Setup(root, DB)
	user.Setup(root, DB)
	portfolio.Setup(root, DB)

	port := ":8001"
	if *devRun {
		port = ":8002"
	}

	r.Run(port)
}
