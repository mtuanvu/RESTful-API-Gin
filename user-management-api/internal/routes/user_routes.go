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
		users := r.Group("/users")
		{
			users.GET("", ur.handler.GetAllUsers)
			users.POST("", ur.handler.CreateUser)
			users.GET("/:uuid", ur.handler.GetUserByUUID)
			users.PUT("/:uuid", ur.handler.UpdateUser)
			users.DELETE("/:uuid", ur.handler.DeleteUser)
		}
}