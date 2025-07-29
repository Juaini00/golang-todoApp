package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"os"
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
	port := os.Getenv("PORT")
	err := router.Run(":" + port)

	if err != nil {
		log.Fatal("Failed to start server", err)
	}
	log.Println("Server started on port", os.Getenv("PORT"))
}
