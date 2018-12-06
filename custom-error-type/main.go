package main

import (
	"fmt"
)

func main() {
	e := CustomErr{
		info: "some info",
	}

	foo(e)
}

// CustomErr .
type CustomErr struct {
	info string
}

func (ce CustomErr) Error() string {
	return fmt.Sprintf("Error: %v", ce.info)
}

func foo(e error) {
	fmt.Println(e)
}
