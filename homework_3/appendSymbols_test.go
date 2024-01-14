package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestAppendSymbols(t *testing.T) {
	TestTable := []struct {
		currentIdx int
		arr        []string
		output     []string
		err        string
	}{
		{
			currentIdx: 0,
			arr: []string{
				"a",
				"2",
				"b",
				"3",
			},
			output: []string(nil),
			err:    "",
		},
		{
			currentIdx: 1,
			arr: []string{
				"a",
				"2",
				"b",
				"3",
			},
			output: []string{
				"a",
				"a",
			},
			err: "",
		},
		{
			currentIdx: 0,
			arr: []string{
				"2",
				"2",
				"b",
				"3",
			},
			output: []string(nil),
			err:    "Некорректная строка! Строка начинается с числового значения",
		},
		{
			currentIdx: 1,
			arr: []string{
				"2",
				"2",
				"b",
				"3",
			},
			output: []string(nil),
			err:    "Некорректная строка! Два числовых значения подряд",
		},
	}

	for _, testCase := range TestTable {
		t.Log("Current index: ", testCase.currentIdx)
		t.Log("Array: ", testCase.arr)
		t.Log("Current symbol: ", testCase.arr[testCase.currentIdx])
		result, err := appendSymbols(testCase.currentIdx, testCase.arr)

		if err != nil {
			require.EqualError(t, err, testCase.err)
		}

		require.Equal(t, testCase.output, result)
	}
}
