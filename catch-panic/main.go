package main

import (
	"fmt"

	"github.com/cheznic/go-sandbox/catch-panic/panic"
)

func main() {
	fmt.Println("before panic")
	panic.Catcher()
	fmt.Println("after panic")
}
