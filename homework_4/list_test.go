package main

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/require"
)

func getList() List {
	list := List{}

	list.PushFront(2)
	list.PushBack(2)
	list.PushBack(3)
	list.PushBack(4)

	return list
}

func getListOne() List {
	list := List{}
	list.PushFront(2)

	return list
}

func TestLen(t *testing.T) {
	list := getList()

	require.Equal(t, list.Len(), 4)
}

func TestFirst(t *testing.T) {
	list := getList()

	require.Equal(t, list.First().Value(), 2)
}

func TestLast(t *testing.T) {
	list := getList()

	require.Equal(t, list.Last().Value(), 4)
}

func PushFront(t *testing.T) {
	TestTable := []struct {
		list        List
		input       interface{}
		resultPrev  interface{}
		resultNext  interface{}
		resultFirst interface{}
		resultLast  interface{}
	}{
		{
			list:        getList(),
			input:       1,
			resultNext:  getList().First().Value(),
			resultPrev:  nil,
			resultFirst: 1,
			resultLast:  4,
		},
		{
			list:        List{},
			input:       1,
			resultNext:  nil,
			resultPrev:  nil,
			resultFirst: 1,
			resultLast:  1,
		},
	}

	for _, testCase := range TestTable {
		testCase.list.PushFront(testCase.input)

		require.Equal(t, testCase.list.First().Next().Value(), testCase.resultNext)
		require.Equal(t, testCase.list.First().Prev().Value(), testCase.resultPrev)
		require.Equal(t, testCase.list.First().Value(), testCase.resultFirst)
		require.Equal(t, testCase.list.Last().Value(), testCase.resultLast)
	}
}

func TestPushBack(t *testing.T) {
	TestTable := []struct {
		list        List
		input       interface{}
		resultPrev  interface{}
		resultNext  interface{}
		resultFirst interface{}
		resultLast  interface{}
	}{
		{
			list:        getList(),
			input:       1,
			resultNext:  (*Item)(nil),
			resultPrev:  getList().Last().Value(),
			resultFirst: 2,
			resultLast:  1,
		},
		{
			list:        List{},
			input:       1,
			resultNext:  (*Item)(nil),
			resultPrev:  (*Item)(nil),
			resultFirst: 1,
			resultLast:  1,
		},
	}

	for _, testCase := range TestTable {

		emptyList := List{}
		listForCheck := testCase.list

		testCase.list.PushBack(testCase.input)

		require.Equal(t, testCase.resultNext, testCase.list.Last().Next())
		require.Equal(t, testCase.list.First().Value(), testCase.resultFirst)
		require.Equal(t, testCase.list.Last().Value(), testCase.resultLast)

		if emptyList == listForCheck {
			require.Equal(t, testCase.resultPrev, testCase.list.Last().Prev())
		} else {
			require.Equal(t, testCase.resultPrev, testCase.list.Last().Prev().Value())
		}
	}
}

func TestCheckItem(t *testing.T) {
	TestList := getList()
	TestTable := []struct {
		list   List
		input  *Item
		result bool
	}{
		{
			list:   TestList,
			input:  TestList.First().Next().Next(),
			result: true,
		},
		{
			list:   getList(),
			input:  getList().First().Next().Next(),
			result: false,
		},
	}

	for _, testCase := range TestTable {
		require.Equal(t, testCase.result, testCase.list.CheckItem(testCase.input))
	}
}

func TestRemove(t *testing.T) {
	TestList := getList()

	TestTable := []struct {
		list List
		item *Item
		err  string
	}{
		{
			list: getList(),
			item: getList().First().Next(),
			err:  "Удаляемый элемент не находится в списке",
		},
		{
			list: TestList,
			item: TestList.First().Next(),
			err:  "",
		},
	}

	for _, testCase := range TestTable {
		err := testCase.list.Remove(testCase.item)
		if err == nil {
			err = errors.New("")
		}
		require.EqualError(t, err, testCase.err)
	}
}

func TestRemoveSelectedItem(t *testing.T) {
	list := getList()
	listOne := getListOne()

	TestTable := []struct {
		list      List
		input     *Item
		listFirst *Item
		listLast  *Item
	}{
		{
			list:      list,
			input:     list.First(),
			listFirst: list.First().Next(),
			listLast:  list.Last(),
		},
		{
			list:      listOne,
			input:     listOne.First(),
			listFirst: (*Item)(nil),
			listLast:  (*Item)(nil),
		},
	}

	for _, testCase := range TestTable {
		testCase.list.removeSelectedItem(testCase.input)
		require.Equal(t, testCase.listFirst, testCase.list.First())
		require.Equal(t, testCase.listLast, testCase.list.Last())
		require.Equal(t, testCase.list.CheckItem(testCase.input), false)
	}
}
