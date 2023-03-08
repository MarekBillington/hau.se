package middleware

import (
	"hause/auth/utility"
	"net/http"

	"github.com/gin-gonic/gin"
)

func JWTAuth() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		_, e := utility.ValidateToken(ctx)
		if e != nil {
			//fmt.Fprintf(os.Stdout, "%+v", e)
			ctx.String(http.StatusUnauthorized, e.Error())
			ctx.Abort()
			return
		}
		ctx.Next()
	}
}
