package commands

import (
	"fmt"
	"testing"
	"time"

	"github.com/andrei1998Front/go_course/homework_8/internal/domain/event"
	mk "github.com/andrei1998Front/go_course/homework_8/internal/domain/event/mocks"
	"github.com/andrei1998Front/go_course/homework_8/internal/pkg/slogdiscard"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func TestAddEventHandler_Handle(t *testing.T) {
	discardLogger := slogdiscard.NewDiscardLogger()

	tests := []struct {
		Name            string
		AddEventRequest AddEventRequest
		MockError       error
	}{
		{
			Name: "Succes",
			AddEventRequest: AddEventRequest{
				Title: "sss",
				Date:  time.Now(),
			},
		},
		{
			Name: "Existent ID",
			AddEventRequest: AddEventRequest{
				Title: "existent ID",
				Date:  time.Now(),
			},
			MockError: fmt.Errorf(
				"%s: %w",
				"interfaceadapters.storage.Add",
				event.ErrExistentID,
			),
		},
		{
			Name: "Existent ID",
			AddEventRequest: AddEventRequest{
				Title: "existent ID",
				Date:  time.Now(),
			},
			MockError: fmt.Errorf(
				"%s: %w",
				"interfaceadapters.storage.Add",
				event.ErrExistentID,
			),
		},
		{
			Name: "Date busy",
			AddEventRequest: AddEventRequest{
				Title: "Date busy",
				Date:  time.Now(),
			},
			MockError: fmt.Errorf(
				"%s: %w",
				"interfaceadapters.storage.Add",
				event.ErrDateBusy,
			),
		},
	}

	for _, tc := range tests {

		mockRepo := mk.NewRepository(t)

		mockRepo.On("Add", mock.AnythingOfType("event.Event")).
			Return(tc.MockError).
			Once()

		h := addEventRequestHandler{
			log:  discardLogger,
			repo: mockRepo,
		}

		err := h.Handle(tc.AddEventRequest)

		if tc.MockError != nil {
			require.ErrorContains(t, err, tc.MockError.Error())
		}
	}
}
