package main

import (
	"fmt"
	"runtime"
	"sync"
	// "time"
)

func main() {
	fmt.Println("CPUs:", runtime.NumCPU())
	fmt.Println("Go routines @ start:", runtime.NumGoroutine())

	counter := 0

	const gs = 100
	var wg sync.WaitGroup
	wg.Add(gs)

	for i := 0; i < gs; i++ {
		go func() {
			v := counter
			// time.Sleep(time.Millisecond / 50)
			runtime.Gosched()
			v++
			counter = v
			wg.Done()
		}()
		fmt.Println("Go routines in for{}:", runtime.NumGoroutine())
	}
	wg.Wait()
	fmt.Println("Go routines @ end:", runtime.NumGoroutine())
	fmt.Println("Counter:", counter)
}
