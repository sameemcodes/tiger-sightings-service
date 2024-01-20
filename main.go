package main

import (
	"log"
	durable "tigerhall-kittens/cmd/durables"
	router "tigerhall-kittens/cmd/routes"
	"tigerhall-kittens/cmd/utils"

	config "tigerhall-kittens/config"

	_ "github.com/swaggo/swag"

	_ "tigerhall-kittens/docs"

	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title tiggerhall-kittens
// @version 1.0
// @description Tiggerhall-Kittens
// @contact.name Mohamed Sameem
// @contact.email mmmohamedsameem@gmail.com

// @BasePath /

func main() {
	config.LoadEnv() // Load environment variables

	durable.InitMysqlDb() // Initialize MySQL database
	r := router.SetupRouter(GinContextToContextMiddleware())
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	r.Run(":8036")
	log.Println("Starting streamerx-backend service at port : 8036")
	utils.HandleError("Error while starting server ", nil)
	defer durable.CloseDbConnection()

}

func GinContextToContextMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()
	}
}
