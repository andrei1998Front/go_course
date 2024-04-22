package addevent

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/andrei1998Front/go_course/homework_8/internal/domain/event"
	"github.com/andrei1998Front/go_course/homework_8/internal/inputports/http-server/handlers/addEvent/mocks"
	"github.com/andrei1998Front/go_course/homework_8/internal/lib/api/responce"
	"github.com/andrei1998Front/go_course/homework_8/internal/pkg/slogdiscard"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func TestAddHandler(t *testing.T) {
	slogDiscard := slogdiscard.NewDiscardLogger()
	eventAdderMock := mocks.NewEventAdder(t)

	tests := []struct {
		name      string
		title     string
		date      string
		mockError error
		respError string
	}{
		{
			name:  "Success",
			title: "title",
			date:  "2024-01-01",
		},
		{
			name:      "Failure. Empty title",
			date:      "2024-01-01",
			respError: "field Title as required field",
		},
		{
			name:      "Failure. Empty title",
			title:     "title",
			respError: "field Date as required field",
		},
		{
			name:      "Failure. Empty request",
			respError: "field Title as required field, field Date as required field",
		},
		{
			name:      "Failure. Existent ID",
			title:     "title",
			date:      "2024-01-01",
			respError: "event already exists",
			mockError: event.ErrExistentID,
		},
		{
			name:      "Failure. Existent ID",
			title:     "title",
			date:      "2024-01-01",
			respError: "event date is busy",
			mockError: event.ErrDateBusy,
		},
		{
			name:      "Failure. Invalid date",
			title:     "title",
			date:      "2024-012-1",
			respError: "field Date is not a valid",
		},
	}

	for _, tc := range tests {

		if tc.respError == "" || tc.mockError != nil {
			eventAdderMock.On("Handle", mock.AnythingOfType("*commands.AddEventRequest")).
				Return(tc.mockError).
				Once()
		}

		handler := New(slogDiscard, eventAdderMock)
		input := fmt.Sprintf(`{"title": "%s", "date": "%s"}`, tc.title, tc.date)

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
