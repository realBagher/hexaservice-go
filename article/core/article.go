package core

import (
	"fmt"
	"strings"
)

type Article struct {
	ID        string `json:"id"`
	Title     string `json:"title"`
	Abstract  string `json:"abstract"`
	AuthorID  string `json:"author_id"`
	JournalID string `json:"journal_id"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

// Validate checks if the article data is valid
func (a Article) Validate() error {
	if strings.TrimSpace(a.ID) == "" {
		return fmt.Errorf("%w: ID cannot be empty", ErrInvalidArticle)
	}

	if strings.TrimSpace(a.Title) == "" {
		return fmt.Errorf("%w: title cannot be empty", ErrInvalidArticle)
	}

	if strings.TrimSpace(a.AuthorID) == "" {
		return fmt.Errorf("%w: author ID cannot be empty", ErrInvalidArticle)
	}

	if strings.TrimSpace(a.JournalID) == "" {
		return fmt.Errorf("%w: journal ID cannot be empty", ErrInvalidArticle)
	}

	return nil
}

type ArticleService struct {
	repository ArticleRepository
}

func NewArticleService(repository ArticleRepository) *ArticleService {
	return &ArticleService{repository: repository}
}

func (s *ArticleService) CreateArticle(article Article) (Article, error) {
	if err := article.Validate(); err != nil {
		return Article{}, err
	}
	return s.repository.CreateArticle(article)
}

func (s *ArticleService) GetArticleByID(id string) (Article, error) {
	return s.repository.GetArticleByID(id)
}

func (s *ArticleService) GetArticleByTitle(title string) (Article, error) {
	return s.repository.GetArticleByTitle(title)
}
