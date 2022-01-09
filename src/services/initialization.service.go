package services

import (
	"dininghall/src/components/constants"
	"dininghall/src/components/types/food"
	"strconv"
	"sync"
)

func InitCoreService() {
	dhallRef = new(Dhall)
	
	dhallRef.Tables = make([]Table, constants.TablesCount)
	dhallRef.Waiters =  make([]Waiter, constants.WaitersCount)
	dhallRef.MenuArray = food.GetMenuArray()
	dhallRef.MenuMap = food.GetMenuMap()
	dhallRef.OrderMap = make(map[int]Order)

	for i := 0; i < len(dhallRef.MenuArray); i++ {
		dhallRef.OrderMap[dhallRef.MenuArray[i].ID] = Order{OrderID: strconv.Itoa(i)}
	}

	for i := 0; i < len(dhallRef.Tables); i++ {
		dhallRef.Tables[i] = Table{IsFree: true, HasOrdered: false, Order: Order{}, Mutex: new(sync.Mutex)}
	}

	for i := 0; i < len(dhallRef.Waiters); i++ {
		dhallRef.Waiters[i] = Waiter{ID: i, Name: constants.Names[i], Mutex: new(sync.Mutex)}
	}
}