package controllers

import (
	coreService "dhall/src/services"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func placeOrder(c *gin.Context) {
	c.SetCookie("id", uuid.New().String(), 10, "http://localhost:4006/order", "http://localhost:4006/order", false, true)
	c.Redirect(http.StatusFound, "http://localhost:4006/order")
}

func getOrderList(c *gin.Context) {
	id := c.Query("id")
	items := c.Query("items")
	priority := c.Query("priority")
	maxWait := c.Query("maxWait")

	c.JSON(200, gin.H{
		"id":       id,
		"items":    items,
		"priority": priority,
		"maxWait":  maxWait,
	})
}

func SetupController(router *gin.Engine) {
  coreService.InitCoreService()

	router.GET("/", func(c *gin.Context) {
		c.JSON(200, "Dining hall server is up!")
	})

	router.POST("/order", placeOrder)
}
