package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	fmt.Println("Start")

	var results <-chan int

	results = fanOut(src())

	for v := range results {
		fmt.Println(v)
	}
	fmt.Println("Exiting")
}

// populate() returns a channel of integers, 0 to 99
func src() <-chan int {

	out := make(chan int)
	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		for i := 0; i < 100; i++ {
			out <- i
		}
		wg.Done()
	}()

	// Wait for the goroutine to complete but allow program flow
	// to continue
	go func() {
		wg.Wait()
		close(out)
	}()

	return out
}

// fanOut functions in a pipeline, accepting an outbound channel
// as input and returning an outbound channel as output.
func fanOut(in <-chan int) <-chan int {

	out := make(chan int)
	var wg sync.WaitGroup

	for v := range in {
		wg.Add(1)

		// fanout the workload to run concurrently
		go func(v2 int) {
			out <- timeConsumingWork(v2)
			wg.Done()
		}(v)
	}

	// Wait for all goroutines to complete then close the channel
	// and allow program flow to continue.
	go func() {
		wg.Wait()
		close(out)
	}()

	return out
}

// this simulates time consuming work
// Note that the application completes in about a second
func timeConsumingWork(x int) int {
	time.Sleep(time.Second)
	return x
}
