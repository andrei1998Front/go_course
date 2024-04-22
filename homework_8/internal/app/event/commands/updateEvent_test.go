package commands

import (
	"testing"
	"time"

	"github.com/andrei1998Front/go_course/homework_8/internal/domain/event"
	"github.com/andrei1998Front/go_course/homework_8/internal/domain/event/mocks"
	"github.com/andrei1998Front/go_course/homework_8/internal/pkg/slogdiscard"
	"github.com/google/uuid"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

const dtFormat = "2006-01-02"

func TestSetEmptyEventFields(t *testing.T) {
	dt, _ := time.Parse(dtFormat, "2024-01-01")

	tests := []struct {
		name    string
		ev      *event.Event
		evByID  *event.Event
		req     *UpdateEventRequest
		title   string
		date    time.Time
		wantErr bool
		err     error
	}{
		{
			name:   "Title and Date in request",
			ev:     &event.Event{},
			evByID: &event.Event{},
			req: &UpdateEventRequest{
				Title: "tt",
				Date:  "2024-01-01",
			},
			title: "tt",
			date:  dt,
		},
		{
			name: "Title and Date in storage",
			ev:   &event.Event{},
			evByID: &event.Event{
				Title: "tt",
				Date:  dt,
			},
			req:   &UpdateEventRequest{},
			title: "tt",
			date:  dt,
		},
		{
			name: "Title in request, date in storage",
			ev:   &event.Event{},
			evByID: &event.Event{
				Title: "tt",
			},
			req:   &UpdateEventRequest{Date: "2024-01-01"},
			title: "tt",
			date:  dt,
		},
		{
			name:   "Title in request, date in storage",
			ev:     &event.Event{},
			evByID: &event.Event{Date: dt},
			req:    &UpdateEventRequest{Title: "tt"},
			title:  "tt",
			date:   dt,
		},
		{
			name:    "Invalid date in request",
			ev:      &event.Event{},
			evByID:  &event.Event{Date: dt},
			req:     &UpdateEventRequest{Title: "tt", Date: "dd"},
			wantErr: true,
			err:     ErrInvalidDate,
		},
		{
			name:    "empty title",
			ev:      &event.Event{},
			evByID:  &event.Event{Date: dt},
			req:     &UpdateEventRequest{Date: "dd"},
			wantErr: true,
			err:     ErrInvalidTitle,
		},
		{
			name:    "empty date",
			ev:      &event.Event{},
			evByID:  &event.Event{Title: "tt"},
			req:     &UpdateEventRequest{Title: "tt"},
			wantErr: true,
			err:     ErrInvalidDate,
		},
	}

	for _, tc := range tests {
		err := setEmptyEventFields(tc.ev, tc.evByID, tc.req)

		if tc.wantErr && err != nil {
			require.ErrorIs(t, err, tc.err)
			continue
		} else if !tc.wantErr && err != nil {
			require.Fail(t, "unxpected error", err)
		} else if tc.wantErr && err == nil {
			require.Fail(t, "expected error missing", tc.err)
		}

		require.Equal(t, tc.title, tc.ev.Title)
		require.Equal(t, tc.date, tc.date)
	}
}

func TestUpdateEventHandler_setEmtyEventFields(t *testing.T) {
	mockRepo := mocks.NewRepository(t)
	slogDiscard := slogdiscard.NewDiscardLogger()

	dt, _ := time.Parse(dtFormat, "2024-01-01")

	tests := []struct {
		name        string
		req         *UpdateEventRequest
		wantErr     bool
		excludeMock bool
		mockError   error
		mockOutput  *event.Event
		eventOutput event.Event
		err         error
	}{
		{
			name: "Success",
			req: &UpdateEventRequest{
				ID:    "3e204a57-4449-4c74-8227-77934cf25322",
				Title: "title_1",
				Date:  "2024-01-01",
			},
			eventOutput: event.Event{
				ID:    uuid.MustParse("3e204a57-4449-4c74-8227-77934cf25322"),
				Title: "title_1",
				Date:  dt,
			},
			excludeMock: true,
		},
		{
			name: "Success. Date in storage",
			req: &UpdateEventRequest{
				ID:    "3e204a57-4449-4c74-8227-77934cf25322",
				Title: "title_1",
			},
			eventOutput: event.Event{
				ID:    uuid.MustParse("3e204a57-4449-4c74-8227-77934cf25322"),
				Title: "title_1",
				Date:  dt,
			},
			mockOutput: &event.Event{Date: dt},
		},
		{
			name: "Success. Title in storage",
			req: &UpdateEventRequest{
				ID: "3e204a57-4449-4c74-8227-77934cf25322",
			},
			eventOutput: event.Event{
				ID:    uuid.MustParse("3e204a57-4449-4c74-8227-77934cf25322"),
				Title: "title_1",
				Date:  dt,
			},
			mockOutput: &event.Event{Title: "title_1", Date: dt},
		},
		{
			name: "Failure. Empty date",
			req: &UpdateEventRequest{
				ID:    "3e204a57-4449-4c74-8227-77934cf25322",
				Title: "title_1",
			},
			eventOutput: event.Event{
				ID:    uuid.MustParse("3e204a57-4449-4c74-8227-77934cf25322"),
				Title: "title_1",
			},
			mockOutput: &event.Event{},
			err:        ErrInvalidDate,
			wantErr:    true,
		},
		{
			name: "Failure. Empty title",
			req: &UpdateEventRequest{
				ID: "3e204a57-4449-4c74-8227-77934cf25322",
			},
			eventOutput: event.Event{
				ID: uuid.MustParse("3e204a57-4449-4c74-8227-77934cf25322"),
			},
			mockOutput: &event.Event{},
			err:        ErrInvalidTitle,
			wantErr:    true,
		},
		{
			name: "Failure. Invalid ID",
			req: &UpdateEventRequest{
				ID: "g",
			},
			eventOutput: event.Event{},
			err:         ErrInvalidUUID,
			wantErr:     true,
			excludeMock: true,
		},
		{
			name: "Failure. Invalid date",
			req: &UpdateEventRequest{
				ID:    "3e204a57-4449-4c74-8227-77934cf25322",
				Title: "title_1",
				Date:  "111",
			},
			eventOutput: event.Event{
				ID:    uuid.MustParse("3e204a57-4449-4c74-8227-77934cf25322"),
				Title: "title_1",
			},
			err:         ErrInvalidDate,
			wantErr:     true,
			excludeMock: true,
		},
		{
			name: "Failure. Non-existent event",
			req: &UpdateEventRequest{
				ID:   "3e204a57-4449-4c74-8227-77934cf25322",
				Date: "111",
			},
			eventOutput: event.Event{
				ID: uuid.MustParse("3e204a57-4449-4c74-8227-77934cf25322"),
			},
			err:         event.ErrNonExistentEvent,
			wantErr:     true,
			mockOutput:  &event.Event{},
			mockError:   event.ErrNonExistentEvent,
			excludeMock: false,
		},
	}

	for _, tc := range tests {
		t.Log(tc.name)

		if !tc.excludeMock {
			mockRepo.On("GetByID", mock.AnythingOfType("uuid.UUID")).
				Return(tc.mockOutput, tc.mockError).
				Once()
		}

		h := NewUpdateEventRequestHandler(slogDiscard, mockRepo)

		ev, err := h.setupUpdatableEvent(tc.req)

		if tc.wantErr && err != nil {
			require.ErrorIs(t, err, tc.err)
		} else if !tc.wantErr && err != nil {
			require.Fail(t, "unxpected error", err)
		} else if tc.wantErr && err == nil {
			require.Fail(t, "expected error missing", tc.err)
		}

		require.Equal(t, tc.eventOutput.ID, ev.ID)
		require.Equal(t, tc.eventOutput.Title, ev.Title)
		require.Equal(t, tc.eventOutput.Date, ev.Date)
	}
}

func TestUpdateEventHandler_Handle(t *testing.T) {
	mockRepo := mocks.NewRepository(t)
	slogDiscard := slogdiscard.NewDiscardLogger()

	dt, _ := time.Parse(dtFormat, "2024-01-01")

	tests := []struct {
		name              string
		req               *UpdateEventRequest
		wantErr           bool
		excludeGetMock    bool
		mockGetError      error
		mockGetOutput     *event.Event
		excludeUpdateMock bool
		mockUpdateError   error
		err               error
	}{
		{
			name: "Success",
			req: &UpdateEventRequest{
				ID:    "3e204a57-4449-4c74-8227-77934cf25322",
				Title: "title_1",
				Date:  "2024-01-01",
			},
			excludeGetMock: true,
		},
		{
			name: "Success. Date in storage",
			req: &UpdateEventRequest{
				ID:    "3e204a57-4449-4c74-8227-77934cf25322",
				Title: "title_1",
			},
			mockGetOutput: &event.Event{Date: dt},
		},
		{
			name: "Success. Title in storage",
			req: &UpdateEventRequest{
				ID:   "3e204a57-4449-4c74-8227-77934cf25322",
				Date: "2004-01-01",
			},
			mockGetOutput: &event.Event{Title: "title_1", Date: dt},
		},
		{
			name: "Failure. Empty date",
			req: &UpdateEventRequest{
				ID:    "3e204a57-4449-4c74-8227-77934cf25322",
				Title: "title_1",
			},
			mockGetOutput:     &event.Event{},
			excludeUpdateMock: true,
			err:               ErrInvalidDate,
			wantErr:           true,
		},
		{
			name: "Failure. Empty title",
			req: &UpdateEventRequest{
				ID:   "3e204a57-4449-4c74-8227-77934cf25322",
				Date: "2024-01-02",
			},
			mockGetOutput:     &event.Event{},
			excludeUpdateMock: true,
			err:               ErrInvalidTitle,
			wantErr:           true,
		},
		{
			name: "Failure. Invalid ID",
			req: &UpdateEventRequest{
				ID:   "g",
				Date: "2024-01-02",
			},
			excludeUpdateMock: true,
			err:               ErrInvalidUUID,
			wantErr:           true,
			excludeGetMock:    true,
		},
		{
			name: "Failure. Invalid date",
			req: &UpdateEventRequest{
				ID:    "3e204a57-4449-4c74-8227-77934cf25322",
				Title: "title_1",
				Date:  "111",
			},
			err:               ErrInvalidDate,
			wantErr:           true,
			excludeGetMock:    true,
			excludeUpdateMock: true,
		},
		{
			name: "Failure. Empty request",
			req: &UpdateEventRequest{
				ID: "3e204a57-4449-4c74-8227-77934cf25322",
			},
			err:               ErrEmptyQuery,
			wantErr:           true,
			excludeGetMock:    true,
			excludeUpdateMock: true,
		},
		{
			name: "Failure. Non-existent event",
			req: &UpdateEventRequest{
				ID:   "3e204a57-4449-4c74-8227-77934cf25322",
				Date: "2024-01-01",
			},
			err:               event.ErrNonExistentEvent,
			wantErr:           true,
			mockGetOutput:     &event.Event{},
			mockGetError:      event.ErrNonExistentEvent,
			excludeUpdateMock: true,
		},
		{
			name: "Failure. Non-existent event",
			req: &UpdateEventRequest{
				ID:    "3e204a57-4449-4c74-8227-77934cf25322",
				Title: "title_1",
				Date:  "2024-01-01",
			},
			err:               event.ErrNonExistentEvent,
			wantErr:           true,
			excludeGetMock:    true,
			excludeUpdateMock: false,
			mockUpdateError:   event.ErrNonExistentEvent,
		},
		{
			name: "Failure. Date busy",
			req: &UpdateEventRequest{
				ID:    "3e204a57-4449-4c74-8227-77934cf25322",
				Title: "title_1",
				Date:  "2024-01-01",
			},
			err:               event.ErrDateBusy,
			wantErr:           true,
			excludeGetMock:    true,
			excludeUpdateMock: false,
			mockUpdateError:   event.ErrDateBusy,
		},
	}

	for _, tc := range tests {
		t.Log(tc.name)

		if !tc.excludeGetMock {
			mockRepo.On("GetByID", mock.AnythingOfType("uuid.UUID")).
				Return(tc.mockGetOutput, tc.mockGetError).
				Once()
		}

		if !tc.excludeUpdateMock {
			mockRepo.On("Update", mock.AnythingOfType("event.Event")).
				Return(tc.mockUpdateError).
				Once()
		}

		h := NewUpdateEventRequestHandler(slogDiscard, mockRepo)

		err := h.Handle(*tc.req)

		if tc.wantErr && err != nil {
			require.ErrorIs(t, err, tc.err)
		} else if !tc.wantErr && err != nil {
			require.Fail(t, "unxpected error", err)
		} else if tc.wantErr && err == nil {
			require.Fail(t, "expected error missing", tc.err)
		}
	}
}
