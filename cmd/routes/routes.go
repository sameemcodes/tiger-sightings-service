package router

import (
	durable "tigerhall-kittens/cmd/durables"
	controller "tigerhall-kittens/cmd/handlers"
	"tigerhall-kittens/cmd/repository"
	"tigerhall-kittens/cmd/service"

	"github.com/gin-gonic/gin"
)

func SetupRouter(nrm gin.HandlerFunc) *gin.Engine {
	r := gin.Default()

	var (
		userRepository   repository.UserRepository   = repository.NewUserRepository(durable.MysqlDb)
		userService      service.UserService         = service.NewUserService(userRepository)
		userController   controller.UserController   = controller.NewUserController(userService)
		healthController controller.HealthController = controller.NewHealthController()
	)
	r.Use(nrm)

	grp1 := r.Group("/user/v1")
	{
		grp1.Use(nrm)
		grp1.GET("/fetch_all", userController.GetAllUsers)
		grp1.GET("/user_id/:userId", userController.GetUserByUserId)
		grp1.POST("/create_new", userController.CreateNewUser)
		grp1.PUT("/update", userController.UpdateUser)
		grp1.DELETE("/deletebyUserId/:userId", userController.DeleteUserById)
	}

	grp2 := r.Group("/")
	{
		grp2.GET("/health", healthController.GetHealth)
	}

	return r
}
