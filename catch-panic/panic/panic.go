package panic

import (
	"fmt"
	"strconv"
)

// Panic panics with a divide by zero
func Panic() {
	defer func() {
		fmt.Println("deferred just before divide by zero")
	}()

	zero, _ := strconv.ParseInt("0", 10, 64)

	// panic happens here
	a := 1 / zero
	// normal execution stops
	fmt.Println("we'll never get here", a)
}

// Catcher calls Panic
// Panic is caught since deferred functions are called in LIFO order after a panic
func Catcher() {
	// waits until end of function to run
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("panic occurred:", r)
		}
	}()

	Panic()
}
