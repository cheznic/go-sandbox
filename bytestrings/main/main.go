package main

import (
	"fmt"

	"github.com/cheznic/go-sandbox/bytestrings"
)

func main() {
	err := bytestrings.WorkingWithBuffer()
	if err != nil {
		panic(err)
	}

	fmt.Println()

	bytestrings.SearchString()
	bytestrings.ModifyString()
	bytestrings.StringReader()
}
