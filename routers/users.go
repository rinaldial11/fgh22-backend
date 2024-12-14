package routers

import (
	"example/postman/controllers"
	"example/postman/middlewares"

	"github.com/gin-gonic/gin"
)

func UserRouter(router *gin.RouterGroup) {
	router.Use(middlewares.ValidateToken())
	router.GET("", controllers.GetAllUsers)
	router.GET("/:id", controllers.GetUserById)
	router.PATCH("/:id", controllers.EditUser)
	router.DELETE("/:id", controllers.DeleteUser)
}
