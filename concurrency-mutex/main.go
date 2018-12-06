package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	fmt.Println("CPUs:", runtime.NumCPU())
	fmt.Println("Go routines @ start:", runtime.NumGoroutine())

	counter := 0

	const gs = 100
	var wg sync.WaitGroup
	wg.Add(gs)

	// initialize Mutex here
	var mu sync.Mutex

	for i := 0; i < gs; i++ {
		go func() {
			// set Mutex lock here
			mu.Lock()
			v := counter
			runtime.Gosched()
			v++
			counter = v

			// release mutex lock here
			mu.Unlock()
			wg.Done()
		}()
		fmt.Println("Go routines in for{}:", runtime.NumGoroutine())
	}
	wg.Wait()
	fmt.Println("Go routines @ end:", runtime.NumGoroutine())
	fmt.Println("Counter:", counter)
}
