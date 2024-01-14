package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func CheckNextValue(t *testing.T) {
	TestTable := []struct {
		value_input string
		outputInt   int
		outputBool  bool
	}{
		{
			value_input: "0",
			outputInt:   0,
			outputBool:  true,
		},
		{
			value_input: "1",
			outputInt:   0,
			outputBool:  true,
		},
		{
			value_input: "e",
			outputInt:   -1,
			outputBool:  false,
		},
	}

	for _, testCase := range TestTable {
		resultInt, resultBool := checkValue(testCase.value_input)

		require.Equal(t, testCase.outputInt, resultInt, "")
		require.Equal(t, testCase.outputBool, resultBool, "")
	}
}
