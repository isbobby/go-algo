package validations

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func workA() error {
	r := rand.Intn(3)
	fmt.Printf("[A]:%v seconds\n", r+1)
	time.Sleep(time.Duration(r+1) * time.Second)
	return nil
}

func workB() error {
	r := rand.Intn(3)
	fmt.Printf("[B]:%v seconds\n", r+1)
	time.Sleep(time.Duration(r+1) * time.Second)
	return nil
}

func Invoke() {
	start := time.Now()

	var wg sync.WaitGroup
	wg.Add(1)

	var newWorkTime time.Duration
	go func(start time.Time) {
		defer wg.Done()
		workB()
		newWorkTime = time.Since(start)
	}(start)

	// original work
	workA()
	originalWorkTime := time.Since(start)
	wg.Wait()

	fmt.Printf("Time difference (og - new):%v, Original:%v, New:%v\n", originalWorkTime-newWorkTime, originalWorkTime.Seconds(), newWorkTime.Seconds())
}
