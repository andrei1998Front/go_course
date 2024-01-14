package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestUnpackString(t *testing.T) {
	TestTable := []struct {
		txt    string
		output string
		err    string
	}{
		{
			txt:    "aaa",
			output: "aaa",
			err:    "",
		},
		{
			txt:    "a2",
			output: "aa",
			err:    "",
		},
		{
			txt:    "a2b3",
			output: "aabbb",
			err:    "",
		},
		{
			txt:    "2aa",
			output: "",
			err:    "Некорректная строка! Строка начинается с числового значения",
		},
		{
			txt:    "a22",
			output: "",
			err:    "Некорректная строка! Два числовых значения подряд",
		},
	}

	for _, testCase := range TestTable {
		result, err := UnpackString(testCase.txt)

		if err != nil {
			require.EqualError(t, err, testCase.err)
		}

		require.Equal(t, testCase.output, result)
	}
}
