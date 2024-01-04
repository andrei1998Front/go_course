package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGetCommonWords(t *testing.T) {

	var TestTable = []struct {
		txt    string
		result CommonWordList
		err    string
	}{
		{
			txt: "Привет привет привет dffd dffd шесть шесть восемь девять dsg кавпвпр укпаваппва купк кеке  ggg",
			result: CommonWordList{
				{
					word:  "привет",
					count: 3,
				},
				{
					word:  "dffd",
					count: 2,
				},
				{
					word:  "шесть",
					count: 2,
				},
				{
					word:  "восемь",
					count: 1,
				},
				{
					word:  "девять",
					count: 1,
				},
				{
					word:  "dsg",
					count: 1,
				},
				{
					word:  "кавпвпр",
					count: 1,
				},
				{
					word:  "укпаваппва",
					count: 1,
				},
				{
					word:  "купк",
					count: 1,
				},
				{
					word:  "кеке",
					count: 1,
				},
			},
			err: "",
		},
		{
			txt:    "r k w",
			result: CommonWordList{},
			err:    "Менее 10 слов!",
		},
		{
			txt:    "Привет привет привет dffd dffd шесть шесть восемь девять dsg ",
			result: CommonWordList{},
			err:    "Менее 10 уникальных слов!",
		},
	}

	for _, testCase := range TestTable {
		result, err := getCommonWords(testCase.txt)

		if err != nil {
			require.EqualError(t, err, testCase.err)
		}

		require.Equal(t, testCase.result, result, "")
	}
}
