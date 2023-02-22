package auth

import (
	"hause/user"
	"hause/utility/token"
	"net/http"

	"github.com/gin-gonic/gin"
)

func AddRoutes(rg *gin.RouterGroup) {
	// Register new user
	rg.POST("/register", registerUser)

	// Refresh user token
	rg.POST("/refresh", refreshUser)

	// Refresh user token
	rg.POST("/login", loginUser)
}

func loginUser(ctx *gin.Context) {
	var reg register

	if err := ctx.ShouldBindJSON(&reg); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// hash the password and check against the matched email
	// respond with correct error message

	// on correct create new token and serve

}

func registerUser(ctx *gin.Context) {
	var reg register

	if err := ctx.ShouldBindJSON(&reg); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user := user.User{
		Email:    reg.Email,
		Password: reg.Password,
	}

	err := user.CreateUser()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	str := user.Email

	token, err := token.GenerateToken(str, 60)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// set to header the refresh token cookie

	ctx.JSON(http.StatusOK, gin.H{"token": token})
}

func refreshUser(ctx *gin.Context) {

}
