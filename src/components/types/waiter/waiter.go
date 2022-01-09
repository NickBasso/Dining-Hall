package waiter

import "sync"

type Waiter struct {
	ID     int
	Name   string
	IsBusy bool
	Mutex  *sync.Mutex
}
