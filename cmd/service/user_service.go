package service

import (
	"context"
	"errors"
	"fmt"
	"tigerhall-kittens/cmd/constants"
	"tigerhall-kittens/cmd/models"
	"tigerhall-kittens/cmd/repository"
	"tigerhall-kittens/config"
	"time"

	"github.com/golang-jwt/jwt"

	"github.com/mitchellh/mapstructure"
	"golang.org/x/crypto/bcrypt"
)

type UserService interface {
	GetUserByUserId(ctx context.Context, UserId string) (_ *models.User, err error)
	GetUserByEmail(ctx context.Context, email string) (_ *models.User, err error)
	GetAllUsers(ctx context.Context) (_ []models.User, err error)
	CreateNewUser(ctx context.Context, user models.User) (_ models.User, err error)
	UpdateUser(ctx context.Context, user models.User) (_ *models.User, err error)
	DeleteUser(ctx context.Context, userId string) (err error)
	SignUp(ctx context.Context, user models.User) (_ models.User, err error)
	Login(ctx context.Context, user models.User) (_ models.User, tokenstr string, err error)
}

type userService struct {
	userRepository repository.UserRepository
}

func NewUserService(userRepo repository.UserRepository) UserService {
	return &userService{
		userRepository: userRepo,
	}
}

func (service *userService) GetUserByEmail(ctx context.Context, email string) (_ *models.User, err error) {
	entity, errorDb := service.userRepository.GetUserByEmail(ctx, email)
	if errorDb != nil {
		return nil, errorDb
	}
	return entity, nil
}

func (service *userService) Login(ctx context.Context, user models.User) (_ models.User, tokenstr string, err error) {
	// Look up the requested user
	result, err := service.userRepository.GetUserByEmail(ctx, user.Email)
	if err != nil {
		return models.User{}, "", err
	}

	if result == nil {
		return models.User{}, "", errors.New("the user does not exist")
	}

	err = bcrypt.CompareHashAndPassword([]byte(result.Password), []byte(user.Password))
	if err != nil {
		return models.User{}, "", errors.New("the password you entered is wrong")
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": user.UserId,
		"exp": time.Now().Add(time.Second * constants.TokenExpiryTime).Unix(),
	})

	// Sign and get the complete encoded token as a string using the secret
	tokenString, err := token.SignedString([]byte(config.GetEnv("JWT_SECRET", "jwt_default,secret")))
	if err != nil {
		return models.User{}, "", errors.New("failed to generate JWT token")
	}

	return user, tokenString, nil
}

func (service *userService) SignUp(ctx context.Context, user models.User) (_ models.User, err error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return user, fmt.Errorf("failed to hash password: %v", err)
	}

	fmt.Println("hash ", hash)

	user.Password = string(hash)

	entity, errorDb := service.userRepository.CreateNewUser(ctx, user)
	if errorDb != nil {
		return user, errorDb
	}
	return entity, nil
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
	errorDb := service.userRepository.DeleteUserById(ctx, userId)
	if errorDb != nil {
		return errorDb
	}
	return nil
}
