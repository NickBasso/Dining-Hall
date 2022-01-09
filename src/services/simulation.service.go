package services

import (
	"dininghall/src/components/constants"
	"fmt"
	"sync"
)

func SimulateOrdersConsecutively(waitGroup *sync.WaitGroup) {
	for i := 0; true; i++ {
		if(pendingOrdersCounter < 2 * constants.TablesCount) {
			waitGroup.Add(1)
			pendingOrdersCounter++;
			fmt.Print("pendingOrdersCounter = ", pendingOrdersCounter)
			go GenerateOrder(i)
		}
	}

	waitGroup.Wait()
}

func EvaluateDeliveryTimes(from int64, to int64, maxWait int64) int {
	timeElapsed := (to - from) / 1000;
	
	fmt.Printf("from: %d  -  to: %d  -  maxWait: %g  -  timeElapsed: %g\n", from, to, float32(maxWait), float32(timeElapsed))
	if maxWait > timeElapsed {
		return 5
	} else if float32(maxWait) * 1.1 > float32(timeElapsed) {
		return 4
	} else if float32(maxWait) * 1.2 > float32(timeElapsed) {
		return 3
	} else if float32(maxWait) * 1.3 > float32(timeElapsed) {
		return 2
	} else if float32(maxWait)* 1.4 > float32(timeElapsed)  {
		return 1
	} else {
		return 0
	}
}