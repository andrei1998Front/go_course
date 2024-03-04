package main

import (
	"fmt"
	"log"

	"github.com/andrei1998Front/go_course/homework_8/internal/domain/event"
	"github.com/andrei1998Front/go_course/homework_8/internal/interfaceadapters"
	randomeventslist "github.com/andrei1998Front/go_course/homework_8/internal/pkg/randomEventsList"
	"github.com/google/uuid"
)

func main() {
	memoryService := interfaceadapters.NewService()

	fmt.Println(memoryService)

	eventsList, err := randomeventslist.GetRandomEventsList(6, 2000, 2023)

	if err != nil {
		log.Fatal(err)
	}

	for _, e := range eventsList {
		newE := event.Event{
			ID:    uuid.New(),
			Title: e.Title,
			Date:  e.DateEvent,
		}

		memoryService.Repo.Add(newE)
	}

	fmt.Println(memoryService.Repo)
}
