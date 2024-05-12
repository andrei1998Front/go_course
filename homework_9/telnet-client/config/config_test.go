package config

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func Test_setTimeOut(t *testing.T) {
	tests := []struct {
		name     string
		toString string
		want     time.Duration
		wantErr  bool
		err      error
	}{
		{
			name:     "Success. Second",
			toString: "10s",
			want:     time.Duration(10) * time.Second,
		},
		{
			name:     "Success. Millisecond",
			toString: "10ms",
			want:     time.Duration(10) * time.Millisecond,
		},
		{
			name:     "Success. Nanosecond",
			toString: "10ns",
			want:     time.Duration(10) * time.Nanosecond,
		},
		{
			name:     "Success. Microseconds",
			toString: "10us",
			want:     time.Duration(10) * time.Microsecond,
		},
		{
			name:     "Success. Hours",
			toString: "10h",
			want:     time.Duration(60) * time.Minute * 10,
		},
		{
			name:     "Success. Minutes",
			toString: "10m",
			want:     time.Duration(10) * time.Minute,
		},
		{
			name:     "Success. Combo",
			toString: "10m10s",
			want:     time.Duration(10)*time.Minute + time.Duration(10)*time.Second,
		},
		{
			name:     "Failure. Numbers",
			toString: "10",
			want:     time.Duration(0),
			wantErr:  true,
			err:      ErrorInvalidTimeout,
		},
		{
			name:     "Failure. String",
			toString: "sss",
			want:     time.Duration(0),
			wantErr:  true,
			err:      ErrorInvalidTimeout,
		},
		{
			name:     "Failure. String, num",
			toString: "s10",
			want:     time.Duration(0),
			wantErr:  true,
			err:      ErrorInvalidTimeout,
		},
		{
			name:     "Failure. String, num 2",
			toString: "10 m",
			want:     time.Duration(0),
			wantErr:  true,
			err:      ErrorInvalidTimeout,
		},
		{
			name:     "Failure. invalid type",
			toString: "10mm",
			want:     time.Duration(0),
			wantErr:  true,
			err:      ErrorInvalidTimeout,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := setTimeOut(tt.toString)

			if (err != nil) != tt.wantErr {
				t.Errorf("setTimeOut() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if err != nil && tt.wantErr {
				require.ErrorIs(t, err, tt.err)
			}

			require.Equal(t, tt.want, got)
		})
	}
}

func Test_setPort(t *testing.T) {
	tests := []struct {
		name      string
		portValue string
		want      int
		wantErr   bool
		err       error
	}{
		{
			name:      "Success",
			portValue: "1",
			want:      1,
		},
		{
			name:      "Failure. Invalid port value",
			portValue: "s",
			want:      0,
			wantErr:   true,
			err:       ErrorInvalidPortValue,
		},
		{
			name:      "Failure. Out range",
			portValue: "-1",
			want:      0,
			wantErr:   true,
			err:       ErrorOutPortRange,
		},
		{
			name:      "Failure. Out range",
			portValue: "65536",
			want:      0,
			wantErr:   true,
			err:       ErrorOutPortRange,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := setPort(tt.portValue)

			if (err != nil) != tt.wantErr {
				t.Errorf("setPort() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if err != nil && tt.wantErr {
				require.ErrorIs(t, err, tt.err)
			}

			require.Equal(t, tt.want, got)
		})
	}
}
