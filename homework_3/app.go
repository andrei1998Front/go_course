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

func checkValue(value string, prevIsEscape bool) (int, bool) {
	val, err := strconv.Atoi(value)

	if err == nil && prevIsEscape == false {
		return val, true
	} else {
		return -1, false
	}
}

func CheckEscapeSymbol(sym string) bool {
	if sym == "\\" {
		return true
	}

	return false
}

func SetPrevEscapeFlag(idx int, arr []string) bool {
	if idx < 0 {
		return false
	}

	return CheckEscapeSymbol(arr[idx])
}

func JoinWithDublicate(sym string, l int) ([]string, error) {
	arrOfDublicate, err := CreateArrOfDuplicateValues(sym, l)

	if err != nil {
		return nil, err
	}

	return append([]string{}, arrOfDublicate...), nil
}

func appendSymbols(currentIdx int, arr []string) ([]string, error) {
	var arrOfUnpackingString []string
	var errMain error
	var nextIsInt bool
	var prevIsEscape bool
	var currentIsEcape bool

	nextIdx := currentIdx + 1
	prevIdx := currentIdx - 1
	lenArr := len(arr)
	lastIdx := lenArr - 1

	prevIsEscape = SetPrevEscapeFlag(prevIdx, arr)

	currentIsEcape = CheckEscapeSymbol(arr[currentIdx])

	if prevIsEscape == false && currentIsEcape == true {
		return arrOfUnpackingString, nil
	}

	currentNum, currentIsInt := checkValue(arr[currentIdx], prevIsEscape)

	if nextIdx > lastIdx {
		if currentIsInt == false || (currentIsInt == true && prevIsEscape == true) {
			arrOfUnpackingString, errMain = append(arrOfUnpackingString, arr[currentIdx]), nil
		} else if currentIsInt == true {
			arrOfUnpackingString, errMain = JoinWithDublicate(arr[prevIdx], currentNum)
		}
	} else {
		_, nextIsInt = checkValue(arr[nextIdx], currentIsEcape)

		if (currentIsInt == false || currentIsInt == false && prevIsEscape == true) && nextIsInt == false {
			arrOfUnpackingString, errMain = append(arrOfUnpackingString, arr[currentIdx]), nil
		} else if currentIsInt == true {

			if currentIdx == 0 {
				return nil, errors.New("Некорректная строка! Строка начинается с числового значения")
			} else {

				_, prevIsInt := checkValue(arr[prevIdx], SetPrevEscapeFlag(prevIdx-1, arr))

				if nextIsInt == true || prevIsInt == true {
					return nil, errors.New("Некорректная строка! Два числовых значения подряд")
				}

				arrOfUnpackingString, errMain = JoinWithDublicate(arr[prevIdx], currentNum)
			}
		}
	}

	return arrOfUnpackingString, errMain
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
