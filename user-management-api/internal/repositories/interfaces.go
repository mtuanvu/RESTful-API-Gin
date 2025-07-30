package repositories

type UserRepository interface {
	FindAll()
	Create()
	FindByUUID()
	Update()
	Delete()
}