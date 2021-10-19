package services

import (
	"dininghall/src/components/constants"
	"dininghall/src/components/types/apparatus"
	"dininghall/src/components/types/dhall"
	"dininghall/src/components/types/food"
	"dininghall/src/components/types/order"
	"dininghall/src/components/types/table"
	"dininghall/src/components/types/waiter"
)

type Dhall = dhall.DiningHall
type Apparatus = apparatus.Apparatus
type Food = food.Food
type Order = order.Order
type Table = table.Table
type Waiter = waiter.Waiter

var dhallRef *Dhall = nil

func InitCoreService() {
	dhallRef = new(Dhall)
	
	dhallRef.Tables = make([]Table, constants.TablesCount, constants.TablesCount * 2)
	dhallRef.Waiters =  make([]Waiter, constants.WaitersCount, constants.WaitersCount * 2)
	dhallRef.MenuArray = food.GetMenuArray()
	dhallRef.MenuMap = food.GetMenuMap()
	fillOrderMap()

	for i := 0; i < len(dhallRef.Tables); i++ {
		dhallRef.Tables[i] = Table{IsFree: true, HasOrdered: false, Order: Order{}}
	}

	for i := 0; i < len(dhallRef.Waiters); i++ {
		dhallRef.Waiters[i] = Waiter{ID: i, Name: constants.Names[i]}
	}
}

func fillOrderMap () {
	dhallRef.OrderMap = make(map[int]Order)

	for i := 0; i < len(dhallRef.MenuArray); i++ {
		dhallRef.OrderMap[dhallRef.MenuArray[i].ID] = Order{OrderID: string(i)}
	}
}

func getMenuArray() []Food {
	return dhallRef.MenuArray
}

func getOrderMap() map[int]Order {
	return dhallRef.OrderMap
}