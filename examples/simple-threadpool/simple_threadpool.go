package examples

import (
	"time"
)

// SimpleThreadpoolOptions ...
type SimpleThreadpoolOptions struct {
	NThreads                 int
	GeneratorSleepMs         int
	LinebreakSleepMs         int
	WorkerSleepMs            int
	EnableLinebreaks         bool
	Printf                   func(format string, a ...interface{}) (n int, err error)
	Exit                     chan struct{}
	Debug                    bool
	ProcessedElementsChannel chan<- int
}

// SimpleThreadpool ...
func SimpleThreadpool(options SimpleThreadpoolOptions) {
	inputChannel := make(chan int)

	if options.EnableLinebreaks {
		go startLinebreaks(options)
	}

	go startNumberGenerator(inputChannel, options)
	go startThreadWorkers(inputChannel, options)
}

func startNumberGenerator(ch chan<- int, options SimpleThreadpoolOptions) {
	for i := 0; ; i++ {
		select {
		case ch <- i:
			<-threadSleep(options.GeneratorSleepMs)
			options.Printf("+%d ", i)
			continue
		case <-options.Exit:
			if options.Debug {
				options.Printf("generator: exited\n")
			}
			return
		}
	}
}

func startLinebreaks(options SimpleThreadpoolOptions) {
	for {
		select {
		case <-threadSleep(options.LinebreakSleepMs):
			options.Printf("\n")
			continue
		case <-options.Exit:
			if options.Debug {
				options.Printf("linebreak: exited\n")
			}
			return
		}
	}
}

func startThreadWorkers(ch chan int, options SimpleThreadpoolOptions) {
	for i := 0; i < options.NThreads; i++ {
		go makeThreadWorker(ch, options)
	}
}

func makeThreadWorker(ch <-chan int, options SimpleThreadpoolOptions) {
	for {
		select {
		case value := <-ch:
			options.ProcessedElementsChannel <- value
			options.Printf("-%d ", value)
			<-threadSleep(options.WorkerSleepMs)
			continue
		case <-options.Exit:
			if options.Debug {
				options.Printf("worker: exited\n")
			}
			return
		}

	}
}

func threadSleep(ms int) <-chan time.Time {
	return time.After(time.Duration(ms) * time.Millisecond)
}
