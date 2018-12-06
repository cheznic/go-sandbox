package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
	"time"
)

func main() {
	fmt.Println("CPUs:", runtime.NumCPU())
	fmt.Println("Go routines @ start:", runtime.NumGoroutine())

	counter := int64(0)

	const cnt = 10
	var wg sync.WaitGroup

	for i := 0; i < cnt; i++ {
		wg.Add(1)
		go func() {
			time.Sleep(time.Second)
			atomic.AddInt64(&counter, 1)
			runtime.Gosched()
			fmt.Println("Counter:", atomic.LoadInt64(&counter))
			wg.Done()
		}()
	}
	fmt.Println("Go routines while waiting:", runtime.NumGoroutine())
	wg.Wait()
	fmt.Println("Go routines @ end:", runtime.NumGoroutine())
	fmt.Println("Counter:", counter)
}

