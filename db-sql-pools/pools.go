package pools

import (
	"database/sql"
	"fmt"
	"os"
)

// Setup configures the db along with pools
// number of connections and more
func Setup() (*sql.DB, error) {
	user := os.Getenv("MYSQLUSERNAME")
	pass := os.Getenv("MYSQLPASSWORD")
	dataSource := fmt.Sprintf("%s:%s@/gocookbook?parseTime=true", user, pass)
	db, err := sql.Open("mysql", dataSource)
	if err != nil {
		return nil, err
	}

	// there will only ever be 24 open connections
	db.SetMaxOpenConns(24)

	// MaxIdleConns can never be less than max open SetMaxOpenConns
	// otherwise it'll default to that value
	db.SetMaxIdleConns(24)

	return db, nil
}
