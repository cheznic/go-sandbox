package numath

import (
	"fmt"
)

func ExampleSum() {
	fmt.Println(Sum(1, 2, 3, 4, 5))
	// Output:
	// 15
}

// test from this directory using:
//    go test
// test from  parent or root directory using:
// go test ./...
