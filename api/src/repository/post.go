package repository

import (
	"api/src/model"
	"database/sql"
)

// Posts represents a repository for managing posts in the database.
type Posts struct {
	db *sql.DB
}

// NewPostsRepository creates a new instance of Posts repository with the given database connection.
func NewPostsRepository(db *sql.DB) *Posts {
	return &Posts{db}
}

// Create inserts a new post into the database.
func (repository Posts) Create(post model.Post) (uint64, error) {
	statement, err := repository.db.Prepare(
		"INSERT INTO posts (title, content, author_id) VALUES (?, ?, ?)",
	)
	if err != nil {
		return 0, err
	}
	defer statement.Close()

	result, err := statement.Exec(post.Title, post.Content, post.AuthorID)
	if err != nil {
		return 0, err
	}

	lastInsertID, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return uint64(lastInsertID), nil
}

// FindByID retrieves a post by its ID.
func (repository Posts) FindByID(postID uint64) (model.Post, error) {
	row, err := repository.db.Query(`
		SELECT p.*, u.nick FROM 
		posts p INNER JOIN users u ON u.id = p.author_id 
		WHERE p.id = ?`,
		postID,
	)
	if err != nil {
		return model.Post{}, err
	}
	defer row.Close()

	var post model.Post

	if row.Next() {
		if err = row.Scan(
			&post.ID,
			&post.Title,
			&post.Content,
			&post.AuthorID,
			&post.Likes,
			&post.CreatedAt,
			&post.AuthorNick,
		); err != nil {
			return model.Post{}, err
		}
	}

	return post, nil
}

// FindAll retrieves all posts visible to the given user, including their posts and posts from followed users.
func (repository Posts) FindAll(userID uint64) ([]model.Post, error) {
	rows, err := repository.db.Query(`
		SELECT DISTINCT p.*, u.nick FROM posts p 
		INNER JOIN users u ON u.id = p.author_id 
		INNER JOIN followers f ON p.author_id = f.user_id 
		WHERE u.id = ? OR f.follower_id = ?
		ORDER BY 1 DESC`,
		userID, userID,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var posts []model.Post

	for rows.Next() {
		var post model.Post

		if err = rows.Scan(
			&post.ID,
			&post.Title,
			&post.Content,
			&post.AuthorID,
			&post.Likes,
			&post.CreatedAt,
			&post.AuthorNick,
		); err != nil {
			return nil, err
		}

		posts = append(posts, post)
	}

	return posts, nil
}

// Update updates the title and content of a post in the database with the given postID. Returns an error if the update fails.
func (repository Posts) Update(postID uint64, post model.Post) error {
	statement, err := repository.db.Prepare("UPDATE posts SET title = ?, content = ? WHERE id = ?")
	if err != nil {
		return err
	}
	defer statement.Close()

	if _, err = statement.Exec(post.Title, post.Content, post.ID); err != nil {
		return err
	}

	return nil
}

// Delete removes a post from the database based on the provided postID and returns an error if the operation fails.
func (repository Posts) Delete(postID uint64) error {
	statement, err := repository.db.Prepare("DELETE FROM posts WHERE id = ?")
	if err != nil {
		return err
	}
	defer statement.Close()

	if _, err = statement.Exec(postID); err != nil {
		return err
	}

	return nil
}

// FindByUserID retrieves all posts authored by a specific user identified by their userID. Returns a slice of posts or an error.
func (repository Posts) FindByUserID(userID uint64) ([]model.Post, error) {
	rows, err := repository.db.Query(`
		SELECT p.*, u.nick FROM posts p
		JOIN users u ON u.id = p.author_id
		WHERE p.author_id = ?`,
		userID,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var posts []model.Post

	for rows.Next() {
		var post model.Post

		if err = rows.Scan(
			&post.ID,
			&post.Title,
			&post.Content,
			&post.AuthorID,
			&post.Likes,
			&post.CreatedAt,
			&post.AuthorNick,
		); err != nil {
			return nil, err
		}

		posts = append(posts, post)
	}

	return posts, nil
}

// Like increments the like count of a post identified by postID in the database. Returns an error if the operation fails.
func (repository Posts) Like(postID uint64) error {
	statement, err := repository.db.Prepare("UPDATE posts SET likes = likes + 1 WHERE id = ?")
	if err != nil {
		return err
	}
	defer statement.Close()

	if _, err = statement.Exec(postID); err != nil {
		return err
	}

	return nil
}

// Unlike decrements the like count of a post identified by postID in the database, ensuring it does not drop below zero.
func (repository Posts) Unlike(postID uint64) error {
	statement, err := repository.db.Prepare(`
		UPDATE posts SET likes = 
		CASE 
			WHEN likes > 0 THEN likes - 1 
			ELSE 0 
		END 
		WHERE id = ?`,
	)
	if err != nil {
		return err
	}
	defer statement.Close()

	if _, err = statement.Exec(postID); err != nil {
		return err
	}

	return nil
}
