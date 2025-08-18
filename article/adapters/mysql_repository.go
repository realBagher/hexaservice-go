package adapters

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/realBagher/hexaservice-go/article/core"
)

type MySQLArticleRepository struct {
	db *sql.DB
}

func NewMySQLArticleRepository(db *sql.DB) *MySQLArticleRepository {
	return &MySQLArticleRepository{db: db}
}

// NewMySQLConnection creates a new MySQL database connection
func NewMySQLConnection(dsn string) (*sql.DB, error) {
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, fmt.Errorf("failed to open database connection: %w", err)
	}

	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("failed to ping database: %w", err)
	}

	return db, nil
}

// InitializeSchema creates the articles table if it doesn't exist
func (r *MySQLArticleRepository) InitializeSchema() error {
	query := `
	CREATE TABLE IF NOT EXISTS articles (
		id VARCHAR(255) PRIMARY KEY,
		title VARCHAR(500) NOT NULL,
		abstract TEXT,
		author_id VARCHAR(255) NOT NULL,
		journal_id VARCHAR(255) NOT NULL,
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
		updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
	)`

	_, err := r.db.Exec(query)
	if err != nil {
		return fmt.Errorf("failed to create articles table: %w", err)
	}

	return nil
}

func (r *MySQLArticleRepository) CreateArticle(article core.Article) (core.Article, error) {
	query := `
	INSERT INTO articles (id, title, abstract, author_id, journal_id, created_at, updated_at) 
	VALUES (?, ?, ?, ?, ?, ?, ?)`

	_, err := r.db.Exec(query, article.ID, article.Title, article.Abstract,
		article.AuthorID, article.JournalID, article.CreatedAt, article.UpdatedAt)
	if err != nil {
		return core.Article{}, fmt.Errorf("failed to create article: %w", err)
	}

	return article, nil
}

func (r *MySQLArticleRepository) GetArticleByID(id string) (core.Article, error) {
	var article core.Article
	query := `
	SELECT id, title, abstract, author_id, journal_id, created_at, updated_at 
	FROM articles 
	WHERE id = ?`

	row := r.db.QueryRow(query, id)
	err := row.Scan(&article.ID, &article.Title, &article.Abstract,
		&article.AuthorID, &article.JournalID, &article.CreatedAt, &article.UpdatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return core.Article{}, core.ErrArticleNotFound
		}
		return core.Article{}, fmt.Errorf("failed to get article by ID: %w", err)
	}

	return article, nil
}

func (r *MySQLArticleRepository) GetArticleByTitle(title string) (core.Article, error) {
	var article core.Article
	query := `
	SELECT id, title, abstract, author_id, journal_id, created_at, updated_at 
	FROM articles 
	WHERE title = ?`

	row := r.db.QueryRow(query, title)
	err := row.Scan(&article.ID, &article.Title, &article.Abstract,
		&article.AuthorID, &article.JournalID, &article.CreatedAt, &article.UpdatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return core.Article{}, core.ErrArticleNotFound
		}
		return core.Article{}, fmt.Errorf("failed to get article by title: %w", err)
	}

	return article, nil
}

// Close closes the database connection
func (r *MySQLArticleRepository) Close() error {
	return r.db.Close()
}
