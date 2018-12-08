package database

import (
	"database/sql"
)

// Create makes a table called example
// and populates it
func Create(db *sql.DB) error {
	// create the database
	query := `CREATE TABLE example (name VARCHAR(20), created DATETIME)`
	if _, err := db.Exec(query); err != nil {
		return err
	}

	query = `INSERT INTO example (name, created) values ("Aaron", NOW())`
	if _, err := db.Exec(query); err != nil {
		return err
	}

	return nil
}
