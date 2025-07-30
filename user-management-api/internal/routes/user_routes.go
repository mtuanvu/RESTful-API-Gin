package routes

import (
	"github.com/gin-gonic/gin"
	"mtuanvu.id.vn/restful-api-gin/internal/handlers"
)

type UserRoute struct {
	handler *handlers.UserHandler
}

func NewUserRoute(handler *handlers.UserHandler) *UserRoute {
	return &UserRoute{
		handler: handler,
	}
}

func (ur *UserRoute) Register(r *gin.RouterGroup) {

}