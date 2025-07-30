package handlers

import (
	"log"

	"github.com/gin-gonic/gin"
	"mtuanvu.id.vn/restful-api-gin/internal/services"
)

type UserHandler struct {
	service services.UserService
}

func NewUserHandler(services services.UserService) *UserHandler {
	return &UserHandler{
		service: services,
	}
}

func (uh *UserHandler) GetAllUsers(ctx *gin.Context) {
	log.Println("Get All User from User Handler")

	uh.service.GetAllUsers()
}

func (uh *UserHandler) CreateUser(ctx *gin.Context) {

}

func (uh *UserHandler) GetUserByUUID(ctx *gin.Context) {

}

func (uh *UserHandler) UpdateUser(ctx *gin.Context) {

}

func (uh *UserHandler) DeleteUser(ctx *gin.Context) {

}
