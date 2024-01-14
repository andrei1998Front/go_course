package main

import (
	"errors"
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/andrei1998Front/go_course/homework_3/scanner"
)

func CreateArrOfDuplicateValues(sym string, l int) ([]string, error) {
	var arr []string

	if l < 0 {
		return nil, errors.New("Количество дублирующихся символов не может быть отрицательным")
	}

	for i := 0; i < l; i++ {
		arr = append(arr, sym)
	}

	return arr, nil
}

func checkValue(value string) (int, bool) {
	if value == "0" {
		return 0, true
	}

	val, err := strconv.Atoi(value)

	if err == nil {
		return val, true
	} else {
		return -1, false
	}
}

func appendSymbols(currentIdx int, arr []string) ([]string, error) {
	var arrOfUnpackingString []string
	var arrOfDublicate []string
	var err error
	var nextIsInt bool

	nextIdx := currentIdx + 1
	prevIdx := currentIdx - 1
	lenArr := len(arr)
	lastIdx := lenArr - 1

	currentNum, currentIsInt := checkValue(arr[currentIdx])

	if nextIdx > lastIdx {
		if currentIsInt == false {
			arrOfUnpackingString = append(arrOfUnpackingString, arr[currentIdx])
		} else if nextIdx > lastIdx && currentIsInt == true {
			arrOfDublicate, err = CreateArrOfDuplicateValues(arr[prevIdx], currentNum)

			if err != nil {
				return nil, err
			}

			arrOfUnpackingString = append(arrOfUnpackingString, arrOfDublicate...)
		}
	} else {
		_, nextIsInt = checkValue(arr[nextIdx])

		if currentIsInt == false && nextIsInt == false {
			arrOfUnpackingString = append(arrOfUnpackingString, arr[currentIdx])
		} else if currentIsInt == true {

			if currentIdx == 0 {
				return nil, errors.New("Некорректная строка! Строка начинается с числового значения")
			} else {
				_, prevIsInt := checkValue(arr[prevIdx])

				if nextIsInt == true || prevIsInt == true {
					return nil, errors.New("Некорректная строка! Два числовых значения подряд")
				}

				arrOfDublicate, err = CreateArrOfDuplicateValues(arr[prevIdx], currentNum)

				if err != nil {
					return nil, err
				}

				arrOfUnpackingString = append(arrOfUnpackingString, arrOfDublicate...)
			}
		}
	}

	return arrOfUnpackingString, nil
}

func CreateArrOfUnpackingString(symbols []string) ([]string, error) {
	var arrOfUnpackingString []string

	for key := range symbols {
		newSymbols, err := appendSymbols(key, symbols)

		if err != nil {
			return nil, err
		}

		arrOfUnpackingString = append(arrOfUnpackingString, newSymbols...)
	}

	return arrOfUnpackingString, nil
}

func UnpackString(txt string) (string, error) {
	symbols := strings.Split(txt, "")

	newSymbols, err := CreateArrOfUnpackingString(symbols)

	if err != nil {
		return "", err
	}

	return strings.Join(newSymbols, ""), nil
}

func main() {
	fmt.Println("Введите строку:")

	txt, err := scanner.ScanUserInput()

	if err != nil {
		log.Fatal(err)
	}

	str, err := UnpackString(txt)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(str)
}
