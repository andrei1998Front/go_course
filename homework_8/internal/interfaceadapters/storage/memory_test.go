package memory

import (
	"testing"
	"time"

	"github.com/andrei1998Front/go_course/homework_8/internal/domain/event"
	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
)

type fields struct {
	events map[string]event.Event
}

func TestRepo_GetByID(t *testing.T) {
	mockUUID := uuid.MustParse("3e204a57-4449-4c74-8227-77934cf25322")
	type args struct {
		id uuid.UUID
	}

	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *event.Event
		wantErr bool
		err     string
	}{
		{
			name: "Success: event exists, should retutn event",
			fields: fields{
				events: func() map[string]event.Event {
					mp := make(map[string]event.Event)
					mp[mockUUID.String()] = event.Event{ID: mockUUID}
					return mp
				}(),
			},
			args: args{
				id: mockUUID,
			},
			want:    &event.Event{ID: mockUUID},
			wantErr: false,
			err:     "",
		},
		{
			name: "Success: event not exists, should retutn nil",
			fields: fields{
				events: make(map[string]event.Event),
			},
			args: args{
				id: mockUUID,
			},
			want:    &event.Event{},
			wantErr: true,
			err:     event.ErrNonExistentEvent.Error(),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo := Repo{
				events: tt.fields.events,
			}

			got, err := repo.GetByID(tt.args.id)

			if (err != nil) != tt.wantErr {
				t.Errorf("Repo.GetByID() error = %v, wantErr %v", err, tt.wantErr)
				return
			} else if err != nil {
				require.ErrorContains(t, err, tt.err)
			}

			require.Equal(t, tt.want, got)
		})
	}
}

func TestRepo_GetAll(t *testing.T) {
	mockUUID1 := uuid.MustParse("3e204a57-4449-4c74-8227-77934cf25322")
	mockUUID2 := uuid.MustParse("4e204a57-4449-4c74-8227-77934cf25322")

	tests := []struct {
		name    string
		fields  fields
		want    []event.Event
		wantErr bool
		err     string
	}{
		{
			name: "Success: should return 2 event",
			fields: fields{
				events: func() map[string]event.Event {
					mp := make(map[string]event.Event)
					mp[mockUUID1.String()] = event.Event{ID: mockUUID1}
					mp[mockUUID2.String()] = event.Event{ID: mockUUID2}
					return mp
				}(),
			},
			want: []event.Event{
				event.Event{ID: mockUUID1},
				event.Event{ID: mockUUID2},
			},
			wantErr: false,
			err:     "",
		},
		{
			name:    "should return 0 events",
			fields:  fields{events: make(map[string]event.Event)},
			want:    ([]event.Event)(nil),
			wantErr: false,
			err:     "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo := Repo{
				events: tt.fields.events,
			}

			got, err := repo.GetAll()

			if (err != nil) != tt.wantErr {
				t.Errorf("Repo.GetByID() error = %v, wantErr %v", err, tt.wantErr)
				return
			} else if err != nil {
				require.ErrorContains(t, err, tt.err)
			}

			require.Equal(t, tt.want, got)
		})
	}
}

func TestRepo_GetByDate(t *testing.T) {
	mockUUID := uuid.MustParse("3e204a57-4449-4c74-8227-77934cf25322")
	dt := time.Date(2024, time.Month(1), 1, 0, 0, 0, 0, time.UTC)

	type args struct {
		dt time.Time
	}

	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *event.Event
		wantErr bool
		err     string
	}{
		{
			name: "Success: have dt, should return event",
			fields: fields{
				events: func() map[string]event.Event {
					m := make(map[string]event.Event)
					m[mockUUID.String()] = event.Event{ID: mockUUID, Date: dt}
					return m
				}(),
			},

			args: args{
				dt: dt,
			},
			want:    &event.Event{ID: mockUUID, Date: dt},
			wantErr: false,
			err:     "",
		},
		{
			name: "not exists dt, should return nil",
			fields: fields{
				events: make(map[string]event.Event),
			},

			args: args{
				dt: dt,
			},
			want:    &event.Event{},
			wantErr: true,
			err:     event.ErrNonExistentDate.Error(),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo := Repo{
				events: tt.fields.events,
			}

			got, err := repo.GetByDate(tt.args.dt)

			if (err != nil) != tt.wantErr {
				t.Errorf("Repo.GetByID() error = %v, wantErr %v", err, tt.wantErr)
				return
			} else if err != nil {
				require.ErrorContains(t, err, tt.err)
			}

			require.Equal(t, tt.want, got)
		})
	}
}

func TestRepo_Add(t *testing.T) {
	mockUUID := uuid.MustParse("3e204a57-4449-4c74-8227-77934cf25322")
	mockUUID1 := uuid.MustParse("4e204a57-4449-4c74-8227-77934cf25322")

	dt := time.Date(2024, time.Month(1), 1, 0, 0, 0, 0, time.UTC)

	type args struct {
		e event.Event
	}

	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
		err     string
	}{
		{
			name: "Success: should add event",
			fields: fields{
				events: make(map[string]event.Event),
			},
			args: args{
				e: event.Event{ID: mockUUID},
			},
			wantErr: false,
			err:     "",
		},
		{
			name: "date busy. should return error",
			fields: fields{
				events: func() map[string]event.Event {
					m := make(map[string]event.Event)
					m[mockUUID.String()] = event.Event{ID: mockUUID, Date: dt}
					return m
				}(),
			},
			args: args{
				e: event.Event{ID: mockUUID1, Date: dt},
			},
			wantErr: true,
			err:     event.ErrDateBusy.Error(),
		},
		{
			name: "ID already exists. should return error",
			fields: fields{
				events: func() map[string]event.Event {
					m := make(map[string]event.Event)
					m[mockUUID.String()] = event.Event{ID: mockUUID}
					return m
				}(),
			},
			args: args{
				e: event.Event{ID: mockUUID},
			},
			wantErr: true,
			err:     event.ErrExistentID.Error(),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo := Repo{
				events: tt.fields.events,
			}

			err := repo.Add(tt.args.e)

			if (err != nil) != tt.wantErr {
				t.Errorf("Repo.GetByID() error = %v, wantErr %v", err, tt.wantErr)
				return
			} else if err != nil {
				require.ErrorContains(t, err, tt.err)
				return
			}

			e, _ := repo.GetByID(mockUUID)
			require.Equal(t, tt.args.e, *e)
		})
	}
}

func TestRepo_Delete(t *testing.T) {
	mockUUID := uuid.MustParse("3e204a57-4449-4c74-8227-77934cf25322")

	type args struct {
		ID uuid.UUID
	}

	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
		err     string
	}{
		{
			name: "Success",
			fields: fields{
				events: func() map[string]event.Event {
					m := make(map[string]event.Event)
					m[mockUUID.String()] = event.Event{ID: mockUUID}
					return m
				}(),
			},
			args: args{
				ID: mockUUID,
			},
			wantErr: false,
			err:     "",
		},
		{
			name: "should be return error",
			fields: fields{
				events: make(map[string]event.Event),
			},
			args: args{
				ID: mockUUID,
			},
			wantErr: true,
			err:     event.ErrNonExistentEvent.Error(),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo := Repo{
				events: tt.fields.events,
			}

			err := repo.Delete(tt.args.ID)

			if (err != nil) != tt.wantErr {
				t.Errorf("Repo.GetByID() error = %v, wantErr %v", err, tt.wantErr)
				return
			} else if err != nil {
				require.ErrorContains(t, err, tt.err)
				return
			}

			_, err = repo.GetByID(mockUUID)
			require.Equal(t, true, err != nil)
		})
	}
}

func TestRepo_Update(t *testing.T) {
	mockUUID := uuid.MustParse("3e204a57-4449-4c74-8227-77934cf25322")
	mockUUID1 := uuid.MustParse("4e204a57-4449-4c74-8227-77934cf25322")

	dt1 := time.Date(2024, time.Month(1), 1, 0, 0, 0, 0, time.UTC)
	dt2 := time.Date(2024, time.Month(2), 1, 0, 0, 0, 0, time.UTC)

	type args struct {
		e event.Event
	}

	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
		err     string
	}{
		{
			name: "Success",
			fields: fields{
				events: func() map[string]event.Event {
					m := make(map[string]event.Event)
					m[mockUUID.String()] = event.Event{ID: mockUUID, Date: dt1}
					return m
				}(),
			},
			args: args{
				e: event.Event{ID: mockUUID, Date: dt2},
			},
			wantErr: false,
			err:     "",
		},
		{
			name: "date is busy, should be return error",
			fields: fields{
				events: func() map[string]event.Event {
					m := make(map[string]event.Event)
					m[mockUUID.String()] = event.Event{ID: mockUUID, Date: dt1}
					return m
				}(),
			},
			args: args{
				e: event.Event{ID: mockUUID, Date: dt1},
			},
			wantErr: true,
			err:     event.ErrDateBusy.Error(),
		},
		{
			name: "event not exists, should be return error",
			fields: fields{
				events: func() map[string]event.Event {
					m := make(map[string]event.Event)
					m[mockUUID.String()] = event.Event{ID: mockUUID, Date: dt1}
					return m
				}(),
			},
			args: args{
				e: event.Event{ID: mockUUID1, Date: dt2},
			},
			wantErr: true,
			err:     event.ErrNonExistentEvent.Error(),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo := Repo{
				events: tt.fields.events,
			}

			err := repo.Update(tt.args.e)

			if (err != nil) != tt.wantErr {
				t.Errorf("Repo.GetByID() error = %v, wantErr %v", err, tt.wantErr)
				return
			} else if err != nil {
				require.ErrorContains(t, err, tt.err)
				return
			}

			e, _ := repo.GetByDate(tt.args.e.Date)
			require.Equal(t, tt.args.e.Date, e.Date)
		})
	}
}
