package app

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"mtuanvu.id.vn/restful-api-gin/internal/config"
	"mtuanvu.id.vn/restful-api-gin/internal/routes"
	"mtuanvu.id.vn/restful-api-gin/internal/validations"
)

type Module interface {
	Routes() routes.Routes
}

type Application struct {
	config  *config.Config
	router  *gin.Engine
	modules []Module
}

func NewApplication(cfg *config.Config) *Application {
	validations.InitValidator()

	loadEnv()

	r := gin.Default()

	modules := []Module{
		NewUserModule(),
	}
	routes.RegisterRoutes(r, getModuleRoutes(modules)...)

	return &Application{
		config:  cfg,
		router:  r,
		modules: modules,
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

func loadEnv() {
	err := godotenv.Load("../../.env")
	if err != nil {
		log.Printf("Not found env")
	}
}
