package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCreateArrOfUnpackingString(t *testing.T) {
	TestTable := []struct {
		symbols_input []string
		output        []string
		err           string
	}{
		{
			symbols_input: []string{
				"a",
				"2",
				"b",
				"3",
			},
			output: []string{
				"a",
				"a",
				"b",
				"b",
				"b",
			},
			err: "",
		},
		{
			symbols_input: []string{
				"a",
				"f",
			},
			output: []string{
				"a",
				"f",
			},
			err: "",
		},
		{
			symbols_input: []string{
				"2",
				"2",
				"b",
				"3",
			},
			output: []string(nil),
			err:    "Некорректная строка! Строка начинается с числового значения",
		},
		{
			symbols_input: []string{
				"a",
				"2",
				"3",
				"3",
			},
			output: []string(nil),
			err:    "Некорректная строка! Два числовых значения подряд",
		},
	}

	for _, testCase := range TestTable {
		arrOfDublicate, err := CreateArrOfUnpackingString(testCase.symbols_input)

		if err != nil {
			require.EqualError(t, err, testCase.err)
		}

		require.Equal(t, testCase.output, arrOfDublicate)
	}
}
