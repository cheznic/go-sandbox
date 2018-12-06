package main

import (
	"context"
	"fmt"
	"runtime"
	"time"
)

func main() {

	ctx, cancel := context.WithCancel(context.Background())

	fmt.Println("first error check:", ctx.Err())
	fmt.Println("Goroutines:", runtime.NumGoroutine())

	f := func(s string) {
		n := 0
		for {
			select {
			case <-ctx.Done():
				return
			default:
				n++
				time.Sleep(time.Millisecond * 200)
				fmt.Println("\tworking:", s+fmt.Sprintf("%v", n))
			}
		}
	}

	go f("A")
	go f("B")

	time.Sleep(time.Second * 2)
	fmt.Println("second error check:", ctx.Err())
	fmt.Println("Goroutines:", runtime.NumGoroutine())

	fmt.Println("canceling")
	cancel()

	time.Sleep(time.Second * 2)
	fmt.Println("third error check:", ctx.Err())
	fmt.Println("Goroutines:", runtime.NumGoroutine())
}
