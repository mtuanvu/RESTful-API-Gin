package app

import (
	"github.com/gin-gonic/gin"
	"mtuanvu.id.vn/restful-api-gin/internal/config"
	"mtuanvu.id.vn/restful-api-gin/internal/routes"
)

type Module interface {
	Routes() routes.Routes
}

type Application struct {
	config *config.Config
	router *gin.Engine
}

func NewApplication(cfg *config.Config) *Application {
	r := gin.Default()

	modules := []Module{
		NewUserModule(),
	}
	routes.RegisterRoutes(r, getModuleRoutes(modules)...)

	return &Application{
		config: cfg,
		router: r,
	}
}

func (a *Application) Run() error {
	return a.router.Run(a.config.ServerAddress)
}

func getModuleRoutes(modules []Module) []routes.Routes {
	routeList := make([]routes.Routes, len(modules))

	for i, module := range modules {
		routeList[i] = module.Routes()
	}

	return routeList
}
