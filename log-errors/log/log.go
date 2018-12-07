package log

import (
	"bytes"
	"fmt"
	"log"
)

// Log uses the setup logger
func Log() {
	// we'll configure the logger to write
	// to a bytes.Buffer
	buf := bytes.Buffer{}

	// second argument is the prefix last argument is about
	// options you combine them with a logical or.
	logger := log.New(&buf, "log: ", log.Lshortfile|log.Ldate|log.Ltime)

	logger.Println("test1")
	logger.Println("test2")
	logger.Println("test3")
	logger.Println("test4")

	logger.SetPrefix("new log prefix: ")

	logger.Println("test5")
	logger.Println("test6")
	logger.Println("test7")
	logger.Println("test8")

	logger.Printf("you can also add args(%v) and use Fataln to log and crash", true)

	fmt.Println(buf.String())
}
