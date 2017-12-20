package psql

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"testing"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/stretchr/testify/assert"
)

func TestClassCreate(t *testing.T) {
	db := CreateDBConnection()
	CreateClassTable(db)
	defer DropClassTable(db)
	err := NewClassRepository(db).Create("class A", 20)
	assert.Equal(t, nil, err)
}

func CreateDBConnection() *sql.DB {
	godotenv.Load()
	dbDlct := os.Getenv("DATABASE_DLCT")
	dbUser := os.Getenv("DATABASE_USER")
	dbPass := os.Getenv("DATABASE_PASS")
	dbHost := os.Getenv("DATABASE_HOST")
	dbPort := os.Getenv("DATABASE_PORT")
	dbName := os.Getenv("DATABASE_NAME")

	connStr := fmt.Sprintf("user=%s password=%s dbname=%s host=%s port=%s sslmode=disable",
		dbUser, dbPass, dbName, dbHost, dbPort)
	db, _ := sql.Open(dbDlct, connStr)
	return db
}

func CreateClassTable(db *sql.DB) {
	_, err := db.Exec("CREATE TABLE class ( id serial, name varchar(32), student_number integer)")
	if err != nil {
		fmt.Println("er: ", err)
	}
}

func DropClassTable(db *sql.DB) {
	if _, err := db.Exec("DROP TABLE class"); err != nil {
		log.Print(err)
	}
}
