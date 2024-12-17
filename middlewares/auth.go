package middlewares

import (
	"example/postman/controllers"
	"example/postman/lib"
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/go-jose/go-jose/v4"
	"github.com/go-jose/go-jose/v4/jwt"
)

func ValidateToken() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		head := ctx.GetHeader("Authorization")
		if head == "" {
			ctx.JSON(http.StatusUnauthorized, controllers.Response{
				Succsess: false,
				Message:  "Unauthorized",
			})
			ctx.Abort()
			return
		}
		token := strings.Split(head, " ")[1:][0]
		fmt.Println(token)
		tok, _ := jwt.ParseSigned(token, []jose.SignatureAlgorithm{jose.HS256})
		out := jwt.Claims{}

		err := tok.Claims(lib.JWT_SECRET, &out)

		if err != nil {
			ctx.JSON(http.StatusUnauthorized, controllers.Response{
				Succsess: false,
				Message:  "Unauthorized",
			})
			ctx.Abort()
		}

		ctx.Next()
	}
}
