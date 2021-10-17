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
	MenuArray []food.Food
	MenuMap map[int]food.Food
	OrderMap map[int]order.Order
}