package core

import (
	"fmt"
	"strings"
)

type Journal struct {
	ID           string  `json:"id"`
	Name         string  `json:"name"`
	Description  string  `json:"description"`
	ImpactFactor float64 `json:"impact_factor"`
}

// Validate checks if the journal data is valid
func (j Journal) Validate() error {
	if strings.TrimSpace(j.ID) == "" {
		return fmt.Errorf("%w: ID cannot be empty", ErrInvalidJournal)
	}

	if strings.TrimSpace(j.Name) == "" {
		return fmt.Errorf("%w: name cannot be empty", ErrInvalidJournal)
	}

	if j.ImpactFactor < 0 {
		return fmt.Errorf("%w: impact factor cannot be negative", ErrInvalidJournal)
	}

	return nil
}

// JournalService contains the core business logic.
type JournalService struct {
	repository JournalRepository // Port interface
}

func NewJournalService(repository JournalRepository) *JournalService {
	return &JournalService{repository: repository}
}

func (s *JournalService) CreateJournal(journal Journal) (Journal, error) {
	if err := journal.Validate(); err != nil {
		return Journal{}, err
	}
	return s.repository.CreateJournal(journal)
}

func (s *JournalService) GetJournal(id string) (Journal, error) {
	return s.repository.GetJournal(id)
}
