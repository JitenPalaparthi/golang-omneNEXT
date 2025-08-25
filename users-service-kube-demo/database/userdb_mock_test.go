package database_test

import (
	"testing"
	"users-service/internal/mocks"
	"users-service/models"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestCreateMockUser(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockUserdb := mocks.NewMockIUserDB(ctrl) // creates an instance of that interface
	out := &models.User{Name: "Jiten", CommonModel: models.CommonModel{ID: 1}}
	var err error
	act := &models.User{Name: "Jiten", CommonModel: models.CommonModel{ID: 1}}
	mockUserdb.EXPECT().Create(gomock.AssignableToTypeOf(&models.User{})).Return(out, err).Times(1)
	got, err := mockUserdb.Create(act) // <-- actually call it
	require.NoError(t, err)
	assert.Equal(t, out, got)
	//mocks.NewMockIUserHandler(ctrl)
}

func TestGetByMockUser(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockUserdb := mocks.NewMockIUserDB(ctrl) // creates an instance of that interface
	out := &models.User{Name: "Jiten", CommonModel: models.CommonModel{ID: 1}}
	var err error
	mockUserdb.EXPECT().GetBy(1).Return(out, err).Times(1)
	got, err := mockUserdb.GetBy(1) // <-- actually call it
	require.NoError(t, err)
	assert.Equal(t, out, got)
	//mocks.NewMockIUserHandler(ctrl)
}
