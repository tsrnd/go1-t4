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

var Db *DB

func init() {
	err := godotenv.Load()

	if err != nil {
	  log.Fatal("Error loading .env file")
	}

	host := os.Getenv("DB_HOST")
	user := os.Getenv("DB_USERNAME")
	password := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_DATABASE")
	dbinfo := fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=disable", host, user, password, dbName)
	
	var connectionErr error
	Db, connectionErr = DBConnection(dbinfo)
	if connectionErr != nil {
		log.Fatal(connectionErr)
	}
}

func DBConnection(dataSource string) (*DB, error) {
	Db, err := gorm.Open("postgres", dataSource)
	if err != nil {
		return nil, err
	}
	defer Db.Close()
	return &DB{Db}, nil
}
