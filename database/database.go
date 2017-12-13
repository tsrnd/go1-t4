package database

import (
	"strings"
	"database/sql"
	"fmt"
	"log"
	"os"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

var dbinfo string

type DB struct {
	Db *sql.DB
  Condition string
	ColsSelect string
	ColsScan []string
}
var DBCon DB

type ModelExtend interface {
	TableName() string
	GetSchema() map[string]interface{}
}

func (db *DB) Find(m ModelExtend) (err error) {
	stmt := "SELECT * FROM " + m.TableName()
	var ColsScan []interface{}
	if db.ColsSelect != "" {
		stmt = strings.Replace(stmt, "*", db.ColsSelect, 1)
  }

	if db.Condition != "" {
		stmt += " " +  db.Condition
	}
	
	rows, err := db.Db.Query(stmt)
	if err != nil {
		return
	}
	defer rows.Close()
	defer db.ClearAttr()
	//set pointer to scan values
	if db.ColsSelect == "" {
		colsName, err := rows.Columns(); if err != nil {
			return err
		}
		ColsScan = MappingPointer(colsName, m.GetSchema())
	} else {
		ColsScan = MappingPointer(db.ColsScan, m.GetSchema())
	}
	err = ScanInto(rows, ColsScan)
	return
}

func (db *DB) Select(params ...string) (*DB) {
	for i, v := range params {
		if i < len(params) - 1 {
			db.ColsSelect += v + ", "
			continue
		}
		db.ColsSelect += v
	}
	db.ColsScan = params
	return db
}

//Get data sources from env
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

//Setup struct for connection to the database
func DBConnection() (*sql.DB, error) {
	Db, err := sql.Open("postgres", dbinfo)
	return Db, err
}
