package main

import (
	"fmt"

	"github.com/cheznic/go-sandbox/error-log/log"
)

func main() {
	fmt.Println("basic logging and modification of logger:")
	log.Log()
	fmt.Println("logging 'handled' errors:")
	log.FinalDestination()
}
