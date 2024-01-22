package repository

import (
	"context"
	"tigerhall-kittens/cmd/models"

	"tigerhall-kittens/cmd/constants"

	"gorm.io/gorm"
)

type UserRepository interface {
	GetUserByUserId(ctx context.Context, UserId string) (_ *models.User, err error)
	GetUserByEmail(ctx context.Context, email string) (_ *models.User, err error)
	GetAllUsers(ctx context.Context) (_ []models.User, err error)
	CreateNewUser(ctx context.Context, user models.User) (_ models.User, err error)
	Save(ctx context.Context, user *models.User) (_ *models.User, err error)
	DeleteUserById(ctx context.Context, userId string) (err error)
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{
		db: db,
	}
}

func (uRepo *userRepository) GetUserByEmail(ctx context.Context, email string) (_ *models.User, err error) {
	var user models.User
	var dbWithCtx = uRepo.db.WithContext(ctx)
	getUser := dbWithCtx.Where(constants.WhereEmail, email).Take(&user)
	return &user, getUser.Error
}

func (uRepo *userRepository) GetUserByUserId(ctx context.Context, UserId string) (_ *models.User, err error) {
	var user models.User
	var dbWithCtx = uRepo.db.WithContext(ctx)
	getUser := dbWithCtx.Where(constants.WhereUserId, UserId).Take(&user)
	return &user, getUser.Error
}

func (uRepo *userRepository) CreateNewUser(ctx context.Context, user models.User) (_ models.User, err error) {
	var dbWithCtx = uRepo.db.WithContext(ctx)
	getUser := dbWithCtx.Create(&user)
	return user, getUser.Error
}

func (uRepo *userRepository) GetAllUsers(ctx context.Context) (_ []models.User, err error) {
	var users []models.User
	var dbWithCtx = uRepo.db.WithContext(ctx)
	getUser := dbWithCtx.Find(&users)
	return users, getUser.Error
}
func (uRepo *userRepository) Save(ctx context.Context, user *models.User) (_ *models.User, err error) {
	var dbWithCtx = uRepo.db.WithContext(ctx)
	getUser := dbWithCtx.Where(constants.WhereUserId, user.UserId).Updates(user)
	return user, getUser.Error
}
func (uRepo *userRepository) DeleteUserById(ctx context.Context, userId string) (err error) {
	var user models.User
	var dbWithCtx = uRepo.db.WithContext(ctx)
	getUser := dbWithCtx.Where(constants.WhereUserId, userId).Delete(&user)
	return getUser.Error
}
