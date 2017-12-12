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
}
var DBCon DB

type ModelExtend interface {
	TableName() string
	GetSchema() []interface{}
}

func (db *DB) Find(m ModelExtend) (err error) {
	stmt := "SELECT * FROM " + m.TableName()
	if db.Condition != "" {
		stmt += " " +  db.Condition
	}
	
	rows, err := db.Db.Query(stmt)
	if err != nil {
		return
	}
	defer rows.Close()
	cols, err := rows.Columns()
	err = ScanInto(rows, cols, m)
	return
}

func ScanInto(rows *sql.Rows, colsName []string, m ModelExtend) (err error) {
	for rows.Next() {
			err = rows.Scan(m.GetSchema()...); if err != nil {
				return
			}
	}
	return
}

func (db *DB) Where(stmt string, cond ...interface{}) (*DB) {
	for _, v := range cond {
		stmt = strings.Replace(stmt, "?", fmt.Sprintf("%v", v), 1)
	}
	db.Condition = "WHERE " + stmt
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
