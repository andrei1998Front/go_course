package readdir

import (
	"errors"
	"fmt"
	"io"
	"io/fs"
	"os"
	"path"
	"path/filepath"
	"regexp"
)

var re_list = []string{".+_ENV$", ".+_VAR$"}

func checkFileName(fileName string) (bool, error) {
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
	fmt.Println(buf)

	if err != nil {
		return "", errors.New("Ошибка чтения файла переменной-окружения, расположенного по пути \"" + pth + "\": " + err.Error())
	}

	return string(buf), nil
}

func ReadDir(dir string) (map[string]string, error) {
	var result = make(map[string]string)

	err := filepath.WalkDir(dir, func(pth string, info fs.DirEntry, err error) error {
		if info.IsDir() {
			return nil
		}

		matched, er := checkFileName(info.Name())

		if er != nil {
			return er
		}

		if !matched {
			return nil
		}

		fPath, er := filepath.Abs(path.Join(dir, pth))

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

	fmt.Println(result)
	return result, nil
}
