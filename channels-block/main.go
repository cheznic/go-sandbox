package main

import (
	"fmt"
)

func main() {

	c := make(chan int)

	// This code as written will deadlock.  If there is no goroutine waiting 
	// to remove the value from the channel the channel will block.  Removing
	// the following comments will unblock the execution.
	//
	// go func() {
		c <- 42
	// }()

	fmt.Println(<-c)
}
