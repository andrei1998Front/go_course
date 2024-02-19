package readdir

import (
	"errors"
	"io"
	"io/fs"
	"os"
	"path/filepath"
	"regexp"
)

var re_list = []string{".+_ENV$", ".+_VAR$"}

func checkFileName(fileName string, re_list []string) (bool, error) {
	if len(re_list) == 0 {
		return false, errors.New("массив регулярных выражений не может быть пустым")
	}

	for _, re_item := range re_list {
		re, err := regexp.Compile(re_item)

		if err != nil {
			return false, errors.New("Ошибка компиляции шаблона \"" + re_item + "\": " + err.Error())
		}

		if re.MatchString(fileName) {
			return true, nil
		}
	}

	return false, nil
}

func checkPathExist(pth string) error {
	_, err := os.Stat(pth)

	if err != nil {
		return errors.New("Ошибка чтения пути: " + err.Error())
	}

	return nil
}

func getFileContent(pth string) (string, error) {

	file, err := os.Open(pth)

	if err != nil {
		return "", errors.New("Ошибка открытия файла переменной-окружения, расположенного по пути \"" + pth + "\": " + err.Error())
	}

	defer file.Close()

	fileStat, err := file.Stat()

	if err != nil {
		return "", errors.New("Ошибка создания статистики файла переменной-окружения, расположенного по пути \"" + pth + "\": " + err.Error())
	}

	var buf = make([]byte, fileStat.Size())

	_, err = io.ReadFull(file, buf)

	if err != nil {
		return "", errors.New("Ошибка чтения файла переменной-окружения, расположенного по пути \"" + pth + "\": " + err.Error())
	}

	return string(buf), nil
}

func ReadDir(dir string) (map[string]string, error) {
	var result = make(map[string]string)

	err := checkPathExist(dir)

	if err != nil {
		return map[string]string{}, err
	}

	err = filepath.WalkDir(dir, func(pth string, info fs.DirEntry, err error) error {

		if info.IsDir() {
			return nil
		}

		matched, er := checkFileName(info.Name(), re_list)

		if er != nil {
			return er
		}

		if !matched {
			return nil
		}

		fPath, er := filepath.Abs(pth)

		if er != nil {
			return errors.New("Ошибка получения абсолютного пути файла: " + er.Error())
		}

		valueFromFile, er := getFileContent(fPath)

		if er != nil {
			return er
		}

		result[info.Name()] = valueFromFile
		return nil
	})

	if err != nil {
		return nil, err
	}

	return result, nil
}
