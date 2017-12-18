package sql

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

// Connect func
func Connect(dbDlct, dbUser, dbPass, dbHost, dbPort, dbName string) (*sql.DB, error) {
	connStr := fmt.Sprintf("user=%s password=%s dbname=%s host=%s port=%s sslmode=disable",
		dbUser, dbPass, dbName, dbHost, dbPort)
	return sql.Open(dbDlct, connStr)
}
