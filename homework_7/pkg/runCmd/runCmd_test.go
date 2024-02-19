package runcmd

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestRunCmd(t *testing.T) {
	var TestTable = []struct {
		cmd         []string
		env         map[string]string
		expected    int
		expectedErr bool
		err         string
	}{
		{
			cmd: []string{"cmd"},
			env: map[string]string{
				"A_VAR": "123",
				"B_ENV": "ccgvb",
			},
			expected:    0,
			expectedErr: false,
			err:         "",
		},
		{
			cmd: []string{},
			env: map[string]string{
				"A_VAR": "123",
				"B_ENV": "ccgvb",
			},
			expected:    0,
			expectedErr: true,
			err:         "ошибка. Команда не передана",
		},
		{
			cmd: []string{"fff"},
			env: map[string]string{
				"A_VAR": "123",
				"B_ENV": "ccgvb",
			},
			expected:    0,
			expectedErr: true,
			err:         "Неизвестная команда:",
		},
		{
			cmd: []string{"cmd", "/C", "date"},
			env: map[string]string{
				"A_VAR": "123",
				"B_ENV": "ccgvb",
			},
			expected:    1,
			expectedErr: true,
			err:         "Ошибка выполнения команды:",
		},
	}

	for _, testCase := range TestTable {
		result, err := RunCmd(testCase.cmd, testCase.env)

		if testCase.expectedErr && err != nil {
			require.ErrorContains(t, err, testCase.err)
			require.Equal(t, testCase.expected, result)
			continue
		} else if !testCase.expectedErr && err != nil {
			t.Fatal("Получена неожиданная ошибка: ", err)
		}

		require.Equal(t, testCase.expected, result)
	}
}
