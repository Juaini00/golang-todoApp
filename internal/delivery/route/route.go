package route

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"todo_app/internal/delivery/http"
	"todo_app/internal/delivery/middleware"
	"todo_app/internal/repository"
	"todo_app/internal/usecase"
)

func SetupRoutes(r *gin.Engine, db *gorm.DB) {

	UserRepository := repository.NewUserRepository(db)
	UserDetailRepository := repository.NewUserDetailRepository(db)
	userUsecase := usecase.NewUserUsecase(UserRepository, UserDetailRepository)
	authHandler := http.NewAuthHandler(userUsecase)

	api := r.Group("/api/v1")
	auth := api.Group("/auth")
	{
		auth.POST("/register", authHandler.Register)
		auth.POST("/login", authHandler.Login)
	}

	protected := api.Group("/")
	protected.Use(middleware.AuthenticationMiddleware(UserRepository, UserDetailRepository))
	{
		protected.GET("/test", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"status":  200,
				"message": "Hello World!",
			})
		})
	}
}
