package dhall

import (
	"dininghall/src/components/types/food"
	"dininghall/src/components/types/order"
	"dininghall/src/components/types/table"
	"dininghall/src/components/types/waiter"
)

type DiningHall struct {
	Tables  []table.Table
	Waiters []waiter.Waiter
	MenuArray []food.Food
	MenuMap map[int]food.Food
	OrderMap map[int]order.Order
}