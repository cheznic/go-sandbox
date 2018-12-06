package main

import (
	"fmt"
)

func main() {

	fmt.Println("start")

	eve := make(chan int)
	odd := make(chan int)
	quit := make(chan int)

	go send(eve, odd, quit)

	receive(eve, odd, quit)

	fmt.Println("exiting")
}

func receive(even, odd, quit <-chan int) {
	for {
		select {
		case v := <-even:
			fmt.Println("From the eve channel", v)
		case v := <-odd:
			fmt.Println("From the odd channel", v)
		case <-quit:
			return
		}
	}
}

func send(even, odd, quit chan<- int) {
	for i := 0; i < 50; i++ {
		if i%2 == 0 {
			even <- i
		} else {
			odd <- i
		}
	}
	close(quit)
}
