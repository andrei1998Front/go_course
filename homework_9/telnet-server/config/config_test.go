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
		name    string
		args    []string
		want    int
		wantErr bool
		err     error
	}{
		{
			name: "Success",
			args: []string{"1"},
			want: 1,
		},
		{
			name: "Success. Empty args",
			args: []string{},
			want: 8080,
		},
		{
			name:    "Failure. Many args",
			args:    []string{"1", "2"},
			want:    0,
			wantErr: true,
			err:     ErrorManyArgs,
		},
		{
			name:    "Failure. Invalid port value",
			args:    []string{"s"},
			want:    0,
			wantErr: true,
			err:     ErrorInvalidPortValue,
		},
		{
			name:    "Failure. Out range",
			args:    []string{"-1"},
			want:    0,
			wantErr: true,
			err:     ErrorOutPortRange,
		},
		{
			name:    "Failure. Out range",
			args:    []string{"65536"},
			want:    0,
			wantErr: true,
			err:     ErrorOutPortRange,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := setPort(tt.args)

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

/*
func TestConfig_Init(t *testing.T) {
	mockCfg := Config{}

	tests := []struct {
		name            string
		cfg             *Config
		mockStringOut   string
		mockArgs        []string
		wantErr         bool
		err             error
		excludeArgsMock bool
		wantTimeout     time.Duration
		wantPort        int
	}{
		{
			name:          "Success",
			cfg:           &mockCfg,
			mockStringOut: "10s",
			mockArgs:      []string{"0"},
			wantTimeout:   time.Duration(10) * time.Second,
			wantPort:      0,
		},
		{
			name:          "Success. Empty ports",
			cfg:           &mockCfg,
			mockStringOut: "10s",
			mockArgs:      []string{},
			wantTimeout:   time.Duration(10) * time.Second,
			wantPort:      8080,
		},
		{
			name:            "Failure. Invalid timeout",
			cfg:             &mockCfg,
			mockStringOut:   "10",
			mockArgs:        []string{},
			wantTimeout:     time.Duration(0),
			excludeArgsMock: true,
			wantPort:        0,
			wantErr:         true,
			err:             ErrorInvalidTimeout,
		},
		{
			name:          "Failure. Invalid port",
			cfg:           &mockCfg,
			mockStringOut: "10s",
			mockArgs:      []string{"s"},
			wantTimeout:   time.Duration(0),
			wantPort:      0,
			wantErr:       true,
			err:           ErrorInvalidPortValue,
		},
		{
			name:          "Failure. Out range",
			cfg:           &mockCfg,
			mockStringOut: "10s",
			mockArgs:      []string{"-1"},
			wantTimeout:   time.Duration(0),
			wantPort:      0,
			wantErr:       true,
			err:           ErrorOutPortRange,
		},
		{
			name:          "Failure. Out range",
			cfg:           &mockCfg,
			mockStringOut: "10s",
			mockArgs:      []string{"65536"},
			wantTimeout:   time.Duration(0),
			wantPort:      0,
			wantErr:       true,
			err:           ErrorOutPortRange,
		},
		{
			name:          "Failure. Many args",
			cfg:           &mockCfg,
			mockStringOut: "10s",
			mockArgs:      []string{"65536", "222"},
			wantTimeout:   time.Duration(0),
			wantPort:      0,
			wantErr:       true,
			err:           ErrorManyArgs,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			stringPointer := &tt.mockStringOut

			mockFG := mocks.NewFlagGetter(t)

			mockFG.On("String",
				mock.AnythingOfType("string"),
				mock.AnythingOfType("string"),
				mock.AnythingOfType("string")).
				Return(stringPointer)
			if !tt.excludeArgsMock {
				mockFG.On("Args").
					Return(tt.mockArgs)
			}

			err := tt.cfg.Init(mockFG)

			if (err != nil) != tt.wantErr {
				t.Errorf("Config.Init() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if err != nil && tt.wantErr {
				require.ErrorIs(t, err, tt.err)
			}

			require.Equal(t, tt.wantPort, tt.cfg.Port)
			require.Equal(t, tt.wantTimeout, tt.cfg.Timeout)

			mockCfg.Port = 0
			mockCfg.Timeout = time.Duration(0)
		})
	}
}*/
