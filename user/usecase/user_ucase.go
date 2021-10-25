package usecase

import (
	"context"
	"net/mail"
	"strings"

	"github.com/h4yfans/case-study/common"
	"github.com/h4yfans/case-study/domain"
	"github.com/h4yfans/case-study/models"
	"golang.org/x/crypto/bcrypt"
)

type UserUsecase struct {
	repo domain.UserRepository
}

func NewUserUsecase(repo domain.UserRepository) *UserUsecase {
	return &UserUsecase{repo: repo}
}

func (u *UserUsecase) Create(ctx context.Context, user *models.User) (*domain.UserResponse, error) {
	ok := u.validate(user, true)
	if !ok {
		return nil, common.BadRequest
	}

	password, err := u.HashPassword(user.Password)
	if err != nil {
		return nil, common.BadRequest
	}

	//hashed password
	user.Password = password

	userData, err := u.repo.Create(ctx, user)
	if err != nil {
		return nil, err
	}

	serializer := domain.UserSerializer(userData)

	return serializer, nil
}

func (u *UserUsecase) Update(ctx context.Context, user *models.User) (*domain.UserResponse, error) {
	ok := u.validate(user, false)
	if !ok {
		return nil, common.BadRequest
	}

	password, err := u.HashPassword(user.Password)
	if err != nil {
		return nil, common.BadRequest
	}

	//hashed password
	user.Password = password

	userData, err := u.repo.Update(ctx, user)
	if err != nil {
		return nil, err
	}

	serializer := domain.UserSerializer(userData)

	return serializer, nil
}

func (u *UserUsecase) Delete(ctx context.Context, id int) error {
	err := u.repo.Delete(ctx, id)
	if err != nil {
		return err
	}

	return nil
}

func (u *UserUsecase) GetByID(ctx context.Context, id int) (*domain.UserResponse, error) {
	user, err := u.repo.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}

	serializer := domain.UserSerializer(user)

	return serializer, nil
}

func (u *UserUsecase) GetAllUser(ctx context.Context) ([]domain.UserResponse, error) {
	users, err := u.repo.GetAllUser(ctx)
	if err != nil {
		return nil, err
	}

	serializers := make([]domain.UserResponse, 0)
	for _, user := range users {
		serializers = append(serializers, *domain.UserSerializer(user))
	}

	return serializers, nil
}

func (u *UserUsecase) validate(user *models.User, create bool) bool {
	var err error
	if create {
		_, err = mail.ParseAddress(user.Email)
	}

	if strings.TrimSpace(user.Name) == "" || strings.TrimSpace(user.Password) == "" || err != nil {
		return false
	}
	return true
}

func (u *UserUsecase) HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}
