package user

import (
	"net/http"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func addRoutes(rg *gin.RouterGroup) {

	user := rg.Group("/user")

	// Register new user
	user.POST("/register", func(ctx *gin.Context) {
		registerUser(ctx)
	})

	// Refresh user token
	user.POST("/refresh", func(ctx *gin.Context) {
		refreshUser(ctx)
	})
}

func loginUser(ctx *gin.Context) {
	var reg register

	if err := ctx.ShouldBindJSON(&reg); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

}

func registerUser(ctx *gin.Context) {
	var reg register

	if err := ctx.ShouldBindJSON(&reg); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user := user{
		Email:    reg.Email,
		Password: reg.Password,
	}

	err := db.Create(&user).Error
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// create token and respond
	token, _ := generateToken(&user.ID)

	ctx.JSON(http.StatusOK, gin.H{"token": token})
}

func generateToken(user_id *int) (string, error) {

	claims := jwt.MapClaims{}
	claims["authorised"] = true
	claims["user_id"] = user_id
	claims["exp"] = time.Now().Add(time.Hour * time.Duration(1)).Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// @todo env replace
	return token.SignedString("shhh it's a secret")
}

func refreshUser(ctx *gin.Context) {

}
