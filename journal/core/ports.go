package core

type JournalRepository interface {
	CreateJournal(journal Journal) (Journal, error)
	GetJournal(id string) (Journal, error)
}
