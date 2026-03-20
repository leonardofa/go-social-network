package model

import (
	"errors"
	"strings"
	"time"
)

// User represents a user in the application (social network).
type User struct {
	ID        uint64    `json:"id,omitempty"`
	Name      string    `json:"name,omitempty"`
	Nick      string    `json:"nick,omitempty"`
	Email     string    `json:"email,omitempty"`
	Password  string    `json:"password,omitempty"`
	CreatedAt time.Time `json:"createdAt,omitempty"`
}

// Prepare prepares the user for insertion into the database.
func (user *User) Prepare(step string) error {
	user.normalizeFields()
	return user.validate(step)
}

// validate validates the user data.
func (user *User) validate(step string) error {
	switch {
	case user.Name == "":
		return errors.New("user name cannot be empty")
	case user.Nick == "":
		return errors.New("user nick cannot be empty")
	case user.Email == "":
		return errors.New("user email cannot be empty")
	case step == "create" && user.Password == "":
		return errors.New("user password cannot be empty")
	default:
		return nil
	}
}

// normalizeFields normalizes the user data.
func (user *User) normalizeFields() {
	user.Name = strings.TrimSpace(user.Name)
	user.Nick = strings.TrimSpace(user.Nick)
	user.Email = strings.TrimSpace(user.Email)
}
