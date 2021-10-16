package table

import "dhall/src/components/types/order"

type Table struct {
	IsFree     bool
	HasOrdered bool
	Order      order.Order
}

func NewTable() Table {
	return Table{true, false, order.Order{Id: "0", Items: nil, Priority: 0, MaxWait: 0}}
}