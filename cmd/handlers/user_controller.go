package controller

import (
	"fmt"
	"net/http"
	"tigerhall-kittens/cmd/models"
	"tigerhall-kittens/cmd/service"

	"github.com/gin-gonic/gin"
)

type UserController interface {
	GetUserByUserId(ctx *gin.Context)
	GetUserByEmail(ctx *gin.Context)
	GetAllUsers(ctx *gin.Context)
	CreateNewUser(ctx *gin.Context)
	UpdateUser(ctx *gin.Context)
	DeleteUserById(ctx *gin.Context)
	SignUp(ctx *gin.Context)
	Login(ctx *gin.Context)
}

type userController struct {
	userService service.UserService
}

func NewUserController(userService service.UserService) UserController {
	return &userController{
		userService: userService,
	}
}

// GetUserByEmail godoc
// @Summary Get a user by email
// @Tags User-Controller
// @Accept */*
// @Produce json
// @Param email path string true "email"
// @Router /user/v1/email/{email} [get]

func (c *userController) GetUserByEmail(ctx *gin.Context) {
	email := ctx.Params.ByName("email")
	userDto, err := c.userService.GetUserByEmail(ctx, email)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{"Error during GetUserByEmail": err.Error()})
	} else {
		ctx.JSON(http.StatusOK, userDto)
	}
}

// GetLogin godoc
// @Summary Login a user
// @Tags User-Controller
// @Accept */*
// @Produce json
// @Param user body models.User true "User credentials for login"
// @Router /user/v1/login [post]
func (c *userController) Login(ctx *gin.Context) {
	var user models.User
	err := ctx.BindJSON(&user)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{"Invalid Input ": err.Error()})
	}
	userDto, tokenstr, err := c.userService.Login(ctx, user)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{"Error during Login ": err.Error()})
	} else {
		// Send the token as a cookie
		ctx.SetSameSite(http.SameSiteLaxMode)
		// have set to 10 seconds of cookie expiry
		ctx.SetCookie("Authorization", tokenstr, 10, "", "", false, true)
		ctx.JSON(http.StatusOK, gin.H{"user": userDto, "token": tokenstr})

	}

}

// PostSignUp godoc
// @Summary Sign up a user
// @Tags User-Controller
// @Accept */*
// @Produce json
// @Param user body models.User true "User credentials for Signup"
// @Router /user/v1/signup [post]
func (c *userController) SignUp(ctx *gin.Context) {
	var user models.User
	err := ctx.BindJSON(&user)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{"Invalid Input ": err.Error()})
	}
	fmt.Println("user", user)
	userDto, err := c.userService.SignUp(ctx, user)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{"Error during CreateNewUser ": err.Error()})
	} else {
		ctx.JSON(http.StatusOK, userDto)
	}

}

// GetUserByUserId godoc
// @Summary Get a user by userId
// @Tags User-Controller
// @Accept */*
// @Produce json
// @Param userId path string true "userId"
// @Router /user/v1/user_id/{userId} [get]
func (c *userController) GetUserByUserId(ctx *gin.Context) {

	userId := ctx.Params.ByName("userId")
	userDto, err := c.userService.GetUserByUserId(ctx, userId)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{"Error during GetUserByUserId": err.Error()})
	} else {
		ctx.JSON(http.StatusOK, userDto)
	}
}

// GetAllUsers godoc
// @Summary Get all users
// @Tags User-Controller
// @Accept */*
// @Produce json
// @Router /user/v1/fetch_all [get]
func (c *userController) GetAllUsers(ctx *gin.Context) {
	userDto, err := c.userService.GetAllUsers(ctx)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{"Error during GetAllUsers": err.Error()})
	} else {
		ctx.JSON(http.StatusOK, userDto)
	}
}

// CreateNewUser godoc
// @Summary Create a new user
// @Tags User-Controller
// @Accept */*
// @Param user body models.User true "User details in JSON format"
// @Success 200
// @Failure 404
// @Failure 500
// @Produce json
// @Router /user/v1/create_new [post]
func (c *userController) CreateNewUser(ctx *gin.Context) {
	var user models.User
	err := ctx.BindJSON(&user)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{"Invalid Input ": err.Error()})
	}
	userDto, err := c.userService.CreateNewUser(ctx, user)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{"Error during CreateNewUser ": err.Error()})
	} else {
		ctx.JSON(http.StatusOK, userDto)
	}
}

// UpdateUser godoc
// @Summary Update a user
// @Tags User-Controller
// @Param userId path string true "userId"
// @Accept */*
// @Produce json
// @Router /user/v1/update [put]
func (c *userController) UpdateUser(ctx *gin.Context) {
	var user models.User
	err := ctx.BindJSON(&user)
	if err != nil {
		fmt.Println("Error during UpdateUser", err.Error())
	}
	userDto, err := c.userService.UpdateUser(ctx, user)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{"Error during  UpdateUser": err.Error()})
	} else {
		ctx.JSON(http.StatusOK, userDto)
	}
}

// DeleteUserById godoc
// @Summary Delete a user by userId
// @Tags User-Controller
// @Accept */*
// @Param userId path string true "userId"
// @Produce json
// @Success 200
// @Router /user/v1/deletebyUserId/{userId} [delete]
func (c *userController) DeleteUserById(ctx *gin.Context) {
	userId := ctx.Params.ByName("userId")
	err := c.userService.DeleteUser(ctx, userId)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{"Error during DeleteUserById": err.Error()})
	} else {
		ctx.JSON(http.StatusOK, gin.H{"message": "User deleted successfully"})
	}
}
