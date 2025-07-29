package main

import (
	"fmt"
	"todo_app/internal/config"
	"todo_app/internal/domain/entity"
)

func init() {
	config.LoadEnv()
	config.InitDB()
}

func main() {
	err := config.DB.AutoMigrate(
		&entity.User{},
		&entity.UserDetail{},
		&entity.Todo{},
	)
	if err != nil {
		return
	}

	fmt.Println("Migration successfully")
}
