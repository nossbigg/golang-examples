package examples

import (
	"sync"
	"time"
)

// SimpleMutex ...
func SimpleMutex(values chan int, exit chan struct{}) {
	currentValue := 0
	mtx := &sync.Mutex{}

	workers := 100
	for i := 0; i < workers; i++ {
		go doIncrementWorker(values, exit, mtx, &currentValue)
	}

}

func doIncrementWorker(values chan int, exit chan struct{}, mtx *sync.Mutex, currentValue *int) {
	for {
		select {
		case <-exit:
			return
		default:
			mtx.Lock()

			newValue := *currentValue + 1
			// simulate delay
			time.Sleep(1 * time.Millisecond)
			*currentValue = newValue
			values <- *currentValue

			mtx.Unlock()
		}
	}
}
