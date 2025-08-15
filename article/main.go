package main

import (
	"fmt"
	"log"

	"github.com/realBagher/hexaservice-go/article/adapters"
	"github.com/realBagher/hexaservice-go/article/core"
)

func main() {
	repo := adapters.NewInMemoryArticleRepository()
	service := core.NewArticleService(repo)

	service.CreateArticle(core.Article{
		ID:        "1",
		Title:     "Test",
		Abstract:  "Test for journal Abstract",
		AuthorID:  "1",
		JournalID: "1",
	})

	article, err := service.GetArticleByID("1")
	if err != nil {
		log.Fatalf("Error getting article: %v", err)
	}

	fmt.Println(article)
}
