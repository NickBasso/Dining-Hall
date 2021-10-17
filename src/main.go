package main

import (
	"dhall/src/configs"
	"dhall/src/controllers"
	"dhall/src/services"

	"github.com/gin-gonic/gin"
)

func main() {
	gin.ForceConsoleColor()
	router := gin.Default()

	configs.SetupENV()
	services.InitCoreService()
	controllers.SetupController(router)
	// services.GenerateOrders(constants.GeneratedOrdersCount)

	router.Run(":4005")
}
