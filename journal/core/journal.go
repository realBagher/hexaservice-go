package core

type Journal struct {
	ID           string  `json:"id"`
	Name         string  `json:"name"`
	Description  string  `json:"description"`
	ImpactFactor float64 `json:"impact_factor"`
}

// JournalService contains the core business logic.
type JournalService struct {
	repository JournalRepository // Port interface
}

func NewJournalService(repository JournalRepository) *JournalService {
	return &JournalService{repository: repository}
}

func (s *JournalService) CreateJournal(journal Journal) (Journal, error) {
	return s.repository.CreateJournal(journal)
}

func (s *JournalService) GetJournal(id string) (Journal, error) {
	return s.repository.GetJournal(id)
}
