package services

import (
	"bytes"
	"dininghall/src/components/constants"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"os"
	"time"

	"github.com/google/uuid"
)

func EvaluateDeliveryTimes(from int64, to int64, maxWait int64) int{
	timeElapsed := from - to;
	if maxWait > timeElapsed {
		return 5
	} else if float32(maxWait) * 1.1 > float32(timeElapsed) {
		return 4
	} else if float32(maxWait) * 1.2 > float32(timeElapsed) {
		return 3
	} else if float32(maxWait) * 1.3 > float32(timeElapsed) {
		return 2
	} else if float32(maxWait) * 1.4 > float32(timeElapsed) {
		return 1
	} else {
		return 0
	}
}

func GenerateOrder (idx int) {
	itemsCount := rand.Intn(constants.ItemsCap) + 1
	maxWait := 0

	items := make([]int, itemsCount)
	for j := 0; j < itemsCount; j++ {
		itemID := rand.Intn(constants.MenuCount) + 1
		items[j] = itemID 
		if dhallRef.MenuMap[itemID].PreparationTime > maxWait {
			maxWait = dhallRef.MenuMap[itemID].PreparationTime
		}
	}
	
	orderID := uuid.NewString()
	pickUpTime := time.Now().UnixMilli()
	priority := rand.Intn(constants.PriorityCap) + 1
	tableID := idx
	waiterID := 0

	isFreeWaiterAvailable := false

	for !isFreeWaiterAvailable {
		for i := 0; i < len(dhallRef.Waiters); i++ {
			if(dhallRef.Waiters[i].IsBusy == false) {
				dhallRef.Waiters[i].IsBusy = true
				isFreeWaiterAvailable = true
				break
			}
		}	
	}

	order := Order{Items: items, MaxWait: maxWait, OrderID: orderID, PickUpTime: pickUpTime, Priority: priority, TableID: tableID, WaiterID: waiterID}

	fmt.Printf("Order %s: %v\n", orderID, order)

	reqBody, reqBodySerializationErr := json.Marshal(order)
	if reqBodySerializationErr != nil {
		log.Fatalln(reqBodySerializationErr)
	}

	fmt.Printf("%v\n", order)
	time.Sleep(constants.WaiterPickUpOrderTime  * time.Second)
	resp, POSTErr := http.Post(os.Getenv("KITCHEN_URL")+"/order", "application/json", bytes.NewBuffer(reqBody))
	if POSTErr != nil {
		log.Fatalln(POSTErr)
	}

	defer resp.Body.Close()

	body, readPOSTResErr := ioutil.ReadAll(resp.Body)
	if readPOSTResErr != nil {
		log.Fatalln(readPOSTResErr)
	}

	var POSTOrderRes string;
	POSTResDeserializationErr := json.Unmarshal(body, &POSTOrderRes)
	if(POSTResDeserializationErr != nil) {
		log.Fatalln(POSTResDeserializationErr)
	}

	fmt.Printf("POST order: %s => %v\n\n", order.OrderID, POSTOrderRes)	
}
func GenerateOrders(amount int) /* []Order */ {
	orders := make([]Order, amount)

	for i := 0; i < amount; i++ {
		itemsCount := rand.Intn(constants.ItemsCap) + 1
		maxWait := 0

		items := make([]int, itemsCount)
		for j := 0; j < itemsCount; j++ {
			itemID := rand.Intn(constants.MenuCount) + 1
			items[j] = itemID 
			if dhallRef.MenuMap[itemID].PreparationTime > maxWait {
				maxWait = dhallRef.MenuMap[itemID].PreparationTime
			}
		}

		orders[i].Items = items
		orders[i].MaxWait = maxWait
		orders[i].OrderID = uuid.NewString()
		orders[i].PickUpTime = time.Now().UnixMilli()
		orders[i].Priority = rand.Intn(constants.PriorityCap) + 1
		orders[i].TableID = i + 1
		orders[i].WaiterID = i % constants.WaitersCount + 1

		fmt.Printf("Order %d: %v\n", i + 1, orders[i])
	}

	for i := 0; i < constants.GeneratedOrdersCount; i++ {
		reqBody, reqBodySerializationErr := json.Marshal(orders[i])
		if reqBodySerializationErr != nil {
			log.Fatalln(reqBodySerializationErr)
		}

		fmt.Printf("%v\n", orders[i])
		resp, POSTErr := http.Post(os.Getenv("KITCHEN_URL")+"/order", "application/json", bytes.NewBuffer(reqBody))
		if POSTErr != nil {
			log.Fatalln(POSTErr)
		}

		defer resp.Body.Close()

		body, readPOSTResErr := ioutil.ReadAll(resp.Body)
		if readPOSTResErr != nil {
			log.Fatalln(readPOSTResErr)
		}

		var POSTOrderRes string;
		POSTResDeserializationErr := json.Unmarshal(body, &POSTOrderRes)
		if(POSTResDeserializationErr != nil) {
			log.Fatalln(POSTResDeserializationErr)
		}

		fmt.Printf("POST order: %s => %v\n\n", orders[i].OrderID, POSTOrderRes)
	}

	// return orders
}

func FinishOrder(waiterID int) {
	dhallRef.Waiters[waiterID].IsBusy = false;
}