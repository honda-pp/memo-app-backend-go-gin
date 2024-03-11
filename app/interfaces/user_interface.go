package interfaces

import "github.com/honda-pp/memo-app-backend-go-gin/generated"

type UserUsecaseInterface interface {
	CreateUser(generated.User) error
	DeleteById(id int) error
	FindAll() ([]generated.User, error)
	FindById(id int) (*generated.User, error)
}

type UserRepositoryInterface interface {
	CreateUser(generated.User) error
	DeleteById(id int) error
	FindAll() ([]generated.User, error)
	FindById(id int) (*generated.User, error)
}
