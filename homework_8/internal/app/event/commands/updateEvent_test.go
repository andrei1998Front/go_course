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

func TestUpdateEventHandler_Handle(t *testing.T) {
	discardLogger := slogdiscard.NewDiscardLogger()

	idMocks := []string{
		"3e204a57-4449-4c74-8227-77934cf25322",
		"3e234a57-4449-4c74-8227-77937cf25322",
		"5e634a57-4449-4c74-8227-77937cf25322",
		"ff",
	}

	tests := []struct {
		Name               string
		UpdateEventRequest UpdateEventRequest
		MockError          error
		WantError          bool
		Err                string
		ExcludeMock        bool
	}{
		{
			Name: "Succes",
			UpdateEventRequest: UpdateEventRequest{
				ID:    idMocks[0],
				Title: "sss",
				Date:  time.Now(),
			},
		},
		{
			Name: "non-existent ID",
			UpdateEventRequest: UpdateEventRequest{
				ID:    idMocks[1],
				Title: "non-existent ID",
				Date:  time.Now(),
			},
			WantError: true,
			MockError: fmt.Errorf(
				"%s: %w",
				"interfaceadapters.storage.GetByID",
				event.ErrNonExistentID,
			),
			Err: event.ErrNonExistentID.Error(),
		},
		{
			Name: "Date busy",
			UpdateEventRequest: UpdateEventRequest{
				ID:    idMocks[2],
				Title: "Date busy",
				Date:  time.Now(),
			},
			WantError: true,
			MockError: fmt.Errorf(
				"%s: %w",
				"interfaceadapters.storage.Update",
				event.ErrDateBusy,
			),
			Err: event.ErrDateBusy.Error(),
		},
		{
			Name: "Invalid ID",
			UpdateEventRequest: UpdateEventRequest{
				ID:    idMocks[3],
				Title: "Invalid ID",
				Date:  time.Now(),
			},
			WantError:   true,
			ExcludeMock: true,
			Err:         ErrInvalidUUID.Error(),
		},
	}

	for _, tc := range tests {

		mockRepo := mk.NewRepository(t)
		if !tc.ExcludeMock {
			mockRepo.On("Update", mock.AnythingOfType("event.Event")).
				Return(tc.MockError).
				Once()
		}

		h := updateEventRequestHandler{
			log:  discardLogger,
			repo: mockRepo,
		}

		err := h.Handle(tc.UpdateEventRequest)

		if tc.WantError {
			require.ErrorContains(t, err, tc.Err)
		} else if !tc.WantError && err != nil {
			t.Errorf("unexpected error: " + err.Error())
			return
		}
	}
}
