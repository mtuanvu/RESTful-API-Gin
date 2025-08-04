package repositories

import (
	"fmt"
	"slices"

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

func (ur *inMemoryUserRepository) FindAll() ([]models.User, error) {
	return ur.users, nil
}

func (ur *inMemoryUserRepository) Create(user models.User) error {
	ur.users = append(ur.users, user)
	return nil
}

func (ur *inMemoryUserRepository) FindByUUID(uuid string) (models.User, bool) {
	for _, user := range ur.users {
		if user.UUID == uuid {
			return user, true
		}
	}

	return models.User{}, false
}

func (ur *inMemoryUserRepository) Update(uuid string, user models.User) error {
	for i, u := range ur.users {
		if u.UUID == uuid {
			ur.users[i] = user
			return nil
		}
	}

	return fmt.Errorf("user not found")
}

func (ur *inMemoryUserRepository) Delete(uuid string) error {
	for i, u := range ur.users {
		if u.UUID == uuid {
			ur.users = slices.Delete(ur.users, i, i + 1)
			return nil
		}
	}

	return fmt.Errorf("User not found")
}

func (ur *inMemoryUserRepository) FindByEmail(email string) (models.User, bool) {
	for _, user := range ur.users {
		if user.Email == email {
			return user, true
		}
	}

	return models.User{}, false
}
