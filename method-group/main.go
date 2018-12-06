package main

import (
	"fmt"
	"strings"
)

type cow struct {
}

type bird struct {
}

type frog struct {
}

type animal interface {
	speak() string
	move() string
}

func (c cow) speak() string {
	return "moo"
}

func (c cow) move() string {
	return "walk"
}

func (b *bird) speak() string {
	return "tweet"
}

func (b *bird) move() string {
	return "fly"
}

func (f frog) speak() string {
	return "ribbit"
}

func (f frog) move() string {
	return "hop"
}

func animate(a animal) {

	s := fmt.Sprintf("%T", a)
	if i := strings.IndexByte(s, '.'); i >= 0 {
		s = s[i+1:]
	}
	fmt.Printf("A %v says %v and can %v\n", s, a.speak(), a.move())
}

func main() {

	// You can use an interface to pass a pointer to a receiver expecting a value
	c := cow{}
	animate(&c)

	// You can use an interface to pass a pointer to a receiver expecting a pointer
	b := bird{}
	animate(&b)

	// You can use an interface to pass a value to a receiver expecting a value
	f := frog{}
	animate(f)

	// You can NOT use an interface to pass a value to a receiver expecting a pointer
	// animate(b)
}
