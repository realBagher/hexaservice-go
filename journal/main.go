package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	"github.com/realBagher/hexaservice-go/journal/adapters"
	"github.com/realBagher/hexaservice-go/journal/core"
	"github.com/realBagher/hexaservice-go/journal/proto"
)

const (
	mysqlDSNEnvVar = "MYSQL_DSN"
	testJournalID  = "1"
	mysqlJournalID = "mysql_1"
	grpcPort       = ":50051"
)

// JournalGRPCServer implements the gRPC server interface
type JournalGRPCServer struct {
	proto.UnimplementedJournalServiceServer
	service *core.JournalService
}

// NewJournalGRPCServer creates a new gRPC server instance
func NewJournalGRPCServer(service *core.JournalService) *JournalGRPCServer {
	return &JournalGRPCServer{service: service}
}

// GetJournal implements the gRPC GetJournal method
func (s *JournalGRPCServer) GetJournal(ctx context.Context, req *proto.GetJournalRequest) (*proto.GetJournalResponse, error) {
	journal, err := s.service.GetJournal(req.Id)
	if err != nil {
		return nil, err
	}

	// Convert core.Journal to proto.Journal
	protoJournal := &proto.Journal{
		Id:           journal.ID,
		Name:         journal.Name,
		Description:  journal.Description,
		ImpactFactor: journal.ImpactFactor,
	}

	return &proto.GetJournalResponse{Journal: protoJournal}, nil
}

func main() {
	// Start gRPC server in a separate goroutine
	go func() {
		if err := startGRPCServer(); err != nil {
			log.Fatalf("gRPC server failed: %v", err)
		}
	}()

	// Run demo
	if err := runDemo(); err != nil {
		log.Fatalf("Demo failed: %v", err)
	}

	// Keep the main goroutine alive
	select {}
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

func startGRPCServer() error {
	// Create a repository and service for the gRPC server
	// Using in-memory repository for simplicity, but could be MySQL based on env var
	var repo core.JournalRepository
	if dsn := os.Getenv(mysqlDSNEnvVar); dsn != "" {
		db, err := adapters.NewMySQLConnection(dsn)
		if err != nil {
			log.Printf("Failed to connect to MySQL, falling back to in-memory: %v", err)
			repo = adapters.NewInMemoryJournalRepository()
		} else {
			mysqlRepo := adapters.NewMySQLJournalRepository(db)
			if err := mysqlRepo.InitializeSchema(); err != nil {
				log.Printf("Failed to initialize MySQL schema, falling back to in-memory: %v", err)
				repo = adapters.NewInMemoryJournalRepository()
			} else {
				repo = mysqlRepo
			}
		}
	} else {
		repo = adapters.NewInMemoryJournalRepository()
	}

	service := core.NewJournalService(repo)

	// Pre-populate with a test journal for the article service to find
	testJournal := createTestJournal("journal_1")
	if _, err := service.CreateJournal(testJournal); err != nil {
		log.Printf("Warning: Failed to create test journal: %v", err)
	}

	// Create gRPC server
	grpcServer := grpc.NewServer()
	journalGRPCServer := NewJournalGRPCServer(service)

	proto.RegisterJournalServiceServer(grpcServer, journalGRPCServer)
	reflection.Register(grpcServer)

	// Start listening
	listener, err := net.Listen("tcp", grpcPort)
	if err != nil {
		return fmt.Errorf("failed to listen on port %s: %w", grpcPort, err)
	}

	log.Printf("gRPC server starting on port %s", grpcPort)
	return grpcServer.Serve(listener)
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
