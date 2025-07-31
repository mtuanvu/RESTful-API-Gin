package routes

import (
	"github.com/gin-gonic/gin"
	"mtuanvu.id.vn/restful-api-gin/internal/middleware"
)

type Routes interface {
	Register(r *gin.RouterGroup)
}

func RegisterRoutes(r *gin.Engine, routes ...Routes) {

	r.Use(
		middleware.LoggerMiddleware(),
		middleware.ApiKeyMiddleware(),
		middleware.AuthMiddleware(),
		middleware.RateLimiterMiddleware())

	api := r.Group("/api/v1")

	for _, route := range routes {
		route.Register(api)
	}
}
