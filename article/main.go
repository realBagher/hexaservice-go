package main

import (
	"fmt"
	"log"
	"os"

	"github.com/realBagher/hexaservice-go/article/adapters"
	"github.com/realBagher/hexaservice-go/article/core"
)

const (
	mysqlDSNEnvVar = "MYSQL_DSN"
	testArticleID  = "1"
	mysqlArticleID = "mysql_1"
)

func main() {
	if err := runDemo(); err != nil {
		log.Fatalf("Demo failed: %v", err)
	}
}

func runDemo() error {
	// Demonstrate InMemory repository
	if err := demonstrateInMemoryRepository(); err != nil {
		return fmt.Errorf("in-memory repository demo failed: %w", err)
	}

	// Demonstrate MySQL repository if DSN is available
	if dsn := os.Getenv(mysqlDSNEnvVar); dsn != "" {
		if err := demonstrateMySQLRepository(dsn); err != nil {
			return fmt.Errorf("MySQL repository demo failed: %w", err)
		}
	}

	return nil
}

func demonstrateInMemoryRepository() error {
	fmt.Println("=== Using InMemory Repository ===")

	repo := adapters.NewInMemoryArticleRepository()
	service := core.NewArticleService(repo)

	testArticle := createTestArticle(testArticleID)

	return demonstrateArticleOperations(service, testArticle)
}

func demonstrateMySQLRepository(dsn string) error {
	fmt.Println("\n=== Live MySQL Repository Demo ===")

	db, err := adapters.NewMySQLConnection(dsn)
	if err != nil {
		return fmt.Errorf("failed to connect to MySQL: %w", err)
	}
	defer func() {
		if closeErr := db.Close(); closeErr != nil {
			log.Printf("Warning: failed to close database connection: %v", closeErr)
		}
	}()

	repo := adapters.NewMySQLArticleRepository(db)
	if err := repo.InitializeSchema(); err != nil {
		return fmt.Errorf("failed to initialize schema: %w", err)
	}

	service := core.NewArticleService(repo)
	testArticle := createTestArticle(mysqlArticleID)

	return demonstrateArticleOperations(service, testArticle)
}

func createTestArticle(id string) core.Article {
	return core.Article{
		ID:        id,
		Title:     "Advanced Machine Learning Techniques",
		Abstract:  "This paper explores cutting-edge machine learning algorithms and their applications in modern data science.",
		AuthorID:  "author_1",
		JournalID: "journal_1",
	}
}

func demonstrateArticleOperations(service *core.ArticleService, article core.Article) error {
	// Create article
	createdArticle, err := service.CreateArticle(article)
	if err != nil {
		return fmt.Errorf("failed to create article: %w", err)
	}
	fmt.Printf("Created article: %+v\n", createdArticle)

	// Retrieve article by ID
	retrievedArticle, err := service.GetArticleByID(article.ID)
	if err != nil {
		return fmt.Errorf("failed to retrieve article by ID: %w", err)
	}
	fmt.Printf("Retrieved article by ID: %+v\n", retrievedArticle)

	// Retrieve article by title
	retrievedByTitle, err := service.GetArticleByTitle(article.Title)
	if err != nil {
		return fmt.Errorf("failed to retrieve article by title: %w", err)
	}
	fmt.Printf("Retrieved article by title: %+v\n", retrievedByTitle)

	return nil
}
