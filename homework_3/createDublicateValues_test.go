package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCreateDublicateValues(t *testing.T) {
	TestTable := []struct {
		sym_input string
		l_input   int
		output    []string
		err       string
	}{
		{
			sym_input: "a",
			l_input:   1,
			output: []string{
				"a",
			},
			err: "",
		},
		{
			sym_input: "a",
			l_input:   2,
			output: []string{
				"a",
				"a",
			},
			err: "",
		},
		{
			sym_input: "a",
			l_input:   -1,
			output:    nil,
			err:       "Количество дублирующихся символов не может быть отрицательным",
		},
	}

	for _, testCase := range TestTable {
		arrOfDublicate, err := CreateArrOfDuplicateValues(testCase.sym_input, testCase.l_input)

		if err != nil {
			require.EqualError(t, err, testCase.err)
		}

		require.Equal(t, testCase.output, arrOfDublicate, "")
	}
}
