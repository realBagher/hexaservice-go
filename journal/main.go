package main

import (
	"fmt"
	"log"

	"github.com/realBagher/hexaservice-go/journal/adapters"
	"github.com/realBagher/hexaservice-go/journal/core"
)

func main() {
	repo := adapters.NewInMemoryJournalRepository()
	service := core.NewJournalService(repo)

	service.CreateJournal(core.Journal{
		ID:           "1",
		Name:         "Test",
		Description:  "Test2",
		ImpactFactor: 1.0,
	})

	journal, err := service.GetJournal("Test")
	if err != nil {
		log.Fatalf("Error getting journal: %v", err)
	}

	fmt.Println(journal)

}
