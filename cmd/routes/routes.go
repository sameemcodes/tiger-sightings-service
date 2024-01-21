package router

import (
	durable "tigerhall-kittens/cmd/durables"
	controller "tigerhall-kittens/cmd/handlers"
	"tigerhall-kittens/cmd/middleware"
	"tigerhall-kittens/cmd/repository"
	"tigerhall-kittens/cmd/service"

	"github.com/gin-gonic/gin"
)

func SetupRouter(nrm gin.HandlerFunc) *gin.Engine {
	r := gin.Default()

	var (
		userRepository          repository.UserRepository          = repository.NewUserRepository(durable.MysqlDb)
		userService             service.UserService                = service.NewUserService(userRepository)
		userController          controller.UserController          = controller.NewUserController(userService)
		healthController        controller.HealthController        = controller.NewHealthController()
		tigerSightingRepository repository.TigerSightingRepository = repository.NewTigerSightingRepository(durable.MysqlDb)
		tigerSightingService    service.TigerSightingService       = service.NewTigerSightingService(tigerSightingRepository)
		tigerSightingController controller.TigerSightingController = controller.NewTigerSightingController(tigerSightingService)
		tigerRepository         repository.TigerRepository         = repository.NewTigerRepository(durable.MysqlDb)
		tigerService            service.TigerService               = service.NewTigerService(tigerRepository)
		tigerController         controller.TigerController         = controller.NewTigerController(tigerService)
	)

	r.Use(nrm)

	// User Routes
	grp1 := r.Group("/user/v1")
	{
		grp1.Use(nrm)
		grp1.GET("/fetch_all", userController.GetAllUsers)
		grp1.GET("/user_id/:userId", userController.GetUserByUserId)
		grp1.POST("/create_new", middleware.VerifyToken, userController.CreateNewUser)
		grp1.PUT("/update", userController.UpdateUser)
		grp1.DELETE("/deletebyUserId/:userId", userController.DeleteUserById)
		grp1.POST("/signup", userController.SignUp)
		grp1.POST("/login", userController.Login)

	}

	// Health Route
	grp2 := r.Group("/")
	{
		grp2.GET("/health", healthController.GetHealth)
	}

	// TigerSighting Routes
	grp3 := r.Group("/tigerSighting/v1")
	{
		grp3.Use(nrm)
		grp3.GET("/fetch_all", tigerSightingController.GetAllTigerSightings)
		grp3.GET("/sighting_id/:sightingId", tigerSightingController.GetTigerSightingById)
		grp3.POST("/create_new", middleware.VerifyToken, tigerSightingController.CreateNewTigerSighting)
		grp3.PUT("/update", tigerSightingController.UpdateTigerSighting)
		grp3.DELETE("/deletebySightingId/:sightingId", tigerSightingController.DeleteTigerSighting)
		grp3.GET("/tiger_id/:tigerId", tigerSightingController.GetTigerSightingsByTigerId)
		grp3.GET(("user_sightings/:tigerId"), tigerSightingController.GetUserSightingsListByTigerId)
	}

	// Tiger Routes
	grp4 := r.Group("/tiger/v1")
	{
		grp4.Use(nrm)
		grp4.GET("/fetch_all", tigerController.GetAllTigers)
		grp4.GET("/tiger_id/:tigerId", tigerController.GetTigerById)
		grp4.POST("/create_new", middleware.VerifyToken, tigerController.CreateNewTiger)
		grp4.PUT("/update", tigerController.UpdateTiger)
		grp4.DELETE("/deletebyTigerId/:tigerId", tigerController.DeleteTigerById)
		grp4.GET("/checkIfTigerExists/:tigerId", tigerController.CheckIfTigerExists)
	}

	return r
}
