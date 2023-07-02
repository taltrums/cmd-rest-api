package main

import (
	"api/controllers"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.POST("/api/cmd", controllers.ExecuteShellCommandHandler)

	router.Run(":8080")
}
