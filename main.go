package main

import (
	"example/postman/routers"

	"github.com/gin-gonic/gin"
)

func main() {
	route := gin.Default()

	routers.Routers(route)
	route.Run(":8888")
}
