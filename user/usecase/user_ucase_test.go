package usecase

import (
	"context"
	"testing"

	mocks "github.com/h4yfans/case-study/mocks/domain"
	"github.com/h4yfans/case-study/models"
	"github.com/stretchr/testify/assert"
)

func TestCreate(t *testing.T) {
	mockRepo := new(mocks.UserRepository)
	user := &models.User{
		Name:     "Kaan",
		Email:    "kaan@test.com",
		Password: "123123",
	}

	userData := user
	userData.ID = 1

	mockRepo.On("Create", context.Background(), user).Return(userData, nil)
	u := NewUserUsecase(mockRepo)
	a, err := u.Create(context.Background(), user)
	assert.NoError(t, err)
	assert.NotNil(t, a)
	mockRepo.AssertExpectations(t)
}

func TestUpdate(t *testing.T) {
	mockRepo := new(mocks.UserRepository)
	user := &models.User{
		Name:     "Kaan",
		Password: "123123",
	}

	userData := user
	userData.ID = 1

	mockRepo.On("Update", context.Background(), user).Return(userData, nil)
	u := NewUserUsecase(mockRepo)
	a, err := u.Update(context.Background(), user)
	assert.NoError(t, err)
	assert.NotNil(t, a)
	mockRepo.AssertExpectations(t)
}

func TestDelete(t *testing.T) {
	mockRepo := new(mocks.UserRepository)

	mockRepo.On("Delete", context.Background(), 1).Return(nil, nil)
	u := NewUserUsecase(mockRepo)
	err := u.Delete(context.Background(), 1)
	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestGetByID(t *testing.T) {
	mockRepo := new(mocks.UserRepository)
	user := &models.User{
		ID:       1,
		Name:     "Kaan",
		Password: "123123",
	}

	mockRepo.On("GetByID", context.Background(), 1).Return(user, nil)
	u := NewUserUsecase(mockRepo)
	a, err := u.GetByID(context.Background(), 1)
	assert.NoError(t, err)
	assert.NotNil(t, a)
	mockRepo.AssertExpectations(t)
}

func TestGetAllUser(t *testing.T) {
	mockRepo := new(mocks.UserRepository)
	userData := models.UserSlice{
		{
			ID:       1,
			Name:     "Kaan",
			Password: "123123",
		},
	}

	mockRepo.On("GetAllUser", context.Background()).Return(userData, nil)
	u := NewUserUsecase(mockRepo)
	a, err := u.GetAllUser(context.Background())
	assert.NoError(t, err)
	assert.NotNil(t, a)
	mockRepo.AssertExpectations(t)
}
