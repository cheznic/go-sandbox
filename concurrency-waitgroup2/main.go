package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wg sync.WaitGroup

func main() {
	fmt.Println("Goroutines at the start:\t", runtime.NumGoroutine())

	// The x in wg.Add(x) is the number of wg.Done() signals the wait
	// group will wait for.  Calling wg.Done() signals to the wait group
	// that it's waiting for one less wg.Done() calls.  The wg.Done() 
	// calls are typically placed in separate goroutines which allows 
	// the calling context to wait for each wg.Done() call before exiting.
	wg.Add(2)

	// Calling wg.Done() here decrements x in wb.Add(x).  This is not
	// typical, just illustrative that wg.Done() can be called outside 
	// a new goroutine.
	wg.Done()

	go func() {
		fmt.Println("First function literal (will not print)")
	}()

	go func() {
		fmt.Println("Second function literal (will print)")

		// wg.Done() signals completion of goroutine
		// The number of wait groups added should match the number of
		// go routines that signal 'Done'.
		wg.Done()
	}()

	fmt.Println("Goroutines while waiting:\t", runtime.NumGoroutine())

	// If the value of x in wg.Add(x) exceeds the total number of
	// wg.Done() calls the app will error.  If the number of wg.Done()
	// calls exceed the value of x in wg.Add(x) the app might exit
	// before completion of all goroutines.
	wg.Wait()

	fmt.Println("Goroutines at the end:\t", runtime.NumGoroutine())
	fmt.Println()
}
