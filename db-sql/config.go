package database

import (
	"database/sql"
	"fmt"
	"os"
	"time"
)

// Example hold the results of our queries
type Example struct {
	Name    string
	Created *time.Time
}

// Setup configures and returns our database
// connection poold
func Setup() (*sql.DB, error) {
	user := os.Getenv("MYSQLUSERNAME")
	fmt.Println("user:", user)
	pass := os.Getenv("MYSQLPASSWORD")
	fmt.Println("pass:", pass)
	dataSource := fmt.Sprintf("%s:%s@/gocookbook?parseTime=true", user, pass)
	db, err := sql.Open("mysql", dataSource)
	if err != nil {
		return nil, err
	}
	return db, nil
}
