package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func CheckNextValue(t *testing.T) {
	TestTable := []struct {
		value_input  string
		prevIsEscape bool
		outputInt    int
		outputBool   bool
	}{
		{
			value_input:  "0",
			prevIsEscape: true,
			outputInt:    -1,
			outputBool:   true,
		},
		{
			value_input:  "1",
			prevIsEscape: false,
			outputInt:    1,
			outputBool:   true,
		},
		{
			value_input: "e",
			outputInt:   -1,
			outputBool:  false,
		},
	}

	for _, testCase := range TestTable {
		resultInt, resultBool := checkValue(testCase.value_input, testCase.prevIsEscape)

		require.Equal(t, testCase.outputInt, resultInt, "")
		require.Equal(t, testCase.outputBool, resultBool, "")
	}
}
