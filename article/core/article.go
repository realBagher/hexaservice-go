package core

type Article struct {
	ID        string `json:"id"`
	Title     string `json:"title"`
	Abstract  string `json:"abstract"`
	AuthorID  string `json:"author_id"`
	JournalID string `json:"journal_id"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

type ArticleService struct {
	repository ArticleRepository
}

func NewArticleService(repository ArticleRepository) *ArticleService {
	return &ArticleService{repository: repository}
}

func (s *ArticleService) CreateArticle(article Article) (Article, error) {
	return s.repository.CreateArticle(article)
}

func (s *ArticleService) GetArticleByID(id string) (Article, error) {
	return s.repository.GetArticleByID(id)
}

func (s *ArticleService) GetArticleByTitle(title string) (Article, error) {
	return s.repository.GetArticleByTitle(title)
}
