package usecases

import (
	"github.com/honda-pp/memo-app-backend-go-gin/generated"
)

type UserRepositoryInterface interface {
	DeleteById(id int) error
	FindAll() ([]generated.User, error)
	FindById(id int) (generated.User, error)
}

func NewUserUsecase(userRepository UserRepositoryInterface) *UserUsecase {
	return &UserUsecase{
		UserRepository: userRepository,
	}
}

type UserUsecase struct {
	UserRepository UserRepositoryInterface
}

// DeleteById delete user by ID
func (u *UserUsecase) DeleteById(id int) error {
	// Your code here
	return nil
}

// FindAll returns a list of users.
func (u *UserUsecase) FindAll() ([]generated.User, error) {
	// Your code here
	return nil, nil
}

// FindById find user by ID
func (u *UserUsecase) FindById(id int) (generated.User, error) {
	// Your code here
	return generated.User{}, nil
}
