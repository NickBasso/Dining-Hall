package order

import "dininghall/src/components/types/order/props"

type Delivery struct {
	OrderID        string
	TableID        int
	WaiterID       int
	Items          []int
	Priority       int
	MaxWait        float64
	PickUpTime     int64
	CookingTime    int64
	CookingDetails []props.CookingDetail
}