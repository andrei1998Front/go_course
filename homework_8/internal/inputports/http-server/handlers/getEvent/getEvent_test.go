package getevent

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/andrei1998Front/go_course/homework_8/internal/app/event/queries"
	"github.com/andrei1998Front/go_course/homework_8/internal/domain/event"
	"github.com/andrei1998Front/go_course/homework_8/internal/inputports/http-server/handlers/getEvent/mocks"
	"github.com/andrei1998Front/go_course/homework_8/internal/pkg/slogdiscard"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func TestDeleteHandler(t *testing.T) {
	slogDiscard := slogdiscard.NewDiscardLogger()
	eventGetterMock := mocks.NewEventGetter(t)

	dt, _ := time.Parse("2006-01-02", "2024-01-02")

	tests := []struct {
		name       string
		id         string
		mockError  error
		mockOutput *queries.GetEventResponce
		respError  string
		respStatus int
	}{
		{
			name: "Success",
			id:   "3e204a57-4449-4c74-8227-77934cf25322",
			mockOutput: &queries.GetEventResponce{
				ID:    uuid.MustParse("3e204a57-4449-4c74-8227-77934cf25322"),
				Title: "title",
				Date:  dt,
			},
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
			respError:  "invalid event id",
			respStatus: http.StatusOK,
		},
		{
			name:       "Failure. Non-existent event",
			id:         "3e204a57-4449-4c74-8227-77934cf25322",
			respError:  "non existent event",
			mockError:  event.ErrNonExistentEvent,
			mockOutput: &queries.GetEventResponce{},
			respStatus: http.StatusOK,
		},
	}

	for _, tc := range tests {
		t.Log(tc.name)
		if tc.respError == "" || tc.mockError != nil {
			eventGetterMock.On("Handle", mock.AnythingOfType("queries.GetEventRequest")).
				Return(tc.mockOutput, tc.mockError).
				Once()
		}

		handler := New(slogDiscard, eventGetterMock)

		r := chi.NewRouter()
		r.Use(middleware.URLFormat)
		r.Get("/events/{event_id}", handler)

		ts := httptest.NewServer(r)
		defer ts.Close()

		cl := ts.Client()

		res, err := cl.Get(ts.URL + "/events/" + tc.id)

		require.NoError(t, err)

		require.Equal(t, tc.respStatus, res.StatusCode)

		if res.StatusCode != http.StatusNotFound {
			var resp Responce

			err = json.NewDecoder(res.Body).Decode(&resp)
			require.NoError(t, err)

			require.Equal(t, tc.respError, resp.Error)
		}
	}
}
