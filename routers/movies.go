package routers

import (
	"example/postman/controllers"
	"example/postman/middlewares"

	"github.com/gin-gonic/gin"
)

func MovieRouter(router *gin.RouterGroup) {
	router.GET("", controllers.GetAllMovies)
	router.GET("/:id", controllers.GetMovieById)

}

func MovieAdminRouter(router *gin.RouterGroup) {
	router.Use(middlewares.ValidateToken())
	router.POST("", controllers.AddMovie)
	router.PATCH("/:id", controllers.EditMovie)
	router.DELETE("/:id", controllers.DeleteMovie)
}
