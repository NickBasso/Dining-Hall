package main

import (
	"dininghall/src/configs"
	"dininghall/src/controllers"
	"dininghall/src/services"
	"sync"

	"github.com/gin-gonic/gin"
)

var waitGroup sync.WaitGroup

func main() {
	gin.ForceConsoleColor()
	router := gin.Default()

	configs.SetupENV()
	services.InitCoreService()
	controllers.SetupController(router)

	router.Run(":4005")
}
