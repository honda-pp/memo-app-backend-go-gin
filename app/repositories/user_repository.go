package repositories

import (
	"database/sql"

	"github.com/honda-pp/memo-app-backend-go-gin/generated"
)

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{
		db: db,
	}
}

// DeleteById delete user by ID
func (r *UserRepository) DeleteById(id int) error {
	// Your code here
	return nil
}

// FindAll returns a list of users.
func (r *UserRepository) FindAll() ([]generated.User, error) {
	// Your code here
	return nil, nil
}

// FindById find user by ID
func (r *UserRepository) FindById(id int) (generated.User, error) {
	// Your code here
	return generated.User{}, nil
}
