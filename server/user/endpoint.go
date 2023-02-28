package user

import (
	"hause/utility/token"
	"net/http"

	"github.com/gin-gonic/gin"
)

func addRoutes(rg *gin.RouterGroup) {

	user := rg.Group("/user")

	// Used to get user from the auth
	user.GET("/setup", getUserFromAuth)

	user.GET(":id", getUser)
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

// Get user when auth is available
func getUserFromAuth(ctx *gin.Context) {

	t, err := token.ExtractBearerToken(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	email, err := token.TokenValid(t)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	var user User
	user.FindByEmail(email)

	user.Password = ""

	ctx.JSON(http.StatusOK, user)
}
