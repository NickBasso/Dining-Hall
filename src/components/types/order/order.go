package order

type Order struct {
	OrderID    string
	TableID    int
	WaiterID   int
	Items      []int
	Priority   int
	MaxWait    int
	PickUpTime int64
}