package patterns

import (
	"fmt"
	"time"
)

func ForSelectLoop() {
	c1 := make(chan interface{})

	// Emit a signal after 5 seconds
	go func() {
		time.Sleep(5 * time.Second)
		close(c1)
	}()

	// For-Select loop - continues to wait for signal
	for {
		received := false
		select {
		case <-c1:
			received = true
		default:
			fmt.Println("No signal, wait for one more second")
			time.Sleep(1 * time.Second)
		}
		if received {
			break
		}
	}

	fmt.Println("Recieved signal, terminate")
}
