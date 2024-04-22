package updateevent

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/andrei1998Front/go_course/homework_8/internal/app/event/commands"
	"github.com/andrei1998Front/go_course/homework_8/internal/domain/event"
	"github.com/andrei1998Front/go_course/homework_8/internal/inputports/http-server/handlers/updateEvent/mocks"
	"github.com/andrei1998Front/go_course/homework_8/internal/lib/api/responce"
	"github.com/andrei1998Front/go_course/homework_8/internal/pkg/slogdiscard"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func TestUpdateHandler(t *testing.T) {
	slogDiscard := slogdiscard.NewDiscardLogger()
	eventUpdaterMock := mocks.NewEventUpdater(t)

	tests := []struct {
		name       string
		id         string
		title      string
		date       string
		mockError  error
		respError  string
		respStatus int
	}{
		{
			name:       "Success",
			id:         "3e204a57-4449-4c74-8227-77934cf25322",
			title:      "title",
			date:       "2004-01-02",
			respStatus: http.StatusOK,
		},
		{
			name:       "Failure. Empty id",
			respError:  "invalid event id",
			respStatus: http.StatusNotFound,
		},
		{
			name:       "Failure. Invalid id",
			id:         "3e",
			title:      "title",
			date:       "2004-01-02",
			respError:  "invalid event id",
			respStatus: http.StatusOK,
			mockError:  commands.ErrInvalidUUID,
		},
		{
			name:       "Failure. Invalid id",
			id:         "3e",
			title:      "title",
			date:       "2004-01-02",
			respError:  "invalid event id",
			respStatus: http.StatusOK,
			mockError:  commands.ErrInvalidUUID,
		},
		{
			name:       "Failure. Invalid date",
			id:         "3e",
			title:      "title",
			date:       "20041-02",
			respError:  "field Date is not a valid",
			respStatus: http.StatusOK,
		},
		{
			name:       "Failure. Invalid date",
			id:         "3e204a57-4449-4c74-8227-77934cf25322",
			title:      "title",
			date:       "2004-01-02",
			respError:  "event date is busy",
			respStatus: http.StatusOK,
			mockError:  event.ErrDateBusy,
		},
		{
			name:       "Failure. Empty query",
			id:         "3e204a57-4449-4c74-8227-77934cf25322",
			respError:  "failed to update event",
			respStatus: http.StatusOK,
			mockError:  commands.ErrEmptyQuery,
		},
	}

	for _, tc := range tests {
		t.Log(tc.name)
		if tc.respError == "" || tc.mockError != nil {
			eventUpdaterMock.On("Handle", mock.AnythingOfType("commands.UpdateEventRequest")).
				Return(tc.mockError).
				Once()
		}

		handler := New(slogDiscard, eventUpdaterMock)

		r := chi.NewRouter()
		r.Use(middleware.URLFormat)
		r.Patch("/events/{event_id}", handler)

		ts := httptest.NewServer(r)
		defer ts.Close()

		input := fmt.Sprintf(`{"title": "%s", "date":"%s"}`, tc.title, tc.date)

		req, err := http.NewRequest(http.MethodPatch, ts.URL+"/events/"+tc.id, bytes.NewReader([]byte(input)))
		require.NoError(t, err)

		cl := ts.Client()

		res, err := cl.Do(req)
		require.NoError(t, err)

		if res.StatusCode != http.StatusNotFound {
			var resp responce.Responce

			err = json.NewDecoder(res.Body).Decode(&resp)
			require.NoError(t, err)

			require.Equal(t, tc.respError, resp.Error)
		}
	}
}
