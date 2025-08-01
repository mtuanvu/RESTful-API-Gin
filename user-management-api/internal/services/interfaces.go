package services

import "mtuanvu.id.vn/restful-api-gin/internal/models"

type UserService interface {
	GetAllUsers() ([]models.User, error)
	CreateUser(user models.User) (models.User, error)
	GetUserByUUID(uuid string) (models.User, error)
	UpdateUser()
	DeleteUser()
}
