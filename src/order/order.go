package orders

type Order struct {
	id       string
	items    []int
	priority string
	maxWait  int
}

/* func (e Orders) LeavesRemaining() {
    fmt.Printf("%s %s has %d leaves remaining\n", e.FirstName, e.LastName, (e.TotalLeaves - e.LeavesTaken))
} */