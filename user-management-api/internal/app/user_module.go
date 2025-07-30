package app

import (
	"mtuanvu.id.vn/restful-api-gin/internal/handlers"
	"mtuanvu.id.vn/restful-api-gin/internal/repositories"
	"mtuanvu.id.vn/restful-api-gin/internal/routes"
	"mtuanvu.id.vn/restful-api-gin/internal/services"
)

type UserModule struct {
	routes routes.Routes
}

func NewUserModule() *UserModule {

	//initialize repository
	userRepo := repositories.NewInMemoryUserRepository()

	//initialize service
	userService := services.NewUserService(userRepo)

	//initialize handler
	userHandler := handlers.NewUserHandler(userService)

	//initialize route
	userRoutes := routes.NewUserRoute(userHandler)

	return &UserModule{
		routes: userRoutes,
	}
}


func (m *UserModule) Routes() routes.Routes {
	return m.routes
}