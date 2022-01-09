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

var waitGroup sync.WaitGroup
var ratingPoints = 0
var ordersCount = 0

func distributeOrder(c *gin.Context) {
	c.JSON(200, "DHall: Delivery received, distributing...");

	time.Sleep(constants.WaiterPickUpOrderTime * time.Second)
	var delivery Delivery;
	jsonDataRaw, _ := ioutil.ReadAll(c.Request.Body)
	_ = json.Unmarshal(jsonDataRaw, &delivery)
	deliveryTime := time.Now().UnixMilli()

	fmt.Printf("POST delivery %s received, distributing...\n", delivery.OrderID)
	fmt.Printf("Delivery details:\n\t%v\n", delivery)
	fmt.Printf("Time delivered:\n\t%v\n", time.Now().UnixMilli())
	fmt.Printf("Time ordered:\n\t%v\n", delivery.PickUpTime)

	currentOrderRating := services.EvaluateDeliveryTimes(delivery.PickUpTime, deliveryTime, int64(delivery.MaxWait))
	ratingPoints += currentOrderRating
	ordersCount++

	services.FinishOrder(delivery.WaiterID, delivery.TableID, &waitGroup)
	println("DONE > ==========================================================")
	fmt.Printf("Rating: %d stars\n\n", currentOrderRating)
	fmt.Printf("Rating points: %d\nOrdersCount: %d\n", ratingPoints, ordersCount)
	fmt.Printf("Average rating for all orders: %g \n\n\n", float32(ratingPoints) / float32(ordersCount))
}

func simulateOrdersConsecutively(c *gin.Context) {
	services.SimulateOrdersConsecutively(&waitGroup);
	waitGroup.Wait()

	fmt.Println("!!!For some reason all go routines finished executing!!!")
}

func SetupController(router *gin.Engine) {
	router.GET("/", simulateOrdersConsecutively)
	router.POST("/distribution", distributeOrder)
}
