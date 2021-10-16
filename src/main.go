package main

import (
	"dhall/src/controllers"
	coreService "dhall/src/services"

	"github.com/gin-gonic/gin"
)

func main() {
	gin.ForceConsoleColor()
	router := gin.Default()

	coreService.InitCoreService()
	controllers.SetupController(router)

	router.Run(":4005")
}
