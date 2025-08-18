package core

import "errors"

var (
	// ErrArticleNotFound is returned when an article is not found
	ErrArticleNotFound = errors.New("article not found")

	// ErrInvalidArticle is returned when article data is invalid
	ErrInvalidArticle = errors.New("invalid article data")
)
