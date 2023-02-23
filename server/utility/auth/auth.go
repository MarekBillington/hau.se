package auth

import (
	"hause/user"
	"hause/utility/token"
	"net/http"
	"time"

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

	exp := time.Now().Add(time.Minute * time.Duration(60)).Unix()

	// on correct create new token which is served as output
	str := user.Email
	t, err := token.GenerateToken(str, 60)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// refresh token is set to browser cookie
	refexp := time.Now().Add(time.Hour * time.Duration(24)).Unix()
	r, _ := token.GenerateToken(str, refexp)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// @todo tidy up url
	ctx.SetCookie("refresh", r, int(refexp), "/", "dev.hau.se", false, true)

	ctx.JSON(http.StatusOK, gin.H{"token": t, "expiry": exp})
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
	exp := time.Now().Add(time.Minute * time.Duration(60)).Unix()

	// create token that is sent for FE memory to handle
	t, err := token.GenerateToken(str, 60)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// create refresh token that is set to cookie
	refexp := time.Now().Add(time.Hour * time.Duration(24)).Unix()
	r, _ := token.GenerateToken(str, refexp)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// @todo tidy up url
	ctx.SetCookie("refresh", r, int(refexp), "/", "dev.hau.se", false, true)
	// set to header the refresh token cookie

	ctx.JSON(http.StatusOK, gin.H{"token": t, "expires": exp})
}

// Refresh access token with provided refresh token provided on first request
// Returns with new JWT and Refresh tokens
func refreshUser(ctx *gin.Context) {

	// get cookie from request, check it is valid
	refToken, err := ctx.Cookie("refresh")
	if err != nil {
		ctx.String(http.StatusUnauthorized, "Unauthorized")
		return
	}

	userEmail, err := token.TokenValid(refToken)
	if err != nil {
		ctx.String(http.StatusUnauthorized, "Unauthorized")
		return
	}

	var u user.User
	u.FindByEmail(userEmail)

	exp := time.Now().Add(time.Minute * time.Duration(60)).Unix()

	// on correct create new token which is served as output
	str := u.Email
	t, err := token.GenerateToken(str, 60)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"token": t, "expiry": exp})
}
