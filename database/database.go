package database

import (
	"fmt"
	"os"
	"log"
	_ "github.com/lib/pq"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/joho/godotenv"
)

type DB struct{
	*gorm.DB
}

var dbinfo string
var DBCon *DB

func init() {
	err := godotenv.Load()

	if err != nil {
	  log.Fatal("Error loading .env file")
	}

	host := os.Getenv("DB_HOST")
	user := os.Getenv("DB_USERNAME")
	password := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_DATABASE")
	dbinfo = fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=disable", host, user, password, dbName)
}

func DBConnection() (*DB, error) {
	Db, err := gorm.Open("postgres", dbinfo)
	if err != nil {
		return nil, err
	}

	return &DB{Db}, nil
}
