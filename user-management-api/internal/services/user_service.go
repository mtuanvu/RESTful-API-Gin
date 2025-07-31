package services

import (
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"mtuanvu.id.vn/restful-api-gin/internal/models"
	"mtuanvu.id.vn/restful-api-gin/internal/repositories"
	"mtuanvu.id.vn/restful-api-gin/internal/utils"
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
}

func (us *userService) CreateUser(user models.User) (models.User, error){
	user.Email = utils.NormalizeString(user.Email)

	if _, exist := us.repo.FindByEmail(user.Email); exist {
		return models.User{}, utils.NewError("Email already exist", utils.ErrorCodeConflict)
	}

	user.UUID = uuid.New().String()

	hashPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return models.User{}, utils.WrapError(err ,"failed to hash password", utils.ErrorCodeInternal)
	}

	user.Password = string(hashPassword)

	if err := us.repo.Create(user); err != nil {
		return models.User{}, utils.WrapError(err ,"failed to create user", utils.ErrorCodeInternal)
	}

	return user, nil
	
}

func (us *userService) GetUserByUUID() {

}

func (us *userService) UpdateUser() {

}

func (us *userService) DeleteUser() {

}
