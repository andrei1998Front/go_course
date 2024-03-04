package interfaceadapters

import (
	"github.com/andrei1998Front/go_course/homework_8/internal/domain/event"
	memory "github.com/andrei1998Front/go_course/homework_8/internal/interfaceadapters/storage"
)

type Service struct {
	Repo event.Repository
}

func NewService() Service {
	return Service{
		Repo: memory.NewRepo(),
	}
}
