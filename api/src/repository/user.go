package repository

import (
	"api/src/security"
	"database/sql"
	"strings"

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

// FindByNameOrNick retrieves users whose name or nickname matches the given input (case-insensitive), performing a SQL LIKE query.
func (repository *User) FindByNameOrNick(nameOrNick string) ([]model.User, error) {
	nameOrNickAsLike := strings.ToLower("%" + nameOrNick + "%")
	lines, err := repository.db.Query(
		"SELECT id, name, nick, email, created_at FROM users where name like ? or nick like ?",
		nameOrNickAsLike, nameOrNickAsLike,
	)
	if err != nil {
		return nil, err
	}
	defer lines.Close()

	var users []model.User
	for lines.Next() {
		var user model.User
		if err := lines.Scan(&user.ID, &user.Name, &user.Nick, &user.Email, &user.CreatedAt); err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	return users, nil
}

// FindByID retrieves a user from the database by their unique ID and returns the user data or an error if not found.
func (repository *User) FindByID(userID uint64) (model.User, error) {
	lines, err := repository.db.Query("SELECT id, name, nick, email, created_at FROM users WHERE id = ?",
		userID,
	)
	if err != nil {
		return model.User{}, err
	}
	defer lines.Close()

	var user model.User
	for lines.Next() {
		if err := lines.Scan(&user.ID, &user.Name, &user.Nick, &user.Email, &user.CreatedAt); err != nil {
			return model.User{}, err
		}
	}

	return user, nil
}

// Update updates an existed user.
func (repository *User) Update(userID uint64, user model.User) error {
	statement, err := repository.db.Prepare("UPDATE users SET name = ?, nick = ?, email = ? WHERE id = ?")
	if err != nil {
		return err
	}
	defer statement.Close()

	_, err = statement.Exec(user.Name, user.Nick, user.Email, userID)
	if err != nil {
		return err
	}

	return nil
}

// DeleteByID deletes a user from the database by their unique ID.
func (repository *User) DeleteByID(userID uint64) error {
	statement, err := repository.db.Prepare("DELETE FROM users WHERE id = ?")
	if err != nil {
		return err
	}
	defer statement.Close()

	_, err = statement.Exec(userID)
	if err != nil {
		return err
	}

	return nil
}

// Login authenticates a user by validating their email and password against the database and returns the user or an error.
func (repository *User) Login(userParam model.User) (model.User, error) {
	var user = model.User{}
	lines, err := repository.db.Query("SELECT id, password FROM users WHERE email = ?", userParam.Email)
	if err != nil {
		return user, nil
	}
	defer lines.Close()

	if lines.Next() {
		if err := lines.Scan(&user.ID, &user.Password); err != nil {
			return user, nil
		}
	} else {
		return user, nil
	}

	err = security.CompareHashAndPassword(user.Password, userParam.Password)

	if err != nil {
		return model.User{}, nil
	}

	return user, nil
}

// Follow establishes a follower relationship between a user and a follower based on their unique IDs.
func (repository *User) Follow(followerUserId uint64, followingUserId uint64) error {
	statement, err := repository.db.Prepare("INSERT IGNORE INTO followers (user_id, follower_id) VALUES (?, ?)")
	if err != nil {
		return err
	}
	defer statement.Close()

	_, err = statement.Exec(followingUserId, followerUserId)
	if err != nil {
		return err
	}

	return nil
}

// Unfollow removes the follower relationship between a user and a follower based on their unique IDs.
func (repository *User) Unfollow(followerUserId uint64, followingUserId uint64) error {
	statement, err := repository.db.Prepare("DELETE FROM followers WHERE user_id = ? AND follower_id = ?")
	if err != nil {
		return err
	}
	defer statement.Close()

	_, err = statement.Exec(followingUserId, followerUserId)
	if err != nil {
		return err
	}

	return nil
}

// ReadFollowingList retrieves the list of users that a specific user is following based on their unique ID.
func (repository *User) ReadFollowingList(followerId uint64) ([]model.User, error) {
	lines, err := repository.db.Query(`
		SELECT u.id, u.name, u.nick, u.email, u.created_at
		FROM users u INNER JOIN followers f ON u.id = f.user_id
		WHERE f.follower_id = ?
	`, followerId)
	if err != nil {
		return nil, err
	}
	defer lines.Close()

	var users []model.User
	for lines.Next() {
		var user model.User
		if err := lines.Scan(&user.ID, &user.Name, &user.Nick, &user.Email, &user.CreatedAt); err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	return users, nil
}

// ReadFollowersList retrieves the list of followers for a specified user.
func (repository *User) ReadFollowersList(userID uint64) ([]model.User, error) {
	lines, err := repository.db.Query(`
		SELECT u.id, u.name, u.nick, u.email, u.created_at
		FROM users u INNER JOIN followers f ON u.id = f.follower_id
		WHERE f.user_id = ?
	`, userID)
	if err != nil {
		return nil, err
	}
	defer lines.Close()

	var users []model.User
	for lines.Next() {
		var user model.User
		if err := lines.Scan(&user.ID, &user.Name, &user.Nick, &user.Email, &user.CreatedAt); err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	return users, nil

}
