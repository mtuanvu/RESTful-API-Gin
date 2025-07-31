package repositories

import (
	"log"

	"mtuanvu.id.vn/restful-api-gin/internal/models"
)

type inMemoryUserRepository struct {
	users []models.User
}

// constructor
func NewInMemoryUserRepository() UserRepository {
	return &inMemoryUserRepository{
		users: make([]models.User, 0),
	}
}

func (ur *inMemoryUserRepository) FindAll() {
	log.Println("Get All User From Repository")
}

func (ur *inMemoryUserRepository) Create(user models.User) error {
	ur.users = append(ur.users, user)
	return nil
}

func (ur *inMemoryUserRepository) FindByUUID() {

}

func (ur *inMemoryUserRepository) Update() {

}

func (ur *inMemoryUserRepository) Delete() {

}

func (ur *inMemoryUserRepository) FindByEmail(email string) (models.User, bool) {
	for _, user := range ur.users {
		if user.Email == email {
			return user, true
		}
	}

	return models.User{}, false
}