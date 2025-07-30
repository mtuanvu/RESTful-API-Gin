package repositories

import "mtuanvu.id.vn/restful-api-gin/internal/models"

type InMemoryUserRepository struct {
	users []models.User
}

// constructor
func NewInMemoryUserRepository() *InMemoryUserRepository {
	return &InMemoryUserRepository{
		users: make([]models.User, 0),
	}
}