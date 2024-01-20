package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCheckEscapeSymbol(t *testing.T) {
	TestTable := []struct {
		sym    string
		output bool
	}{
		{
			sym:    "f",
			output: false,
		},
		{
			sym:    "\\",
			output: true,
		},
	}

	for _, testCase := range TestTable {
		result := CheckEscapeSymbol(testCase.sym)

		require.Equal(t, testCase.output, result)
	}
}
