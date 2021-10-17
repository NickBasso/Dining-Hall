package table

import "dhall/src/components/types/order"

type Table struct {
	IsFree     bool
	HasOrdered bool
	Order      order.Order
}

func NewTable() Table {
	return Table{true, false, order.Order{}}
}