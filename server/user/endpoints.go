package user

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func addRoutes(rg *gin.RouterGroup) {

	user := rg.Group("/user")

	user.GET(":id", func(ctx *gin.Context) {
		getUser(ctx)
	})
}

// Get User by Id
// No password provided on response
func getUser(ctx *gin.Context) {
	id := ctx.Param("id")

	var user User

	db.Find(&user, id)

	user.Password = ""

	ctx.JSON(http.StatusOK, user)
}
