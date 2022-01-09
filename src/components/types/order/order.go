package order

import "sync"

type Order struct {
	OrderID    string
	TableID    int
	WaiterID   int
	Items      []int
	Priority   int
	MaxWait    float64
	PickUpTime int64
	Mutex 		 *sync.Mutex
}