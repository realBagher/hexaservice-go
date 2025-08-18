package main

import (
	"fmt"
	"log"

	"github.com/realBagher/hexaservice-go/article/adapters"
	"github.com/realBagher/hexaservice-go/article/core"
)

const testArticleID = "1"

func main() {
	if err := runDemo(); err != nil {
		log.Fatalf("Demo failed: %v", err)
	}
}

func runDemo() error {
	repo := adapters.NewInMemoryArticleRepository()
	service := core.NewArticleService(repo)

	testArticle := core.Article{
		ID:        testArticleID,
		Title:     "Test Article",
		Abstract:  "Test abstract for article",
		AuthorID:  "author_1",
		JournalID: "journal_1",
	}

	// Create article
	createdArticle, err := service.CreateArticle(testArticle)
	if err != nil {
		return fmt.Errorf("failed to create article: %w", err)
	}
	fmt.Printf("Created article: %+v\n", createdArticle)

	// Retrieve article
	retrievedArticle, err := service.GetArticleByID(testArticleID)
	if err != nil {
		return fmt.Errorf("failed to retrieve article: %w", err)
	}
	fmt.Printf("Retrieved article: %+v\n", retrievedArticle)

	return nil
}
