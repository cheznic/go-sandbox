package main

import (
	"github.com/cheznic/go-sandbox/db-sql"
	"github.com/cheznic/go-sandbox/db-sql-txn"
	_ "github.com/go-sql-driver/mysql" //we import supported libraries for database/sql
)

func main() {
	db, err := database.Setup()
	if err != nil {
		panic(err)
	}

	tx, err := db.Begin()
	if err != nil {
		panic(err)
	}
	// this wont do anything if commit is successful
	defer tx.Rollback()

	if err := dbinterface.Exec(db); err != nil {
		panic(err)
	}
	if err := tx.Commit(); err != nil {
		panic(err)
	}
}
