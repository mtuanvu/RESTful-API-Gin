package repositories

import "mtuanvu.id.vn/restful-api-gin/internal/models"

type UserRepository interface {
	FindAll()
	Create(user models.User) error
	FindByUUID()
	Update()
	Delete()
	FindByEmail(email string) (models.User, bool)
}
