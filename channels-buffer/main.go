package main

import (
	"fmt"
)

func main() {

	// Adding a buffer when creating channels will allow values to be sent
	// on the channel without blocking.  The number values that can be sent
	// before blocking occurs is equal tot he buffer.
	buff := 1
	c := make(chan int, buff)

	// One value was placed on the buffered channel without blocking.
	c <- 42

	// The number of buffered values will exceed the buffer size if the
	// next line is uncommented, and the app will deadlocked.
	// c <- 86

	fmt.Println(<-c)
}
