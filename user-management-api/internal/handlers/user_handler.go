package handlers

import "mtuanvu.id.vn/restful-api-gin/internal/services"

type UserHandler struct {
	service *services.UserService
}

func NewUserHandler(services *services.UserService) *UserHandler {
	return &UserHandler{
		service: services,
	}
}