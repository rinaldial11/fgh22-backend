package middlewares

import (
	"example/postman/controllers"
	"net/http"

	"github.com/gin-gonic/gin"
)

func SetAuthorize() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		isLog := ctx.GetHeader("X-Logged-User")
		if isLog != "true" {
			ctx.JSON(http.StatusUnauthorized, controllers.Response{
				Succsess: false,
				Message:  "Unauthorized",
			})
			ctx.Abort()
			return
		}
		ctx.Next()
	}
}
