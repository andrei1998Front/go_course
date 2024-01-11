package main

import (
	"errors"
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/andrei1998Front/go_course/homework_3/scanner"
)

func CreateArrOfDuplicateValues(sym string, l int) []string {
	var arr []string

	for i := 0; i < l; i++ {
		arr = append(arr, sym)
	}

	return arr
}

func checkNextValue(value string) bool {
	if value == "0" {
		return true
	}

	_, err := strconv.Atoi(value)

	if err == nil {
		return true
	} else {
		return false
	}
}

func CreateArrOfUnpackingString(symbols []string) ([]string, error) {
	var arrOfUnpackingString []string
	var nextIsInt bool

	for key, value := range symbols {

		currentNum, err := strconv.Atoi(value)

		if err == nil && key == 0 {
			return nil, errors.New("Некорректная строка! Строка начинается с числового значения")
		}

		if key+1 > len(symbols)-1 {
			arrOfUnpackingString = append(arrOfUnpackingString, value)
			continue
		}

		nextIsInt = checkNextValue(symbols[key+1])

		if err != nil && !nextIsInt {
			arrOfUnpackingString = append(arrOfUnpackingString, value)
		} else if err == nil && !nextIsInt {
			arrOfUnpackingString = append(arrOfUnpackingString, CreateArrOfDuplicateValues(symbols[key-1], currentNum)...)
		} else if err == nil && nextIsInt {
			return nil, errors.New("Некорректная строка! Два числовых значения подряд")
		}
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
