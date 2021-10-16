package order

type Order struct {
	Id       string
	Items    []int16
	Priority byte
	MaxWait  int
}