package main

import (
	"errors"
	"fmt"
	"log"
	"math/rand"
)

func generateArrOfFunc(l int) []func() error {
	var arr []func() error

	for i := 0; i < l; i++ {
		arr = append(arr, func() error {
			if rand.Intn(10) == 1 {
				return errors.New("some error")
			}

			return nil
		})
	}

	return arr
}

func runMultipleParallelJobs(arrOfFunc []func() error, countParallelJobs int, maxErrCount int) (int, error) {
	var currentCountErrors int

	if countParallelJobs > len(arrOfFunc) {
		return 0, errors.New("Число параллельно выполняемых задач больше, чем их есть")
	}

	errChan := make(chan error, maxErrCount)

	for i := 0; i < countParallelJobs; i++ {
		if currentCountErrors < maxErrCount {
			go func(f func() error) {
				err := f()
				errChan <- err
			}(arrOfFunc[i])

		} else {
			break
		}

		x := <-errChan
		if x != nil {
			fmt.Println(x, ": ", i)
			currentCountErrors++
		}
	}

	return currentCountErrors, nil
}

func main() {
	arrOfFunc := generateArrOfFunc(20)

	v, err := runMultipleParallelJobs(arrOfFunc, 20, 1)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Количество ошибок: ", v)
}
