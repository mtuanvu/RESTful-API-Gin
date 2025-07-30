package services

type UserService interface {
	GetAllUsers()
	CreateUser()
	GetUserByUUID()
	UpdateUser()
	DeleteUser()
}