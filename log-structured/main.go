package main

import (
	"fmt"

	"github.com/cheznic/go-sandbox/log-structured/structured"
)

func main() {
	fmt.Println("Logrus:")
	structured.Logrus()

	fmt.Println()
	fmt.Println("Apex:")
	structured.Apex()
}
