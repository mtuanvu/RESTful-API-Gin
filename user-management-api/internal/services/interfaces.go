package services

import "mtuanvu.id.vn/restful-api-gin/internal/models"

type UserService interface {
	GetAllUsers(search string, page, limit int) ([]models.User, error)
	CreateUser(user models.User) (models.User, error)
	GetUserByUUID(uuid string) (models.User, error)
	UpdateUser(uuid string, user models.User) (models.User, error)
	DeleteUser()
}
