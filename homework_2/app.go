package main

import (
	"errors"
	"fmt"
	"log"
	"sort"
	"strings"
)

type CommonWord struct {
	word  string
	count int8
}

type CommonWordList []CommonWord

func (commonWordList CommonWordList) Len() int {
	return len(commonWordList)
}

func (commonWordList CommonWordList) Less(i, j int) bool {
	return commonWordList[i].count > commonWordList[j].count
}

func (commonWordList CommonWordList) Swap(i, j int) {
	commonWordList[i], commonWordList[j] = commonWordList[j], commonWordList[i]
}

func searchRepeat(commonWord CommonWord, commonWordList CommonWordList) (bool, int) {
	for key, value := range commonWordList {
		if value.word == commonWord.word {
			return true, key
		}
	}

	return false, -1
}

func getListOfWordsCount(txt []string) CommonWordList {
	var count_of_words CommonWordList

	for _, word := range txt {
		word := CommonWord{word, 1}

		result, idx := searchRepeat(word, count_of_words)

		if result {
			count_of_words[idx].count += word.count
		} else {
			count_of_words = append(count_of_words, word)
		}
	}

	sort.Sort(count_of_words)

	return count_of_words
}

func printResult(str string, commonWordList CommonWordList) {
	fmt.Printf("Входные данные: \"%s\"\n\n", str)
	fmt.Println("Десять наиболее повторяющихся слов:")

	for _, word := range commonWordList {
		fmt.Printf("Слово  \"%s\". Количество вхождений: %d\n", word.word, word.count)
	}
}

func getCommonWords(txt string) (CommonWordList, error) {
	txt = strings.ToLower(txt)

	strs := strings.Split(txt, " ")

	if len(strs) < 10 {
		return CommonWordList{}, errors.New("Менее 10 слов!")
	}

	count_of_words := getListOfWordsCount(strs)

	if len(count_of_words) < 10 {
		return CommonWordList{}, errors.New("Менее 10 уникальных слов!")
	}

	return count_of_words[0:10], nil
}

func main() {
	str := "Привет привет привет dffd dffd шесть шесть восемь девять dsg кавпвпр укпаваппва купк кеке  ggg"
	v, err := getCommonWords(str)

	if err != nil {
		log.Fatal(err)
	}

	printResult(str, v)
}
