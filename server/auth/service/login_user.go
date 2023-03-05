package service

import (
	"hause/auth/dto"
	"hause/auth/repository"
	"hause/auth/utility"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

// Login user with provided Email and Password
// Returns JWT for use as Bearer token
func (h Handler) LoginUser(ctx *gin.Context) {
	var login dto.Login

	if err := ctx.ShouldBindJSON(&login); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user := repository.GetUserFromEmail(login.Email, h.DB)
	if user.ID == 0 {
		ctx.String(http.StatusUnauthorized, "Unauthorized")
		return
	}

	// hash the password and check against the matched email
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(login.Password))
	if err != nil {
		// respond with correct error message
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// on correct create new token which is served as output
	t, exp, err := utility.ServeJWT(user.Email, ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// refresh token is set to browser cookie
	err = utility.ServeRefCookie(user.Email, ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"token": t, "expiry": exp})
}
