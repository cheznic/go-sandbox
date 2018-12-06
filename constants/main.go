package main

import (
	"fmt"
)

const a int32 = 23
const b = 32

const (
	c string = "James Bond"
	d        = "Daffy Duck"
)

func main() {
	fmt.Println("Typed constant:", a)
	fmt.Println("Constant:", b)
	fmt.Println("Typed constant:", c)
	fmt.Println("Constant:", d)
}
