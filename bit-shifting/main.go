package main

import (
	"fmt"
)

const (
	_  = iota
	kb = 1 << (iota * 10)
	mb = 1 << (iota * 10)
	gb = 1 << (iota * 10)
)

func main() {

	fmt.Printf("Decimal %16d\t Binary: %32b\n", kb, kb)
	fmt.Printf("Decimal %16d\t Binary: %32b\n", mb, mb)
	fmt.Printf("Decimal %16d\t Binary: %32b\n", gb, gb)

	v1 := 50
	v2 := v1 << 1
	v3 := v1 >> 1
	fmt.Printf("Decimal %4d\t Binary: %8b\t Hex: %#5x\n", v3, v3, v3)
	fmt.Printf("Decimal %4d\t Binary: %8b\t Hex: %#5x\n", v1, v1, v1)
	fmt.Printf("Decimal %4d\t Binary: %8b\t Hex: %#5x\n", v2, v2, v2)

}
