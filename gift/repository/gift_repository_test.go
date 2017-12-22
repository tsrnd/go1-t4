package repository

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/golang/mock/gomock"

	"gopkg.in/DATA-DOG/go-sqlmock.v1"

	psql "github.com/goweb4/gift/repository/psql"
)

func TestGetByFromUserID(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	db, mockSQL, err := sqlmock.New()
	if err != nil {
		fmt.Println("err creating mock database")
		return
	}

	columns := []string{"id", "from_user_id", "to_user_id", "product_id", "message"}
	mockSQL.
		ExpectQuery("SELECT id, from_user_id, to_user_id, product_id, message FROM gifts WHERE from_user_id = (.+)").
		WithArgs(1).WillReturnRows(sqlmock.NewRows(columns).AddRow(1, 1, 1, 2, "test"))

	repo := psql.NewGiftRepository(db)
	_, err = repo.GetByFromUserID(1)

	assert.NoError(t, err)
}
