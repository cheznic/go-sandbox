package main

import "github.com/cheznic/go-sandbox/db-mongodb"

func main() {
	if err := mongodb.Exec(); err != nil {
		panic(err)
	}
}
