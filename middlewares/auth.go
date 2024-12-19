package middlewares

import (
	"example/postman/models"
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/go-jose/go-jose/v4"
	"github.com/go-jose/go-jose/v4/jwt"
	"github.com/joho/godotenv"
)

func ValidateToken() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		head := ctx.GetHeader("Authorization")
		if head == "" {
			ctx.JSON(http.StatusUnauthorized, models.Response{
				Succsess: false,
				Message:  "Unauthorized",
			})
			ctx.Abort()
			return
		}
		token := strings.Split(head, " ")[1:][0]
		tok, _ := jwt.ParseSigned(token, []jose.SignatureAlgorithm{jose.HS256})
		out := jwt.Claims{}

		godotenv.Load()
		var SECRET_KEY = os.Getenv("SECRET_KEY")
		err := tok.Claims(SECRET_KEY, &out)

		if err != nil {
			ctx.JSON(http.StatusUnauthorized, models.Response{
				Succsess: false,
				Message:  "Unauthorized",
			})
			ctx.Abort()
		}

		ctx.Next()
	}
}
