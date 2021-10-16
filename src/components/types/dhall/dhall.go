package dhall

import (
	"dhall/src/components/types/food"
	"dhall/src/components/types/order"
	"dhall/src/components/types/table"
	"dhall/src/components/types/waiter"
)

type DiningHall struct {
	Tables  []table.Table
	Waiters []waiter.Waiter
	Menu []food.Food
	OrderMap map[int16]order.Order
}