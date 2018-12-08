package main

import (
	"github.com/cheznic/go-sandbox/db-sql-pools"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	if err := pools.ExecWithTimeout(); err != nil {
		panic(err)
	}
}
