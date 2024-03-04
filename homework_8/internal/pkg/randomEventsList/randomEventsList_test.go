package randomeventslist

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCheckMaxMin(t *testing.T) {
	TestTable := []struct {
		minYY   int
		maxYY   int
		wantErr bool
		err     string
	}{
		{
			minYY:   1,
			maxYY:   200,
			wantErr: false,
			err:     "",
		},
		{
			minYY:   200,
			maxYY:   1,
			wantErr: true,
			err:     "максимальное значение года меньше минимального",
		},
		{
			minYY:   0,
			maxYY:   1,
			wantErr: true,
			err:     "значение года не может быть нулевым",
		},
		{
			minYY:   -1,
			maxYY:   1,
			wantErr: true,
			err:     "значение года не может быть меньше нуля",
		},
	}

	for _, testCase := range TestTable {
		err := checkMaxMin(testCase.minYY, testCase.maxYY)

		if err != nil && testCase.wantErr {
			require.EqualError(t, err, testCase.err)
		} else if err != nil && !testCase.wantErr {
			t.Log("Минимальное значение года: " + fmt.Sprint(testCase.minYY))
			t.Log("Максимальное значение года: " + fmt.Sprint(testCase.maxYY))
			t.Fatal("Неожиданная ошибка: " + err.Error())
		}
	}
}

func TestGetRandomDate(t *testing.T) {
	TestTable := []struct {
		minYY   int
		maxYY   int
		wantErr bool
		err     string
	}{
		{
			minYY:   1,
			maxYY:   200,
			wantErr: false,
			err:     "",
		},
		{
			minYY:   200,
			maxYY:   1,
			wantErr: true,
			err:     "максимальное значение года меньше минимального",
		},
		{
			minYY:   0,
			maxYY:   1,
			wantErr: true,
			err:     "значение года не может быть нулевым",
		},
		{
			minYY:   -1,
			maxYY:   1,
			wantErr: true,
			err:     "значение года не может быть меньше нуля",
		},
	}

	for _, testCase := range TestTable {
		_, err := GetRandomDate(testCase.minYY, testCase.maxYY)

		if err != nil && testCase.wantErr {
			require.EqualError(t, err, testCase.err)
		} else if err != nil && !testCase.wantErr {
			t.Log("Минимальное значение года: " + fmt.Sprint(testCase.minYY))
			t.Log("Максимальное значение года: " + fmt.Sprint(testCase.maxYY))
			t.Fatal("Неожиданная ошибка: " + err.Error())
		}
	}
}

func TestGetRandomEventsList(t *testing.T) {
	TestTable := []struct {
		minYY   int
		maxYY   int
		size    int
		wantErr bool
		err     string
		outLen  int
	}{
		{
			minYY:   1,
			maxYY:   200,
			size:    5,
			wantErr: false,
			err:     "",
			outLen:  5,
		},
		{
			minYY:   1,
			maxYY:   200,
			size:    0,
			wantErr: true,
			err:     "размер слайса событий не может быть нулевым",
			outLen:  0,
		},
		{
			minYY:   1,
			maxYY:   200,
			size:    -1,
			wantErr: true,
			err:     "размер слайса событий не может быть отрицательным",
			outLen:  0,
		},
		{
			minYY:   200,
			maxYY:   1,
			size:    5,
			wantErr: true,
			err:     "максимальное значение года меньше минимального",
			outLen:  0,
		},
		{
			minYY:   0,
			maxYY:   1,
			wantErr: true,
			err:     "значение года не может быть нулевым",
			outLen:  0,
		},
		{
			minYY:   -1,
			maxYY:   1,
			size:    5,
			wantErr: true,
			err:     "значение года не может быть меньше нуля",
			outLen:  0,
		},
	}

	for _, testCase := range TestTable {
		t.Log("Минимальное значение года: " + fmt.Sprint(testCase.minYY))
		t.Log("Максимальное значение года: " + fmt.Sprint(testCase.maxYY))
		t.Log("Размер слайса событий: " + fmt.Sprint(testCase.size))

		result, err := GetRandomEventsList(testCase.size, testCase.minYY, testCase.maxYY)

		if err != nil && testCase.wantErr {
			require.EqualError(t, err, testCase.err)
		} else if err != nil && !testCase.wantErr {
			t.Fatal("Неожиданная ошибка: " + err.Error())
		}

		require.Len(t, result, testCase.outLen)
	}
}
