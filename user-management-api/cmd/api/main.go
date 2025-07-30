package main

import (
	"github.com/gin-gonic/gin"
	"mtuanvu.id.vn/restful-api-gin/internal/config"
	"mtuanvu.id.vn/restful-api-gin/internal/handlers"
	"mtuanvu.id.vn/restful-api-gin/internal/repositories"
	"mtuanvu.id.vn/restful-api-gin/internal/routes"
	"mtuanvu.id.vn/restful-api-gin/internal/services"
)

func main() {
	//initialize configuration
	cfg := config.NewConfig()

	//initialize repository
	userRepo := repositories.NewInMemoryUserRepository()

	//initialize service
	userService := services.NewUserService(userRepo)

	//initialize handler
	userHandler := handlers.NewUserHandler(userService)

	//initialize route
	userRoutes := routes.NewUserRoute(userHandler)

	r := gin.Default()

	routes.RegisterRoutes(r, userRoutes)

	if err := r.Run(cfg.ServerAddress); err != nil {
		panic(err)
	}
}
