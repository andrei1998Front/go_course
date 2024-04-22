package getallevent

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/andrei1998Front/go_course/homework_8/internal/app/event/queries"
	"github.com/andrei1998Front/go_course/homework_8/internal/inputports/http-server/handlers/getAllEvent/mocks"
	"github.com/andrei1998Front/go_course/homework_8/internal/pkg/slogdiscard"
	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
)

func TestAddHandler(t *testing.T) {
	slogDiscard := slogdiscard.NewDiscardLogger()
	eventsGetterMock := mocks.NewEventsGetter(t)
	dt, _ := time.Parse("2006-01-02", "2024-01-02")

	tests := []struct {
		name       string
		mockError  error
		mockOutput []*queries.GetAllEventResponce
		lenOutput  int
		respError  string
	}{
		{
			name: "Success. 1 event",
			mockOutput: []*queries.GetAllEventResponce{
				&queries.GetAllEventResponce{
					ID:    uuid.MustParse("3e204a57-4449-4c74-8227-77934cf25322"),
					Title: "title",
					Date:  dt,
				},
			},
			lenOutput: 1,
		},
		{
			name: "Success. 2 events",
			mockOutput: []*queries.GetAllEventResponce{
				&queries.GetAllEventResponce{
					ID:    uuid.MustParse("2e204a57-4449-4c74-8227-77934cf25322"),
					Title: "title",
					Date:  dt,
				},
				&queries.GetAllEventResponce{
					ID:    uuid.MustParse("4e204a57-4449-4c74-8227-77934cf25322"),
					Title: "title",
					Date:  dt,
				},
			},
			lenOutput: 2,
		},
		{
			name:       "Success. Zero events",
			mockOutput: []*queries.GetAllEventResponce{},
			lenOutput:  0,
		},
	}

	for _, tc := range tests {

		if tc.respError == "" || tc.mockError != nil {
			eventsGetterMock.On("Handle").
				Return(tc.mockOutput, tc.mockError).
				Once()
		}

		handler := New(slogDiscard, eventsGetterMock)
		input := ""

		req, err := http.NewRequest(http.MethodGet, "/events/", bytes.NewReader([]byte(input)))
		require.NoError(t, err)

		rr := httptest.NewRecorder()
		handler.ServeHTTP(rr, req)

		require.Equal(t, rr.Code, http.StatusOK)

		body := rr.Body.String()

		var resp Responce

		require.NoError(t, json.Unmarshal([]byte(body), &resp))
		require.Equal(t, tc.respError, resp.Error)
		require.Equal(t, tc.lenOutput, len(resp.Events))
	}
}
