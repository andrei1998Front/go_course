package queries

import (
	"fmt"
	"testing"
	"time"

	"github.com/andrei1998Front/go_course/homework_8/internal/domain/event"
	mk "github.com/andrei1998Front/go_course/homework_8/internal/domain/event/mocks"
	"github.com/andrei1998Front/go_course/homework_8/internal/pkg/slogdiscard"
	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
)

func TestGetEventQueryHandler_Handle(t *testing.T) {
	discardLogger := slogdiscard.NewDiscardLogger()

	idMocks := []uuid.UUID{
		uuid.MustParse("3e204a57-4449-4c74-8227-77934cf25322"),
		uuid.MustParse("3e234a57-4449-4c74-8227-77937cf25322"),
	}

	tests := []struct {
		Name            string
		GetEventRequest GetEventRequest
		mockError       error
		mockOutput      *event.Event
	}{
		{
			Name: "Success",
			GetEventRequest: GetEventRequest{
				EventID: idMocks[0],
			},
			mockOutput: &event.Event{
				ID:    idMocks[0],
				Title: "ddddd",
				Date:  time.Now(),
			},
		},
		{
			Name: "Non-existent event",
			GetEventRequest: GetEventRequest{
				EventID: idMocks[1],
			},
			mockOutput: &event.Event{},
			mockError:  fmt.Errorf("%s: %w", "interfaceadapters.storage.GetByID", event.ErrNonExistentEvent),
		},
	}

	for _, tc := range tests {
		mockRepo := mk.NewRepository(t)
		mockRepo.On("GetByID", tc.GetEventRequest.EventID).
			Return(tc.mockOutput, tc.mockError).
			Once()

		h := getEventRequestHandler{
			log:  discardLogger,
			repo: mockRepo,
		}

		got, err := h.Handle(tc.GetEventRequest)

		if err != nil {
			require.ErrorContains(t, err, tc.mockError.Error())
			continue
		}

		require.Equal(t, tc.mockOutput.ID, got.ID)
		require.Equal(t, tc.mockOutput.Title, got.Title)
		require.Equal(t, tc.mockOutput.Date, got.Date)
	}
}
