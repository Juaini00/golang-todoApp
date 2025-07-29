package main

// Package main provides the entry point for the todo REST API.
//
// @title Todo App API
// @version 1.0
// @description REST API for managing todos
// @BasePath /api/v1
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"log"
	"os"
	docs "todo_app/docs"
	"todo_app/internal/config"
	"todo_app/internal/delivery/middleware"
	"todo_app/internal/delivery/route"
)

func main() {
	config.LoadEnv()
	config.InitDB()

	router := gin.Default()

	router.Use(middleware.Logger())
	route.SetupRoutes(router, config.DB)

	docs.SwaggerInfo.BasePath = "/api/v1"
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	port := os.Getenv("PORT")
	err := router.Run(":" + port)

	if err != nil {
		log.Fatal("Failed to start server", err)
	}
	log.Println("Server started on port", os.Getenv("PORT"))
}
