// Package main .
package main

import (
	"fmt"
	"sync"
	"time"
)

// The example below approximates the relationship between client and
// server processes.  It demonstrates the power and flexibility of
// channels in go.  Note the speed in which the process completes.
// Each of the 20,000 clients sends 5 requests to the server with a one
// second wait between each request.  The program runs in a little over
// four seconds, yet handling 100,000 total requests.
func main() {

	fmt.Println("start")

	// Initialize two channels of string.  One for sending requests from
	// the client to the server and one for responding from the server
	// to the client.  The channels are declared with out a specific
	// direction.  When passed as arguments to the server and client,
	// functions, the function arguments declare the direction which
	// data is allowed to flow.
	req := make(chan string)
	res := make(chan string)

	// Pass the request and response channels to the server function and
	// launch it as a goroutine.
	go Server(req, res)

	// Create a wait group to assure client goroutines complete before
	// exiting.
	var wg sync.WaitGroup

	// Launching numerous goroutines in this way uncovered a limitation.
	// When attempting to call 10,000 client functions as goroutines
	// the following error was raised:
	//      "race: limit on 8128 simultaneously
	//       alive goroutines is exceeded, dying"
	// This limit is enforced when using the -race flag to run your go
	// app.  Removing -race when running the application allows for
	// higher concurrent counts of goroutines.
	x := 20000
	// Fanout x client goroutines by passing the request and response
	// channels.  Each client goroutine is also passed the wait group
	// and a client id.
	for i := 0; i < x; i++ {
		wg.Add(1)
		go Client(req, res, &wg, i)
	}

	// Wait for x client goroutines to complete.
	wg.Wait()

	// Close the request channel and end the server process.
	close(req)

	// Wait for and print the final response from the server
	fmt.Println(<-res)

	fmt.Println("exiting")
}

// Server function is launched as a goroutine, and accepts the request
// and response channels so the client function can communicate with
// the it.
func Server(req <-chan string, res chan<- string) {
	start := time.Now()
	i := 0
	// Echo the request message to response.
	for msg := range req {
		res <- msg
		i++
	}

	t := time.Since(start)
	res <- fmt.Sprintf("\n%v total requests handled in %v", i, t)
}

// Client function is launched multiple times as a goroutines.  It
// produces requests and sends them to the server via the req channel.
// The client function then waits for the response from the server on
// the res channel.
func Client(req chan<- string, res <-chan string, wg *sync.WaitGroup, id int) {
	for i := 0; i < 5; i++ {
		if i != 0 {
			time.Sleep(time.Second)
		}
		req <- fmt.Sprintf("\rmessage %v from client %v", i+1, id)
		fmt.Printf(<-res)
	}

	wg.Done()
}
