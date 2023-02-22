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

func getUser(ctx *gin.Context) {
	id := ctx.Param("id")

	var user User

	db.Find(&user, id)

	ctx.JSON(http.StatusOK, user)
}
