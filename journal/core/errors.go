package core

import "errors"

var (
	// ErrJournalNotFound is returned when a journal is not found
	ErrJournalNotFound = errors.New("journal not found")

	// ErrInvalidJournal is returned when journal data is invalid
	ErrInvalidJournal = errors.New("invalid journal data")
)
