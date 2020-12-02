package examples

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestSimpleThreadpool(t *testing.T) {
	processedElementsChannel := make(chan int)
	processedElements := make([]int, 0)
	exit := make(chan struct{})

	go func() {
		for {
			select {
			case v := <-processedElementsChannel:
				processedElements = append(processedElements, v)
			case <-exit:
				return
			}
		}
	}()

	options := &SimpleThreadpoolOptions{
		NThreads:                 2,
		GeneratorSleepMs:         1,
		LinebreakSleepMs:         0,
		WorkerSleepMs:            10,
		EnableLinebreaks:         false,
		Printf:                   stubPrintf,
		Exit:                     exit,
		Debug:                    true,
		ProcessedElementsChannel: processedElementsChannel,
	}
	SimpleThreadpool(*options)

	time.Sleep(time.Duration(20) * time.Millisecond)
	close(exit)

	expectedElements := []int{0, 1, 2, 3}
	for _, element := range expectedElements {
		assert.Contains(t, processedElements, element)
	}
}

func stubPrintf(format string, a ...interface{}) (n int, err error) {
	return
}
