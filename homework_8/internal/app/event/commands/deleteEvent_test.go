package commands

import (
	"fmt"
	"testing"

	"github.com/andrei1998Front/go_course/homework_8/internal/domain/event"
	mk "github.com/andrei1998Front/go_course/homework_8/internal/domain/event/mocks"
	"github.com/andrei1998Front/go_course/homework_8/internal/pkg/slogdiscard"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func TestDeleteEventHandler_Handle(t *testing.T) {
	discardLogger := slogdiscard.NewDiscardLogger()

	idMocks := []string{
		"3e204a57-4449-4c74-8227-77934cf25322",
		"3e234a57-4449-4c74-8227-77937cf25322",
		"5e634a57-4449-4c74-8227-77937cf25322",
		"ff",
	}

	tests := []struct {
		Name               string
		DeleteEventRequest DeleteEventRequest
		MockError          error
		WantError          bool
		Err                string
		ExcludeMock        bool
	}{
		{
			Name: "Succes",
			DeleteEventRequest: DeleteEventRequest{
				ID: idMocks[0],
			},
		},
		{
			Name: "non-existent ID",
			DeleteEventRequest: DeleteEventRequest{
				ID: idMocks[1],
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
			Name: "Invalid ID",
			DeleteEventRequest: DeleteEventRequest{
				ID: idMocks[3],
			},
			WantError:   true,
			ExcludeMock: true,
			Err:         ErrInvalidUUID.Error(),
		},
	}

	for _, tc := range tests {

		mockRepo := mk.NewRepository(t)
		if !tc.ExcludeMock {
			mockRepo.On("Delete", mock.AnythingOfType("uuid.UUID")).
				Return(tc.MockError).
				Once()
		}

		h := deleteEventRequestHandler{
			log:  discardLogger,
			repo: mockRepo,
		}

		err := h.Handle(tc.DeleteEventRequest)

		if tc.WantError {
			require.ErrorContains(t, err, tc.Err)
		} else if !tc.WantError && err != nil {
			t.Errorf("unexpected error: " + err.Error())
			return
		}
	}
}
