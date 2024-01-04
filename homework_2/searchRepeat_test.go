package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSearchRepeat(t *testing.T) {
	var commonWordList = CommonWordList([]CommonWord{
		CommonWord{"word1", 1},
		CommonWord{"word2", 1},
		CommonWord{"word3", 1},
	})

	TestTable := []struct {
		commonWord     CommonWord
		commonWordList CommonWordList
		resultBool     bool
		resultInt      int
	}{
		{
			commonWord:     CommonWord{"word1", 1},
			commonWordList: commonWordList,
			resultBool:     true,
			resultInt:      0,
		},
		{
			commonWord:     CommonWord{"fff", 1},
			commonWordList: commonWordList,
			resultBool:     false,
			resultInt:      -1,
		},
	}

	for _, testCase := range TestTable {
		resultBool, resultInt := searchRepeat(testCase.commonWord, testCase.commonWordList)

		t.Logf("Calling searchRepeat(commonWord, commonWordList)\n")
		t.Log("commonWord: ", testCase.commonWord)
		t.Log("commonWord: ", testCase.commonWordList)
		require.Equal(t, testCase.resultBool, resultBool, "")
		require.Equal(t, testCase.resultInt, resultInt, "")
	}
}
