package main

import (
	"fmt"

	"github.com/cheznic/go-sandbox/error-wrap/errwrap"
)

func main() {
	errwrap.Wrap()
	fmt.Println()
	fmt.Println("Unwrap:")
	errwrap.Unwrap()
	fmt.Println()
	fmt.Println("Stack trace:")
	errwrap.StackTrace()
}
