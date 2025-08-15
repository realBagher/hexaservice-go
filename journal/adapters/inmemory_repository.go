package adapters

import (
	"errors"

	"github.com/realBagher/hexaservice-go/journal/core"
)

type InMemoryJournalRepository struct {
	journals map[string]core.Journal
}

func NewInMemoryJournalRepository() *InMemoryJournalRepository {
	return &InMemoryJournalRepository{journals: make(map[string]core.Journal)}
}

func (r *InMemoryJournalRepository) CreateJournal(journal core.Journal) (core.Journal, error) {
	r.journals[journal.ID] = journal
	return journal, nil
}

func (r *InMemoryJournalRepository) GetJournal(id string) (core.Journal, error) {
	journal, ok := r.journals[id]
	if !ok {
		return core.Journal{}, errors.New("journal not found")
	}
	return journal, nil
}
