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

	// _, err := http.Post(os.Getenv("DHALL_URL")+"/", "application/json", nil)
	// if err != nil {}

	// services.SimulateOrdersConsecutively(&waitGroup)
	// controllers.SimulateOrdersConsecutively(nil)

	router.Run(":4005")
}
