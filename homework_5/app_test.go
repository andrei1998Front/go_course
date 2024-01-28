package main

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGenereArrOfFunc(t *testing.T) {
	TestTable := []struct {
		input      int
		output_len int
		err        string
	}{
		{
			input:      3,
			output_len: 3,
			err:        "",
		},
		{
			input:      -1,
			output_len: 0,
			err:        "Число элементов массива не может быть отрицательным.",
		},
	}

	for _, testCase := range TestTable {
		result, err := generateArrOfFunc(testCase.input)

		if err != nil {
			require.EqualError(t, err, testCase.err)
		}

		require.Equal(t, testCase.output_len, len(result))
	}
}

func getExample() []func() error {
	var firstExample []func() error

	for i := 0; i < 10; i++ {
		firstExample = append(firstExample, func() error {
			return errors.New("some error")
		})
	}

	for i := 0; i < 10; i++ {
		firstExample = append(firstExample, func() error {
			return nil
		})
	}

	return firstExample
}

func TestRunMultipleParallelJobs(t *testing.T) {

	TestTable := []struct {
		arrOfFunc         []func() error
		countParallelJobs int
		maxErrCount       int
		output            int
		err               string
	}{
		{
			arrOfFunc:         getExample(),
			countParallelJobs: 100,
			maxErrCount:       2,
			output:            0,
			err:               "Число параллельно выполняемых задач больше, чем их есть",
		},
		{
			arrOfFunc:         getExample(),
			countParallelJobs: 20,
			maxErrCount:       -2,
			output:            0,
			err:               "Аргументы функции не могут быть отрицательными",
		},
		{
			arrOfFunc:         getExample(),
			countParallelJobs: -20,
			maxErrCount:       2,
			output:            0,
			err:               "Аргументы функции не могут быть отрицательными",
		},
		{
			arrOfFunc:         getExample(),
			countParallelJobs: 20,
			maxErrCount:       20,
			output:            10,
			err:               "",
		},
	}

	for _, testCase := range TestTable {
		result, err := runMultipleParallelJobs(testCase.arrOfFunc, testCase.countParallelJobs, testCase.maxErrCount)

		if err != nil {
			require.EqualError(t, err, testCase.err)
		}

		require.Equal(t, testCase.output, result)
	}
}
