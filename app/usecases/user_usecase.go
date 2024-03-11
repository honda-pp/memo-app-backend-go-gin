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

// CreateUser create a new user
func (u *UserUsecase) CreateUser(user generated.User) error {
	return u.UserRepository.CreateUser(user)
}

// DeleteById delete user by ID
func (u *UserUsecase) DeleteById(id int) error {
	return u.UserRepository.DeleteById(id)
}

// FindAll returns a list of users.
func (u *UserUsecase) FindAll() ([]generated.User, error) {
	return u.UserRepository.FindAll()
}

// FindById find user by ID
func (u *UserUsecase) FindById(id int) (*generated.User, error) {
	return u.UserRepository.FindById(id)
}
