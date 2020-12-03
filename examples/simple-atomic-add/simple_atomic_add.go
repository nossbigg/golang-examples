package examples

import (
	"sync/atomic"
)

// SimpleAtomicAdd ...
func SimpleAtomicAdd(values chan int32, exit chan struct{}) {
	var currentValue int32 = 0

	workers := 100
	for i := 0; i < workers; i++ {
		go doIncrementWorker2(values, exit, &currentValue)
	}

}

func doIncrementWorker2(values chan int32, exit chan struct{}, currentValue *int32) {
	for {
		select {
		case <-exit:
			return
		default:
			atomic.AddInt32(currentValue, 1)
			values <- *currentValue
		}
	}
}
