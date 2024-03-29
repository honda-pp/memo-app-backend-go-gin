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

// CreateUser create a new user
func (r *UserRepository) CreateUser(user generated.User) error {
	query := "INSERT INTO users (name) VALUES ($1)"
	_, err := r.db.Exec(query, user.Name)
	if err != nil {
		return err
	}
	return nil
}

// DeleteById delete user by ID
func (r *UserRepository) DeleteById(id int) error {
	query := "DELETE FROM users WHERE id = $1"
	_, err := r.db.Exec(query, id)
	if err != nil {
		return err
	}
	return nil
}

// FindAll returns a list of users.
func (r *UserRepository) FindAll() ([]generated.User, error) {
	query := "SELECT id, name FROM users"

	users := []generated.User{}

	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		user := generated.User{}
		err = rows.Scan(&user.Id, &user.Name)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	return users, nil
}

// FindById find user by ID
func (r *UserRepository) FindById(id int) (*generated.User, error) {
	query := "SELECT id, name FROM users WHERE id = $1"

	user := &generated.User{}

	row := r.db.QueryRow(query, id)
	err := row.Scan(&user.Id, &user.Name)
	if err != nil {
		return nil, err
	}
	return user, nil
}
