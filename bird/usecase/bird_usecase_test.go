package usecase

import (
	"testing"
	"github.com/golang/mock/gomock"
	"github.com/goweb4/bird/usecase/mocks"
	"github.com/stretchr/testify/assert"
	models "github.com/goweb4/bird"
)

func TestUsecaseCreateBird(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockUc := mocks.NewMockBirdUsecase(mockCtrl)
	var dataExpect *models.Bird

	mockUc.EXPECT().Create("My Bird", "Yellow", "Yellow bird").Return(dataExpect, nil).Times(1)
	rs, err := mockUc.Create("My Bird", "Yellow", "Yellow bird")
	
	assert.Empty(t, err, nil)
	assert.Equal(t, rs, dataExpect)
}
