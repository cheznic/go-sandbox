package main

import (
	"fmt"
)

// https://en.wikipedia.org/wiki/ASCII

func main() {
	s := fmt.Sprintln("Hello playground")

	bs := []byte(s)

	fmt.Println(bs)
	fmt.Printf("%T \n", bs)

	for i := 0; i < len(s); i++ {
		fmt.Printf("UTF-8 symbols: %#U, Int value %v\n", s[i], s[i])
	}

	fmt.Println()

	for i, v := range s {
		fmt.Printf("Index: %v, Hex: %#x, Value: %v\n", i, v, v)
	}

	x := `this is the first line of a raw string literal
	second line
	third`

	fmt.Println(x)
}
