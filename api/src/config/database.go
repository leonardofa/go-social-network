package config

import (
	"database/sql"
	"fmt"

	// Registers the MySQL driver with database/sql.
	_ "github.com/go-sql-driver/mysql"
)

// GetConnection establishes and returns a database connection.
func GetConnection() (*sql.DB, error) {
	// Open the database connection.
	dbConn, err := sql.Open("mysql", DBUrl + "?parseTime=true")
	if err != nil {
		return nil, fmt.Errorf("error opening database: %w", err)
	}

	// Test the connection.
	if err := dbConn.Ping(); err != nil {
		err := dbConn.Close()
		if err != nil {
			return nil, err
		}

		return nil, fmt.Errorf("error connecting to database: %w", err)
	}

	return dbConn, nil
}
