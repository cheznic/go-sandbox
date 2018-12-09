package main

import "github.com/cheznic/go-sandbox/db-persistence"

func main() {
	if err := storage.Exec(); err != nil {
		panic(err)
	}
}
