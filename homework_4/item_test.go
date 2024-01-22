package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func getTestCases() []Item {
	first := Item{value: "первый"}
	average := Item{value: "второй"}
	last := Item{value: "третий"}

	first.next = &average
	average.prev = &first
	average.next = &last
	last.prev = &average

	return []Item{first, average, last}
}

func TestValue(t *testing.T) {
	testTable := getTestCases()

	for _, testCase := range testTable {
		require.Equal(t, testCase.Value(), testCase.value)
	}
}

func TestNext(t *testing.T) {
	testTable := getTestCases()

	require.Equal(t, testTable[0].Next().Value(), testTable[1].Value())
}

func TestPrev(t *testing.T) {
	testTable := getTestCases()

	require.Equal(t, testTable[1].Prev().Value(), testTable[0].Value())
}
