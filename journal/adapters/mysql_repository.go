package adapters

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/realBagher/hexaservice-go/journal/core"
)

type MySQLJournalRepository struct {
	db *sql.DB
}

func NewMySQLJournalRepository(db *sql.DB) *MySQLJournalRepository {
	return &MySQLJournalRepository{db: db}
}

// NewMySQLConnection creates a new MySQL database connection
func NewMySQLConnection(dsn string) (*sql.DB, error) {
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, fmt.Errorf("failed to open database connection: %w", err)
	}

	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("failed to ping database: %w", err)
	}

	return db, nil
}

// InitializeSchema creates the journals table if it doesn't exist
func (r *MySQLJournalRepository) InitializeSchema() error {
	query := `
	CREATE TABLE IF NOT EXISTS journals (
		id VARCHAR(255) PRIMARY KEY,
		name VARCHAR(255) NOT NULL,
		description TEXT,
		impact_factor DECIMAL(10,3)
	)`

	_, err := r.db.Exec(query)
	if err != nil {
		return fmt.Errorf("failed to create journals table: %w", err)
	}

	return nil
}

func (r *MySQLJournalRepository) CreateJournal(journal core.Journal) (core.Journal, error) {
	query := `
	INSERT INTO journals (id, name, description, impact_factor) 
	VALUES (?, ?, ?, ?)`

	_, err := r.db.Exec(query, journal.ID, journal.Name, journal.Description, journal.ImpactFactor)
	if err != nil {
		return core.Journal{}, fmt.Errorf("failed to create journal: %w", err)
	}

	return journal, nil
}

func (r *MySQLJournalRepository) GetJournal(id string) (core.Journal, error) {
	var journal core.Journal
	query := `
	SELECT id, name, description, impact_factor 
	FROM journals 
	WHERE id = ?`

	row := r.db.QueryRow(query, id)
	err := row.Scan(&journal.ID, &journal.Name, &journal.Description, &journal.ImpactFactor)
	if err != nil {
		if err == sql.ErrNoRows {
			return core.Journal{}, core.ErrJournalNotFound
		}
		return core.Journal{}, fmt.Errorf("failed to get journal: %w", err)
	}

	return journal, nil
}

// Close closes the database connection
func (r *MySQLJournalRepository) Close() error {
	return r.db.Close()
}
