package main

import (
	"github.com/HunCoding/golang-architecture/hexagonal-news-api/adapter/input/routes"
	"github.com/HunCoding/golang-architecture/hexagonal-news-api/configuration/logger"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	logger.Info("About to start application")

	logger.Info("Starting the application...")
	router := gin.Default()

	routes.InitRoutes(router)

	if err := router.Run(":8080"); err != nil {
		logger.Error("Error starting the application", err)
		return
	}

}
