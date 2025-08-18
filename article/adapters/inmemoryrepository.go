package adapters

import (
	"github.com/realBagher/hexaservice-go/article/core"
)

type InMemoryArticleRepository struct {
	articles map[string]core.Article
}

func NewInMemoryArticleRepository() *InMemoryArticleRepository {
	return &InMemoryArticleRepository{articles: make(map[string]core.Article)}
}

func (r *InMemoryArticleRepository) CreateArticle(article core.Article) (core.Article, error) {
	r.articles[article.ID] = article
	return article, nil
}

func (r *InMemoryArticleRepository) GetArticleByID(id string) (core.Article, error) {
	article, ok := r.articles[id]
	if !ok {
		return core.Article{}, core.ErrArticleNotFound
	}
	return article, nil
}

func (r *InMemoryArticleRepository) GetArticleByTitle(title string) (core.Article, error) {
	for _, article := range r.articles {
		if article.Title == title {
			return article, nil
		}
	}
	return core.Article{}, core.ErrArticleNotFound
}
