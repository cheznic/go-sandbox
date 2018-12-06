package main

import "fmt"

var a int
var b string
var c bool

func main() {
	x := 42
	y := "James Bond"
	z := true

	fmt.Printf("Value of x: '%v', Value of y: '%v', Value of z: '%v'\n", x, y, z)

	fmt.Println("Value of x:", x)
	fmt.Println("Value of y:", y)
	fmt.Println("Value of z:", z)

	fmt.Println("Default value for a:", a)
	fmt.Println("Default value for b:", b)
	fmt.Println("Default value for c:", c)

	a = 42
	b = "James Bond"
	c = true

	line := fmt.Sprintf("'%v', '%v', '%v'", a, b, c)

	fmt.Println("New values for a, b, and c are:", line)

	fmt.Println("Zero value of d:", d)
	fmt.Printf("Type for d: %T\n", d)
	d = 42
	fmt.Println("Revised value of d:", d)

	fmt.Println("Zero value of q:", q)
	q = 96
	fmt.Println("Updated value of q:", q)
	p = int(q)
	fmt.Printf("The underlying type of q: %T\n", p)
}

type hotdog int

var d hotdog
var q hotdog
var p int
