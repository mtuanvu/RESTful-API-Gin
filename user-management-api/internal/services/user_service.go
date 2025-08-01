package services

import (
	"strings"

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

func (us *userService) GetAllUsers(search string, page, limit int) ([]models.User, error) {
	users, err := us.repo.FindAll()
	if err != nil {
		return nil, utils.WrapError(err, "failed to fetch users", utils.ErrorCodeInternal)
	}

	var filteredUsers []models.User

	if search != "" {
		search = strings.ToLower(search)
		for _, user := range users {
			name := strings.ToLower(user.Name)
			email := strings.ToLower(user.Email)

			if strings.Contains(name, search) || strings.Contains(email, search) {
				filteredUsers = append(filteredUsers, user)
			}
		}
	} else {
		filteredUsers = users
	}

	start := (page - 1) * limit
	if start >= len(filteredUsers) {
		return []models.User{}, nil
	}

	end := start + limit
	if end > len(filteredUsers) {
		end = len(filteredUsers)
	}

	return filteredUsers[start:end], nil
}

func (us *userService) CreateUser(user models.User) (models.User, error) {
	user.Email = utils.NormalizeString(user.Email)

	if _, exist := us.repo.FindByEmail(user.Email); exist {
		return models.User{}, utils.NewError("Email already exist", utils.ErrorCodeConflict)
	}

	user.UUID = uuid.New().String()

	hashPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return models.User{}, utils.WrapError(err, "failed to hash password", utils.ErrorCodeInternal)
	}

	user.Password = string(hashPassword)

	if err := us.repo.Create(user); err != nil {
		return models.User{}, utils.WrapError(err, "failed to create user", utils.ErrorCodeInternal)
	}

	return user, nil

}

func (us *userService) GetUserByUUID(uuid string) (models.User, error) {
	user, found := us.repo.FindByUUID(uuid)
	if !found {
		return models.User{}, utils.NewError("user not found", utils.ErrorCodeNotFound)
	}

	return user, nil
}

func (us *userService) UpdateUser(uuid string, updateUser models.User) (models.User, error) {
	updateUser.Email = utils.NormalizeString(updateUser.Email)

	if u, exist := us.repo.FindByEmail(updateUser.Email); exist && u.UUID != uuid {
		return models.User{}, utils.NewError("Email already exist", utils.ErrorCodeConflict)
	}

	currentUser, found := us.repo.FindByUUID(uuid)
	if !found {
		return models.User{}, utils.NewError("User not found", utils.ErrorCodeNotFound)
	}

	currentUser.Name = updateUser.Name
	currentUser.Email = updateUser.Email
	currentUser.Age = updateUser.Age
	currentUser.Status = updateUser.Status
	currentUser.Level = updateUser.Level

	if updateUser.Password != "" {
		hashPassword, err := bcrypt.GenerateFromPassword([]byte(updateUser.Password), bcrypt.DefaultCost)
		if err != nil {
			return models.User{}, utils.WrapError(err, "failed to hash password", utils.ErrorCodeInternal)
		}

		currentUser.Password = string(hashPassword)
	}

	if err := us.repo.Update(uuid, currentUser); err != nil {
		return models.User{}, utils.WrapError(err, "failed to update user", utils.ErrorCodeInternal)
	}

	return currentUser, nil
}

func (us *userService) DeleteUser() {

}
