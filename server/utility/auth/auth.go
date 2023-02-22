package auth

import (
	"hause/user"
	"hause/utility/token"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func AddRoutes(rg *gin.RouterGroup) {
	// Register new user
	rg.POST("/register", registerUser)

	// Refresh user token
	rg.POST("/refresh", refreshUser)

	// Refresh user token
	rg.POST("/login", loginUser)
}

// Login user with provided Email and Password
// Returns JWT for use as Bearer token
func loginUser(ctx *gin.Context) {
	var reg register

	if err := ctx.ShouldBindJSON(&reg); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var user user.User

	user.FindByEmail(reg.Email)

	// hash the password and check against the matched email
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(reg.Password))
	if err != nil {
		// respond with correct error message
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// on correct create new token and serve
	str := user.Email
	token, err := token.GenerateToken(str, 60)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"token": token})
}

// Register creates the user with a provided email and password
// Returns JWT for use as Bearer token
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

// Refresh access token with provided refresh token provided on first request
// Returns with new JWT and Refresh tokens
func refreshUser(ctx *gin.Context) {
	// @todo
}
