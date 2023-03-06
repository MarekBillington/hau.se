package utility

import (
	"hause/auth/repository"
	"hause/entity"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GetRequestingUser(ctx *gin.Context, db *gorm.DB) (entity.User, error) {
	signature, err := ValidateToken(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err})
		return entity.User{}, err
	}

	user := repository.GetUserFromEmail(signature, db)
	return user, nil
}
