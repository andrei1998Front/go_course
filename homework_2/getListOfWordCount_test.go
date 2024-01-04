package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGetListOfWordsCount(t *testing.T) {

	var TestTable = []struct {
		txt    []string
		result CommonWordList
	}{
		{
			txt: []string{
				"ddd",
				"dfff",
				"ffff",
			},
			result: CommonWordList{
				CommonWord{
					word:  "ddd",
					count: 1,
				},
				CommonWord{
					word:  "dfff",
					count: 1,
				},
				CommonWord{
					word:  "ffff",
					count: 1,
				},
			},
		},

		{
			txt: []string{
				"ddd",
				"ddd",
				"ffff",
			},
			result: CommonWordList{
				CommonWord{
					word:  "ddd",
					count: 2,
				},
				CommonWord{
					word:  "ffff",
					count: 1,
				},
			},
		},

		{
			txt: []string{
				"ddd",
				"ddd",
				"ffff",
				"ffff",
			},
			result: CommonWordList{
				CommonWord{
					word:  "ddd",
					count: 2,
				},
				CommonWord{
					word:  "ffff",
					count: 2,
				},
			},
		},

		{
			txt: []string{
				"ddd",
				"ddd",
				"ddd",
				"ddd",
			},
			result: CommonWordList{
				CommonWord{
					word:  "ddd",
					count: 4,
				},
			},
		},
	}

	for _, testCase := range TestTable {
		result := getListOfWordsCount(testCase.txt)

		require.Equal(t, result, testCase.result, "")
	}
}
