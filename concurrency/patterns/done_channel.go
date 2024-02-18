package patterns

import (
	"fmt"
	"time"
)

// This patterns is used to prevent go-routine leaks
func DoneChannel() {
	doWork := func(dataStream <-chan int, done <-chan interface{}) {
		for {
			select {
			case data, channelOpen := <-dataStream:
				if channelOpen {
					fmt.Printf("received data %v, working on it for 1 second\n", data)
					time.Sleep(1 * time.Second)
				} else {
					fmt.Printf("channel closed, waiting for 1 second\n")
					time.Sleep(1 * time.Second)
				}

			case <-done:
				fmt.Println("received termination signal, closing go routine")
				return
			}
		}
	}

	dataStream := make(chan int)
	doneChannel := make(chan interface{})
	go doWork(dataStream, doneChannel)

	timeoutAfterFiveSeconds := func() {
		time.Sleep(5 * time.Second)
		close(doneChannel)
	}

	go timeoutAfterFiveSeconds()

	dataStream <- 1
	dataStream <- 2
	dataStream <- 3
	close(dataStream)

	// simulate long running parent routine
	time.Sleep(10 * time.Second)
}
