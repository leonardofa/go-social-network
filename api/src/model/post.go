package model

import (
	"errors"
	"strings"
	"time"
)

// Post represents a post made by a user.
type Post struct {
	ID         uint64    `json:"id,omitempty"`
	Title      string    `json:"title,omitempty"`
	Content    string    `json:"content,omitempty"`
	AuthorID   uint64    `json:"authorId,omitempty"`
	AuthorNick string    `json:"authorNick,omitempty"`
	Likes      uint64    `json:"likes"`
	CreatedAt  *time.Time `json:"createdAt,omitempty"`
}

// Prepare prepares the post for insertion or update.
func (post *Post) Prepare() error {
	post.normalizeFields()
	return post.validate()
}

// validate validates the post-data.
func (post *Post) validate() error {
	if post.Title == "" {
		return errors.New("post title cannot be empty")
	}

	if post.Content == "" {
		return errors.New("post content cannot be empty")
	}

	return nil
}

// normalizeFields normalizes the post-data by trimming whitespace.
func (post *Post) normalizeFields() {
	post.Title = strings.TrimSpace(post.Title)
	post.Content = strings.TrimSpace(post.Content)
}
