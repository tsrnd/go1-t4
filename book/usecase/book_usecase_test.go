package usecase_test

import (
	"strconv"
	"testing"

	"github/monstar-lab/fr-circle-api/models"
	"github/monstar-lab/fr-circle-api/usecase"

	"github.com/bxcodec/faker"
	"github.com/goweb4/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestFetch(t *testing.T) {
	mockBookRepo := new(mocks.BookRepository)
	var mockBook models.Book
	err := faker.FakeData(&mockBook)
	assert.NoError(t, err)

	mockListBook := make([]*models.Book, 0)
	mockListBook = append(mockListBook, &mockBook)
	mockBookRepo.On("Fetch", mock.AnythingOfType("string"), mock.AnythingOfType("int64")).Return(mockListBook, nil)
	uc := usecase.NewBookUsecase(mockBookRepo)
	limit := int64(1)
	offset := "12"
	list, err := uc.Fetch(offset, limit)
	offsetExpected := strconv.Itoa(int(mockBook.ID))
	assert.NoError(t, err)
	assert.Len(t, list, len(mockListBook))

	mockBookRepo.AssertCalled(t, "Fetch", mock.AnythingOfType("string"), mock.AnythingOfType("int64"))
}
