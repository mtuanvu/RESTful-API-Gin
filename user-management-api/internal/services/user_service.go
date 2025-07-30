package services

import "mtuanvu.id.vn/restful-api-gin/internal/repositories"

type UserService struct {
	repo *repositories.InMemoryUserRepository
}

func NewUserService(repo *repositories.InMemoryUserRepository) *UserService {
	return &UserService{
		repo: repo,
	}
}