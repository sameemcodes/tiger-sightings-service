package service

import (
	"context"
	"fmt"
	"tigerhall-kittens/cmd/models"
	"tigerhall-kittens/cmd/repository"

	"github.com/mitchellh/mapstructure"
)

type UserService interface {
	GetUserByUserId(ctx context.Context, UserId string) (_ *models.User, err error)
	GetAllUsers(ctx context.Context) (_ []models.User, err error)
	CreateNewUser(ctx context.Context, user models.User) (_ models.User, err error)
	UpdateUser(ctx context.Context, user models.User) (_ *models.User, err error)
	DeleteUser(ctx context.Context, userId string) (err error)
}

// user service implementation
type userService struct {
	userRepository repository.UserRepository
}

// NewUserService creates a new instance of UserService
func NewUserService(userRepo repository.UserRepository) UserService {
	return &userService{
		userRepository: userRepo,
	}
}

func (service *userService) GetUserByUserId(ctx context.Context, userId string) (_ *models.User, err error) {
	entity, errorDb := service.userRepository.GetUserByUserId(ctx, userId)
	if errorDb != nil {
		return nil, errorDb
	}
	return entity, nil
}

func (service *userService) GetAllUsers(ctx context.Context) (_ []models.User, err error) {
	entity, errorDb := service.userRepository.GetAllUsers(ctx)
	if errorDb != nil {
		return nil, errorDb
	}
	return entity, nil
}

func (service *userService) CreateNewUser(ctx context.Context, user models.User) (_ models.User, err error) {
	entity, errorDb := service.userRepository.CreateNewUser(ctx, user)
	if errorDb != nil {
		return user, errorDb
	}
	return entity, nil
}

func (service *userService) UpdateUser(ctx context.Context, user models.User) (_ *models.User, err error) {
	var userId = user.UserId
	newUser, err := service.GetUserByUserId(ctx, userId)
	fmt.Println("newUser  ", newUser, "err ", err)
	// Bind the JSON request body to the user object
	fmt.Println("user ", user)
	err2 := mapstructure.Decode(user, &newUser)
	fmt.Println("newUser", newUser, "user ", user)
	if err2 != nil {
		fmt.Println("err2  ", err2)
		return &user, err2
	}
	entity, errorDb := service.userRepository.Save(ctx, newUser)
	if errorDb != nil {
		return &user, errorDb
	}
	return entity, nil
}
func (service *userService) DeleteUser(ctx context.Context, userId string) (err error) {
	//Delete User by Id
	errorDb := service.userRepository.DeleteUserById(ctx, userId)
	if errorDb != nil {
		return errorDb
	}
	return nil
}
