package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

type DB struct {
	*sql.DB
}

var dbinfo string
var DBCon *DB

func init() {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// host := os.Getenv("DB_HOST")
	user := os.Getenv("DB_USERNAME")
	password := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_DATABASE")
	dbinfo = fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable", user, password, dbName)
}

func DBConnection() (*DB, error) {
	Db, err := sql.Open("postgres", dbinfo)
	if err != nil {
		return nil, err
	}

	return &DB{Db}, nil
}
