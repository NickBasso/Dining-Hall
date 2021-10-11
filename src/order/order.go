package orders

type Order struct {
	id       string
	items    []int16
	priority byte
	maxWait  int
}