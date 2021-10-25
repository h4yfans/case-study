package domain

import (
	"context"

	"github.com/h4yfans/case-study/models"
)

type UserRepository interface {
	Create(ctx context.Context, user *models.User) (*models.User, error)
	Update(ctx context.Context, user *models.User) (*models.User, error)
	Delete(ctx context.Context, id int) error
	GetByID(ctx context.Context, id int) (*models.User, error)
	GetAllUser(ctx context.Context) (models.UserSlice, error)
}

type UserUsecase interface {
	Create(ctx context.Context, user *models.User) (*UserResponse, error)
	Update(ctx context.Context, user *models.User) (*UserResponse, error)
	Delete(ctx context.Context, id int) error
	GetByID(ctx context.Context, id int) (*UserResponse, error)
	GetAllUser(ctx context.Context) ([]UserResponse, error)
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
