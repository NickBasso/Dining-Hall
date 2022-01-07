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

// var ratingPoints = 0
// var ordersCount = 0

func EvaluateDeliveryTimes(from int64, to int64, maxWait int64) int{
	
	timeElapsed := (to - from) / 1000;
	fmt.Printf("from: %d  -  to: %d  -  maxWait: %g  -  timeElapsed: %g\n", from, to, float32(maxWait), float32(timeElapsed))
	if maxWait > timeElapsed {
		return 5
	} else if float32(maxWait) * 1.1 > float32(timeElapsed) {
		return 4
	} else if float32(maxWait) * 1.2 > float32(timeElapsed) {
		return 3
	} else if float32(maxWait) * 1.3 > float32(timeElapsed) {
		return 2
	} else if float32(maxWait)* 1.4 > float32(timeElapsed)  {
		return 1
	} else {
		return 0
	}
}

func GenerateOrder (idx int) {
	isFreeTableAvailable := false
	isFreeWaiterAvailable := false


	for !isFreeTableAvailable {
		for i := 0; i < len(dhallRef.Tables); i++ {
			if(dhallRef.Tables[i].IsFree) {
				dhallRef.Tables[i].IsFree = true
				isFreeTableAvailable = true
				break
			}
		}
	}

	itemsCount := rand.Intn(constants.ItemsCap) + 1
	// anti-null value mechanism
	maxWait := 0
	maxFloat := 0.0

	items := make([]int, itemsCount)
	for j := 0; j < itemsCount; j++ {
		itemID := rand.Intn(constants.MenuCount) + 1
		items[j] = itemID 
		if dhallRef.MenuMap[itemID].PreparationTime > maxWait {
			maxWait = dhallRef.MenuMap[itemID].PreparationTime
		}
	}
	maxFloat = float64(maxWait) * 1.3

	orderID := uuid.NewString()
	pickUpTime := time.Now().UnixMilli()
	priority := rand.Intn(constants.PriorityCap) + 1
	tableID := idx
	waiterID := 0

	for !isFreeWaiterAvailable {
		for i := 0; i < len(dhallRef.Waiters); i++ {
			if(!dhallRef.Waiters[i].IsBusy) {
				dhallRef.Waiters[i].IsBusy = true
				isFreeWaiterAvailable = true
				break
			}
		}	
	}

	order := Order{Items: items, MaxWait: maxFloat, OrderID: orderID, PickUpTime: pickUpTime,
		 Priority: priority, TableID: tableID, WaiterID: waiterID}

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

	// fmt.Printf("POST order: %s => %v\n\n", order.OrderID, POSTOrderRes)	
	// fmt.Printf("Rating points: %d\nOrdersCount: %d\n", ratingPoints, ordersCount)
	// fmt.Printf("Average rating for all orders: %g \n\n\n", float32(ratingPoints) / float32(ordersCount))
}

func FinishOrder(waiterID int, tableID int) {
	dhallRef.Waiters[waiterID].IsBusy = false;
	dhallRef.Tables[tableID].IsFree = true;
	dhallRef.Tables[tableID].HasOrdered = false;
}