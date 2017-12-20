package psql

import (
	// _ "database/sql"
	// _ "database/sql/driver"
	"fmt"

	// _ "github.com/lib/pq"

	"testing"

	"github.com/stretchr/testify/assert"
	mockDb "gopkg.in/DATA-DOG/go-sqlmock.v1"
)

func TestBirdCreateSuccess(t *testing.T) {
	db, mock, err := mockDb.New()
	if err != nil {
		fmt.Println("error creating mock database")
		return
	}
	defer db.Close()

	columns := []string{"id"}

	mock.ExpectQuery(`insert into birds`).
		WithArgs("name", "color", "des").
		WillReturnRows(mockDb.NewRows(columns).AddRow(1))
	id, err := NewBirdRepository(db).Create("name", "color", "des")

	if err != nil {
		t.Error(err)
	}
	assert.Equal(t, id, int64(1))
	assert.Equal(t, err, nil)
}
