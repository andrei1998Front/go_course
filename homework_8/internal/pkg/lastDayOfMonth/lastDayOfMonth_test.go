package lastdayofmonth

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestGetLastDayOfMonth(t *testing.T) {
	TestTable := []struct {
		yy        int
		mm        int
		wantError bool
		expected  time.Time
		err       string
	}{
		{
			yy:        2024,
			mm:        12,
			wantError: false,
			expected:  time.Date(2024, time.Month(12), 31, 0, 0, 0, 0, time.UTC),
			err:       "",
		},
		{
			yy:        2024,
			mm:        -1,
			wantError: true,
			expected:  time.Time{},
			err:       "неверное значение месяца. Значение месяца должно быть в диапозоне от 1 до 12",
		},
		{
			yy:        2024,
			mm:        13,
			wantError: true,
			expected:  time.Time{},
			err:       "неверное значение месяца. Значение месяца должно быть в диапозоне от 1 до 12",
		},
		{
			yy:        -1,
			mm:        1,
			wantError: true,
			expected:  time.Time{},
			err:       "неверно значение года. Должно быть не меньшу нуля",
		},
	}

	for _, testCase := range TestTable {
		result, err := GetLastDayOfMonth(testCase.yy, testCase.mm)

		if err != nil && !testCase.wantError {
			t.Fatal("Неожиданная ошибка")
		} else if err != nil {
			require.EqualError(t, err, testCase.err)
		}

		require.Equal(t, testCase.expected, result)
	}
}
