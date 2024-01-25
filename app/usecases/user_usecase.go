package usecases

import (
	"github.com/honda-pp/memo-app-backend-go-gin/app/interfaces"
	"github.com/honda-pp/memo-app-backend-go-gin/generated"
)

func NewUserUsecase(userRepository interfaces.UserRepositoryInterface) *UserUsecase {
	return &UserUsecase{
		UserRepository: userRepository,
	}
}

type UserUsecase struct {
	UserRepository interfaces.UserRepositoryInterface
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
