package psql

import (
	"testing"

	_ "github.com/lib/pq"
	sqlmock "gopkg.in/DATA-DOG/go-sqlmock.v1"
)

func TestClassCreate(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()
	rows := sqlmock.NewRows([]string{"id", "name", "student_number"}).AddRow(1, "class A", 20)
	mock.ExpectQuery("INSERT into class").WillReturnRows(rows)
	a := NewClassRepository(db)
	err = a.Create("class A", 20)
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}
