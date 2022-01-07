package controllers

import (
	"dininghall/src/components/constants"
	"dininghall/src/components/types/order"
	"dininghall/src/services"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
)

type Delivery order.Delivery

/* func placeOrder(c *gin.Context) {
	c.SetCookie("id", uuid.New().String(), 10, "http://localhost:4006/order", "http://localhost:4006/order", false, true)
	c.Redirect(http.StatusFound, "http://localhost:4006/order")
} */

var wg sync.WaitGroup

var ratingPoints = 0
var ordersCount = 0

func distributeOrder(c *gin.Context) {
	defer wg.Done()
	c.JSON(200, "DHall: Delivery received, distributing...");

	time.Sleep(constants.WaiterPickUpOrderTime * time.Second)
	var delivery Delivery;
	jsonDataRaw, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {}
	e := json.Unmarshal(jsonDataRaw, &delivery)
	if e != nil {}
	deliveryTime := time.Now().UnixMilli()
	fmt.Printf("POST delivery %s received, distributing...\n", delivery.OrderID)
	fmt.Printf("Delivery details:\n\t%v\n", delivery)
	fmt.Printf("Time delivered:\n\t%v\n", time.Now().UnixMilli())
	fmt.Printf("Time ordered:\n\t%v\n", delivery.PickUpTime)

	currentOrderRating := services.EvaluateDeliveryTimes(delivery.PickUpTime, deliveryTime, int64(delivery.MaxWait))
	ratingPoints += currentOrderRating
	ordersCount++

	services.FinishOrder(delivery.WaiterID, delivery.TableID)
	println("DONE > ==========================================================")
	fmt.Printf("Rating: %d stars\n\n", currentOrderRating)
	fmt.Printf("Rating points: %d\nOrdersCount: %d\n", ratingPoints, ordersCount)
	fmt.Printf("Average rating for all orders: %g \n\n\n", float32(ratingPoints) / float32(ordersCount))
}

/* func test(c *gin.Context) {
	services.GenerateOrders(constants.GeneratedOrdersCount)
} */

func simulateOrdersConsecutively(c *gin.Context) {
	// TODO: awaiting => waitGroup

	// generate orders infinitely
	// for i := 0; true; i++ {
	// 	wg.Add(1)
	// 	go services.GenerateOrder(i)
	// }
	go services.SimulateOrdersConsecutively(&wg);

	// fmt.Println("Main: Waiting for workers to finish")
	// wg.Wait()
	// fmt.Printf("Rating points: %d\nOrdersCount: %d\n", ratingPoints, ordersCount)
	// fmt.Printf("Average rating for all orders: %g \n\n\n", float32(ratingPoints) / float32(ordersCount))
}

// func getOrderList(c *gin.Context) {
// 	id := c.Query("id")
// 	items := c.Query("items")
// 	priority := c.Query("priority")
// 	maxWait := c.Query("maxWait")

// 	c.JSON(200, gin.H{
// 		"id":       id,
// 		"items":    items,
// 		"priority": priority,
// 		"maxWait":  maxWait,
// 	})
// }

func SetupController(router *gin.Engine) {
  services.InitCoreService()

	router.GET("/", simulateOrdersConsecutively)

	// router.GET("/test", test)
	router.POST("/distribution", distributeOrder)
}
