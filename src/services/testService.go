package services

import (
	"bytes"
	"dhall/src/components/constants"
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

func GenerateOrders(amount int) []Order {
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
		orders[i].PickUpTime = time.Now().Unix()
		orders[i].Priority = rand.Intn(constants.PriorityCap) + 1
		orders[i].TableID = i + 1
		orders[i].WaiterID = i % constants.WaitersCount + 1

		fmt.Printf("Order %d: %v\n", i + 1, orders[i])
	}

	/*

				"order_id": 1,
		"table_id": 1,
		"waiter_id": 1,
		"items": [ 3, 4, 4, 2 ],
		"priority": 3,
		"max_wait": 45,
		"pick_up_time": 1631453140 // UNIX timestamp

	*/

	reqBody, err := json.Marshal(orders)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Printf("%v", orders)

	for i := 0; i < int(len(reqBody)); i++ {
		println(i , ": ", reqBody[i])
	}

	resp, err := http.Post(os.Getenv("KITCHEN_URL")+"/order", "application/json", bytes.NewBuffer(reqBody))
	if err != nil {
		log.Fatalln(err)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	log.Println(string(body))

	return orders
}