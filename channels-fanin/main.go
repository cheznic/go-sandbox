package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {

	fmt.Println("start")
	combined := make(<-chan int)

	combined = combine(splitEven(50), splitOdd(50))

	for v := range combined {
		fmt.Println(v)
	}

	fmt.Println("exiting")
}

// return a channel of only odd numbers
func splitOdd(x int) <-chan int {
	c := make(chan int)

	go func() {
		for i := 0; i < x; i++ {
			if i%2 == 0 {
				c <- i
			}
			time.Sleep(time.Second / 4)
		}
		close(c)
	}()

	return c
}

// return a channel of only even numbers
func splitEven(x int) <-chan int {
	c := make(chan int)

	go func() {
		for i := 0; i < x; i++ {
			if i%2 != 0 {
				c <- i
			}
			time.Sleep(time.Second / 4)
		}
		close(c)
	}()

	return c
}

// combine the inputs of channels 1 and 2 into single channel returned
func combine(input1, input2 <-chan int) <-chan int {
	c := make(chan int)

	var wg sync.WaitGroup
	wg.Add(2)

	// Direct input channel 1 to output channel.
	go func() {
		for v := range input1 {
			c <- v
		}
		wg.Done()
	}()

	// Direct input channel 2 to output channel.
	go func() {
		for v := range input2 {
			c <- v
		}
		wg.Done()
	}()

	// Wait to close channel but allow execution flow to continue.
	go func() {
		wg.Wait()
		close(c)
	}()

	return c
}
