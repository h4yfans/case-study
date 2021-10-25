package repository

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/h4yfans/case-study/common"
	"github.com/h4yfans/case-study/domain"
	"github.com/h4yfans/case-study/models"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) domain.UserRepository {
	return &UserRepository{
		db: db,
	}
}

func (u *UserRepository) Create(ctx context.Context, user *models.User) (*models.User, error) {
	exists, err := u.getByEmail(ctx, user.Email)
	if err != nil {
		return nil, err
	}
	if exists {
		return nil, common.UserAlreadyExist
	}

	err = user.Insert(ctx, u.db, boil.Infer())
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (u *UserRepository) Update(ctx context.Context, user *models.User) (*models.User, error) {
	whitelist := []string{models.UserColumns.Name, models.UserColumns.Password}
	effected, err := user.Update(ctx, u.db, boil.Whitelist(whitelist...))
	if err != nil {
		return nil, err
	}

	if effected != 1 {
		err = fmt.Errorf("Weird  Behavior. Total Affected: %d", effected)
		return nil, err
	}

	userData, err := u.GetByID(ctx, user.ID)
	if err != nil {
		return nil, err
	}

	return userData, nil
}

func (u *UserRepository) Delete(ctx context.Context, id int) error {
	user := models.User{ID: id}
	effected, err := user.Delete(ctx, u.db)
	if err != nil {
		return err
	}

	if effected == 0 {
		return common.UserNotExist
	}

	return nil
}

func (u *UserRepository) GetByID(ctx context.Context, id int) (*models.User, error) {
	user, err := models.FindUser(ctx, u.db, id)
	if err != nil {
		return nil, common.UserNotExist
	}

	return user, nil
}

func (u *UserRepository) GetAllUser(ctx context.Context) (models.UserSlice, error) {
	users, err := models.Users().All(ctx, u.db)
	if err != nil {
		return nil, err
	}

	return users, nil
}

func (u *UserRepository) getByEmail(ctx context.Context, email string) (bool, error) {
	exists, err := models.Users(models.UserWhere.Email.EQ(email)).Exists(ctx, u.db)
	if err != nil {
		return exists, err
	}
	return exists, err
}
