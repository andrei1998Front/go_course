package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSetEcapeFlag(t *testing.T) {
	TestTable := []struct {
		idx    int
		arr    []string
		output bool
	}{
		{
			idx: -1,
			arr: []string{
				"a",
				"v",
				"b",
			},
			output: false,
		},
		{
			idx: 0,
			arr: []string{
				"a",
				"v",
				"v",
			},
			output: false,
		},
		{
			idx: 0,
			arr: []string{
				"\\",
				"a",
				"v",
			},
			output: true,
		},
	}

	for _, testCase := range TestTable {
		result := SetPrevEscapeFlag(testCase.idx, testCase.arr)

		require.Equal(t, testCase.output, result)
	}
}
