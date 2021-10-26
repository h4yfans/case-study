package domain

import (
	"context"

	"github.com/h4yfans/case-study/models"
)

type UserRepository interface {
	Create(c context.Context, user *models.User) (*models.User, error)
	Update(c context.Context, user *models.User) (*models.User, error)
	Delete(c context.Context, id int) error
	GetByID(c context.Context, id int) (*models.User, error)
	GetAllUser(c context.Context) (models.UserSlice, error)
}

type UserUsecase interface {
	Create(c context.Context, user *models.User) (*UserResponse, error)
	Update(c context.Context, user *models.User) (*UserResponse, error)
	Delete(c context.Context, id int) error
	GetByID(c context.Context, id int) (*UserResponse, error)
	GetAllUser(c context.Context) ([]UserResponse, error)
}

type UserResponse struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

func UserSerializer(user *models.User) *UserResponse {
	return &UserResponse{
		ID:    user.ID,
		Name:  user.Name,
		Email: user.Email,
	}
}
