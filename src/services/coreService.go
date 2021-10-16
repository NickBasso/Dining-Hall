package coreService

import (
	"dhall/src/components/constants"
	"dhall/src/components/types/apparatus"
	"dhall/src/components/types/dhall"
	"dhall/src/components/types/food"
	"dhall/src/components/types/order"
	"dhall/src/components/types/table"
	"dhall/src/components/types/waiter"
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
	fillFoodList()
	fillOrderMap()

	for i := 0; i < len(dhallRef.Tables); i++ {
		dhallRef.Tables[i] = Table{IsFree: true, HasOrdered: false, Order: Order{}}
	}

	for i := 0; i < len(dhallRef.Waiters); i++ {
		dhallRef.Waiters[i] = Waiter{Id: byte(i), Name: constants.Names[i]}
	}
}

func HandleOrder() {

}

func fillFoodList () {
	menu := append(make([]Food, 0), Food{
		Id:              1,
		Name:            "pizza",
		PreparationTime: 20,
		Complexity:      2,
		Apparatus:       apparatus.Oven},
		Food{
			Id:              2,
			Name:            "salad",
			PreparationTime: 10,
			Complexity:      1,
			Apparatus:       apparatus.None},
		Food{
			Id:              3,
			Name:            "zeama",
			PreparationTime: 7,
			Complexity:      1,
			Apparatus:       apparatus.Stove},
		Food{
			Id:              4,
			Name:            "Scallop Sashimi with Meyer Lemon Confit",
			PreparationTime: 32,
			Complexity:      3,
			Apparatus:       apparatus.None},
		Food{
			Id:              5,
			Name:            "Island Duck with Mulberry Mustard",
			PreparationTime: 35,
			Complexity:      3,
			Apparatus:       apparatus.Oven},
		Food{
			Id:              6,
			Name:            "Waffles",
			PreparationTime: 10,
			Complexity:      1,
			Apparatus:       apparatus.Stove},
		Food{
			Id:              7,
			Name:            "Aubergine",
			PreparationTime: 20,
			Complexity:      2,
			Apparatus:       apparatus.None},
		Food{
			Id:              8,
			Name:            "Lasagna",
			PreparationTime: 30,
			Complexity:      2,
			Apparatus:       apparatus.Oven},
		Food{
			Id:              9,
			Name:            "Burger",
			PreparationTime: 15,
			Complexity:      1,
			Apparatus:       apparatus.Oven},
		Food{
			Id:              10,
			Name:            "Gyros",
			PreparationTime: 15,
			Complexity:      1,
			Apparatus:       apparatus.None})

			dhallRef.Menu = menu
}

func fillOrderMap () {
	dhallRef.OrderMap = make(map[int16]Order)

	for i := 0; i < len(GetMenu()); i++ {
		dhallRef.OrderMap[dhallRef.Menu[i].Id] = Order{Id: string(i)}
	}
}

func GetMenu() []Food {
	return dhallRef.Menu
}

func GetOrderMap() map[int16]Order {
	return dhallRef.OrderMap
}