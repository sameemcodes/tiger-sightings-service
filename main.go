package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
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
	config.LoadEnv()
	durable.MessageQueueVariable = durable.NewMessageQueue(10)

	ctx, cancel := context.WithCancel(context.Background())

	go func() {
		sigCh := make(chan os.Signal, 1)
		signal.Notify(sigCh, syscall.SIGINT, syscall.SIGTERM)
		<-sigCh
		cancel()
	}()

	workerPool := durable.NewWorkerPool(10, durable.MessageQueueVariable)
	workerPool.StartWorkers()

	durable.InitMysqlDb()

	r := router.SetupRouter(GinContextToContextMiddleware())
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	go func() {
		log.Println("Server listening on :8036")
		err := r.Run(":8036")
		if err != nil && err != http.ErrServerClosed {
			log.Printf("HTTP server error: %v\n", err)
			cancel()
		}
	}()

	<-ctx.Done()

	log.Println("Shutting down...")

	utils.HandleError("Error while starting server ", nil)
	defer durable.CloseDbConnection()
	defer durable.MessageQueueVariable.Dequeue() // Close the message queue
}

func GinContextToContextMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()
	}
}
