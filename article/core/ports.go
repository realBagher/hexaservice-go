package core

type ArticleRepository interface {
	CreateArticle(article Article) (Article, error)
	GetArticleByID(id string) (Article, error)
	GetArticleByTitle(title string) (Article, error)
}
