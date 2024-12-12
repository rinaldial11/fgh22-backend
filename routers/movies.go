package routers

import (
	"example/postman/controllers"

	"github.com/gin-gonic/gin"
)

func MovieRouter(router *gin.RouterGroup) {
	router.GET("", controllers.GetAllMovies)
	router.GET("/:id", controllers.GetMovieById)
	router.POST("", controllers.AddMovie)
	router.PATCH("/:id", controllers.EditMovie)
	router.DELETE("/:id", controllers.DeleteMovie)
}
