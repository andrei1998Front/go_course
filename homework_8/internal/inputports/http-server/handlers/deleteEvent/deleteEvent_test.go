package deleteevent

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/andrei1998Front/go_course/homework_8/internal/domain/event"
	"github.com/andrei1998Front/go_course/homework_8/internal/inputports/http-server/handlers/deleteEvent/mocks"
	"github.com/andrei1998Front/go_course/homework_8/internal/lib/api/responce"
	"github.com/andrei1998Front/go_course/homework_8/internal/pkg/slogdiscard"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func TestDeleteHandler(t *testing.T) {
	slogDiscard := slogdiscard.NewDiscardLogger()
	eventDeleterMock := mocks.NewEventDeleter(t)

	tests := []struct {
		name      string
		id        string
		mockError error
		respError string
	}{
		{
			name: "Success",
			id:   "3e204a57-4449-4c74-8227-77934cf25322",
		},
		{
			name:      "Failure. Empty id",
			respError: "field ID as required field",
		},
		{
			name:      "Failure. Invalid id",
			id:        "3e",
			respError: "field ID is not a valid",
		},
		{
			name:      "Failure. Non-existent event",
			id:        "3e204a57-4449-4c74-8227-77934cf25322",
			respError: "non existent id",
			mockError: event.ErrNonExistentEvent,
		},
	}

	for _, tc := range tests {

		if tc.respError == "" || tc.mockError != nil {
			eventDeleterMock.On("Handle", mock.AnythingOfType("commands.DeleteEventRequest")).
				Return(tc.mockError).
				Once()
		}

		handler := New(slogDiscard, eventDeleterMock)
		input := fmt.Sprintf(`{"id": "%s"}`, tc.id)

		req, err := http.NewRequest(http.MethodPost, "/events/", bytes.NewReader([]byte(input)))
		require.NoError(t, err)

		rr := httptest.NewRecorder()
		handler.ServeHTTP(rr, req)

		require.Equal(t, rr.Code, http.StatusOK)

		body := rr.Body.String()

		var resp responce.Responce

		require.NoError(t, json.Unmarshal([]byte(body), &resp))
		require.Equal(t, tc.respError, resp.Error)
	}
}
