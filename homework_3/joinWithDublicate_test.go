package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestJoinWithDublicate(t *testing.T) {
	TestTable := []struct {
		sym    string
		l      int
		output []string
		err    string
	}{
		{
			sym: "a",
			l:   3,
			output: []string{
				"a",
				"a",
				"a",
			},
			err: "",
		},
		{
			sym:    "a",
			l:      -1,
			output: []string(nil),
			err:    "Количество дублирующихся символов не может быть отрицательным",
		},
	}

	for _, testCase := range TestTable {
		result, err := JoinWithDublicate(testCase.sym, testCase.l)

		if err != nil {
			require.EqualError(t, err, testCase.err)
		}

		require.Equal(t, testCase.output, result)
	}
}
