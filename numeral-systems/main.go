package main

import (
	"fmt"
)

func main() {

	for i := 0; i < 40; i++ {
		fmt.Printf("Default: %3v, Binary: %8b, Octal: %3o, Decimal: %3d, Hex: %2X\n", i, i, i, i, i)
	}

}
