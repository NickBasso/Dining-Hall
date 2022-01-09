package services

import (
	"bytes"
	"dininghall/src/components/constants"
	"dininghall/src/components/types/apparatus"
	"dininghall/src/components/types/dhall"
	"dininghall/src/components/types/food"
	"dininghall/src/components/types/order"
	"dininghall/src/components/types/table"
	"dininghall/src/components/types/waiter"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"os"
	"sync"
	"time"

	"github.com/google/uuid"
)

type Dhall = dhall.DiningHall
type Apparatus = apparatus.Apparatus
type Food = food.Food
type Order = order.Order
type Table = table.Table
type Waiter = waiter.Waiter

var dhallRef *Dhall = nil
var pendingOrdersCounter = 0

func GenerateOrder(idx int) {
	isFreeTableAvailable := false
	isFreeWaiterAvailable := false

	orderID := uuid.NewString()
	pickUpTime := time.Now().UnixMilli()
	priority := rand.Intn(constants.PriorityCap) + 1
	tableID := 0
	waiterID := 0

	for !isFreeTableAvailable {
		for i := 0; i < len(dhallRef.Tables); i++ {
			dhallRef.Tables[i].Mutex.Lock()

			if(dhallRef.Tables[i].IsFree) {
				dhallRef.Tables[i].IsFree = false
				isFreeTableAvailable = true
				dhallRef.Tables[i].Mutex.Unlock()
				tableID = i
				break
			}

			dhallRef.Tables[i].Mutex.Unlock()
		}

		time.Sleep(50 * time.Millisecond)
	}

	itemsCount := rand.Intn(constants.ItemsCap) + 1
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

	for !isFreeWaiterAvailable {
		for i := 0; i < len(dhallRef.Waiters); i++ {
			dhallRef.Waiters[i].Mutex.Lock()

			if(!dhallRef.Waiters[i].IsBusy) {
				dhallRef.Waiters[i].IsBusy = true
				isFreeWaiterAvailable = true
			  dhallRef.Waiters[i].Mutex.Unlock()
				waiterID = i
				break
			}

			dhallRef.Waiters[i].Mutex.Unlock()
		}	

		time.Sleep(50 * time.Millisecond)
		dhallRef.Waiters[waiterID].Mutex.Lock()
		dhallRef.Waiters[waiterID].IsBusy = false
		dhallRef.Waiters[waiterID].Mutex.Unlock()	
	}

	order := Order {Items: items, MaxWait: maxFloat, OrderID: orderID, PickUpTime: pickUpTime,
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
}

func FinishOrder(waiterID int, tableID int, waitGroup *sync.WaitGroup) {
	dhallRef.Waiters[waiterID].Mutex.Lock()
	defer dhallRef.Waiters[waiterID].Mutex.Unlock()
	dhallRef.Waiters[waiterID].IsBusy = false;

	fmt.Printf("tableID = %d", tableID)
	dhallRef.Tables[tableID].Mutex.Lock()
	dhallRef.Tables[tableID].IsFree = true;
	dhallRef.Tables[tableID].HasOrdered = false;
	dhallRef.Tables[tableID].Mutex.Unlock()

	pendingOrdersCounter--;
	waitGroup.Done()
}