package services

import "mtuanvu.id.vn/restful-api-gin/internal/models"

type UserService interface {
	GetAllUsers()
	CreateUser(user models.User) (models.User, error)
	GetUserByUUID()
	UpdateUser()
	DeleteUser()
}
