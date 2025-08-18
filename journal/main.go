package main

import (
	"fmt"
	"log"
	"os"

	"github.com/realBagher/hexaservice-go/journal/adapters"
	"github.com/realBagher/hexaservice-go/journal/core"
)

const (
	mysqlDSNEnvVar = "MYSQL_DSN"
	testJournalID  = "1"
	mysqlJournalID = "mysql_1"
)

func main() {
	if err := runDemo(); err != nil {
		log.Fatalf("Demo failed: %v", err)
	}
}

func runDemo() error {
	// Demonstrate InMemory repository
	if err := demonstrateInMemoryRepository(); err != nil {
		return fmt.Errorf("in-memory repository demo failed: %w", err)
	}

	// Demonstrate MySQL repository if DSN is available
	if dsn := os.Getenv(mysqlDSNEnvVar); dsn != "" {
		if err := demonstrateMySQLRepository(dsn); err != nil {
			return fmt.Errorf("MySQL repository demo failed: %w", err)
		}
	}

	return nil
}

func demonstrateInMemoryRepository() error {
	fmt.Println("=== Using InMemory Repository ===")

	repo := adapters.NewInMemoryJournalRepository()
	service := core.NewJournalService(repo)

	testJournal := createTestJournal(testJournalID)

	return demonstrateJournalOperations(service, testJournal)
}

func demonstrateMySQLRepository(dsn string) error {
	fmt.Println("\n=== Live MySQL Repository Demo ===")

	db, err := adapters.NewMySQLConnection(dsn)
	if err != nil {
		return fmt.Errorf("failed to connect to MySQL: %w", err)
	}
	defer func() {
		if closeErr := db.Close(); closeErr != nil {
			log.Printf("Warning: failed to close database connection: %v", closeErr)
		}
	}()

	repo := adapters.NewMySQLJournalRepository(db)
	if err := repo.InitializeSchema(); err != nil {
		return fmt.Errorf("failed to initialize schema: %w", err)
	}

	service := core.NewJournalService(repo)
	testJournal := createTestJournal(mysqlJournalID)

	return demonstrateJournalOperations(service, testJournal)
}

func createTestJournal(id string) core.Journal {
	return core.Journal{
		ID:           id,
		Name:         "Nature",
		Description:  "Leading scientific journal",
		ImpactFactor: 64.8,
	}
}

func demonstrateJournalOperations(service *core.JournalService, journal core.Journal) error {
	// Create journal
	createdJournal, err := service.CreateJournal(journal)
	if err != nil {
		return fmt.Errorf("failed to create journal: %w", err)
	}
	fmt.Printf("Created journal: %+v\n", createdJournal)

	// Retrieve journal
	retrievedJournal, err := service.GetJournal(journal.ID)
	if err != nil {
		return fmt.Errorf("failed to retrieve journal: %w", err)
	}
	fmt.Printf("Retrieved journal: %+v\n", retrievedJournal)

	return nil
}
