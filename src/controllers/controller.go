package controllers

import (
	"dininghall/src/components/constants"
	"dininghall/src/components/types/order"
	"dininghall/src/services"
	coreService "dininghall/src/services"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"time"

	"github.com/gin-gonic/gin"
)

type Delivery order.Delivery

/* func placeOrder(c *gin.Context) {
	c.SetCookie("id", uuid.New().String(), 10, "http://localhost:4006/order", "http://localhost:4006/order", false, true)
	c.Redirect(http.StatusFound, "http://localhost:4006/order")
} */

func distributeOrder(c *gin.Context) {
	c.JSON(200, "DHall: Delivery received, distributing...");

	var delivery Delivery;
	jsonDataRaw, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {}
	e := json.Unmarshal(jsonDataRaw, &delivery)
	if e != nil {}
	fmt.Printf("POST delivery %s received, distributing...\n", delivery.OrderID)
	fmt.Printf("Delivery details:\n\t%v\n", delivery)
	fmt.Printf("Time delivered:\n\t%v\n", time.Now().UnixMilli())
	fmt.Printf("Time ordered:\n\t%v\n", delivery.PickUpTime)

	fmt.Printf("Rating: %d stars\n\n", services.EvaluateDeliveryTimes(delivery.PickUpTime, time.Now().UnixMilli(), int64(delivery.MaxWait)))
}

func test(c *gin.Context) {
	services.GenerateOrders(constants.GeneratedOrdersCount)
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

	router.GET("/test", test)
	router.POST("/distribution", distributeOrder)
}
