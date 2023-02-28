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
	var login login

	if err := ctx.ShouldBindJSON(&login); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var user user.User
	user.FindByEmail(login.Email)

	// hash the password and check against the matched email
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(login.Password))
	if err != nil {
		// respond with correct error message
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// on correct create new token which is served as output
	t, exp, err := serveJWT(user, ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// refresh token is set to browser cookie
	err = serveRefCookie(user, ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

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
		Email:     reg.Email,
		Password:  reg.Password,
		FirstName: reg.FirstName,
		LastName:  reg.LastName,
	}

	err := user.CreateUser()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	t, exp, err := serveJWT(user, ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = serveRefCookie(user, ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
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

	t, exp, err := serveJWT(u, ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"token": t, "expiry": exp})
}

// Creates the refresh token and sends to set as cookie in response
func serveRefCookie(u user.User, ctx *gin.Context) error {

	refexp := time.Now().Add(time.Hour * time.Duration(24)).Unix()
	r, err := token.GenerateToken(u.Email, refexp)
	if err != nil {
		return err
	}

	// @todo tidy up url and for prod set to secure
	ctx.SetCookie("refresh", r, int(refexp), "/", "dev.hau.se", false, true)
	return nil
}

// Creates the auth token used as bearer for rest of api
func serveJWT(u user.User, ctx *gin.Context) (string, int64, error) {
	str := u.Email
	exp := time.Now().Add(time.Minute * time.Duration(10)).Unix()

	// on correct create new token which is served as output
	t, err := token.GenerateToken(str, exp)
	if err != nil {
		return "", 0, err
	}

	return t, exp, nil
}
