package service

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func LogoutUser(ctx *gin.Context) {
	ctx.SetCookie("refresh", "", 0, "/", "dev.hau.se", false, true)

	ctx.JSON(http.StatusOK, "logged out")
}
