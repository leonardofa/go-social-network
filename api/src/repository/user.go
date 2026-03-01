package repository

import (
	"database/sql"

	"api/src/model"
)

// User represents a repository for managing user data.
type User struct {
	db *sql.DB
}

// New returns a new UserRepository instance.
func New(db *sql.DB) *User {
	return &User{db}
}

// Create creates a new user.
func (repository *User) Create(user model.User) (uint64, error) {
	statement, err := repository.db.Prepare("INSERT INTO users (name, nick, email, password) VALUES (?, ?, ?, ?)")
	if err != nil {
		return 0, err
	}

	defer statement.Close()

	created, err := statement.Exec(user.Name, user.Nick, user.Email, user.Password)
	if err != nil {
		return 0, err
	}

	id, err := created.LastInsertId()
	if err != nil {
		return 0, err
	}

	return uint64(id), nil
}
