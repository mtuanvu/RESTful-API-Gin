package services

import (
	"log"

	"mtuanvu.id.vn/restful-api-gin/internal/repositories"
)

type userService struct {
	repo repositories.UserRepository
}

func NewUserService(repo repositories.UserRepository) UserService {
	return &userService{
		repo: repo,
	}
}

func (us *userService) GetAllUsers() {
	log.Println("Get All User from User service")
	us.repo.FindAll()
}

func (us *userService) CreateUser() {

}

func (us *userService) GetUserByUUID() {

}

func (us *userService) UpdateUser() {

}

func (us *userService) DeleteUser() {

}
