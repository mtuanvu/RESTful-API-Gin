package main

import (
	"mtuanvu.id.vn/restful-api-gin/internal/app"
	"mtuanvu.id.vn/restful-api-gin/internal/config"
)

func main() {
	//initialize configuration
	cfg := config.NewConfig()

	//initialize application
	application := app.NewApplication(cfg)

	//start server
	if err := application.Run(); err != nil {
		panic(err)
	}
}
