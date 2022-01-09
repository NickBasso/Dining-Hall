package table

import (
	"dininghall/src/components/types/order"
	"sync"
)

type Table struct {
	IsFree     bool
	HasOrdered bool
	Order      order.Order
	Mutex 		 *sync.Mutex
}