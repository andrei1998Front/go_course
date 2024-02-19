package readdir

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCheckfileName(t *testing.T) {
	var TestTable = []struct {
		re_list      []string
		fileName     string
		boolExpected bool
		expectedErr  string
	}{
		{
			re_list: []string{
				".+_ENV$",
				".+_VAR$",
			},
			fileName:     "ABC_ENV",
			boolExpected: true,
			expectedErr:  "",
		},
		{
			re_list: []string{
				".+_ENV$",
				".+_VAR$",
			},
			fileName:     "ABC_VAR",
			boolExpected: true,
			expectedErr:  "",
		},
		{
			re_list: []string{
				".+_ENV$",
				".+_VAR$",
			},
			fileName:     "abc_var",
			boolExpected: false,
			expectedErr:  "",
		},
		{
			re_list: []string{
				".+_ENV$",
				".+_VAR$",
			},
			fileName:     "abc",
			boolExpected: false,
			expectedErr:  "",
		},
		{
			re_list:      []string{},
			fileName:     "abc",
			boolExpected: false,
			expectedErr:  "массив регулярных выражений не может быть пустым",
		},
	}

	for _, testCase := range TestTable {
		result, err := checkFileName(testCase.fileName, testCase.re_list)

		if err != nil {
			require.ErrorContains(t, err, testCase.expectedErr)
			require.Equal(t, testCase.boolExpected, result)
			continue
		}
		t.Log(testCase.re_list)
		t.Log(testCase.fileName)
		t.Log(testCase.boolExpected)
		require.Equal(t, testCase.boolExpected, result)
	}
}

func TestGetFileContent(t *testing.T) {
	var TestTable = []struct {
		pth             string
		expectedContent string
		expectedError   string
	}{
		{
			pth:             "C:\\Users\\Андрей\\go\\src\\homework_7\\dir_for_test\\A_VAR",
			expectedContent: "123",
			expectedError:   "",
		},
		{
			pth:             "ffff",
			expectedContent: "",
			expectedError:   "Ошибка открытия файла переменной-окружения, расположенного по пути",
		},
	}

	for _, testCase := range TestTable {
		result, err := getFileContent(testCase.pth)

		if err != nil {
			require.ErrorContains(t, err, testCase.expectedError)
			require.Equal(t, testCase.expectedContent, result)
			continue
		}
		t.Log(testCase.pth)
		t.Log(testCase.expectedContent)
		require.Equal(t, testCase.expectedContent, result)
	}
}

func TestCheckPathExist(t *testing.T) {
	var TestTable = []struct {
		pth         string
		expectedErr string
	}{
		{
			pth:         "ggg",
			expectedErr: "Ошибка чтения пути:",
		},
		{
			pth:         "C:\\Users\\Андрей\\go\\src\\homework_7\\",
			expectedErr: "",
		},
	}

	for _, testCase := range TestTable {
		err := checkPathExist(testCase.pth)
		t.Log(testCase.pth)
		t.Log(testCase.expectedErr)
		t.Log(err)
		if err != nil {
			require.ErrorContains(t, err, testCase.expectedErr)
			continue
		}
	}
}

func TestReadDir(t *testing.T) {
	var TestTable = []struct {
		dir           string
		expectedMap   map[string]string
		expectedError string
	}{
		{
			dir: "C:\\Users\\Андрей\\go\\src\\homework_7\\",
			expectedMap: map[string]string{
				"A_VAR": "123",
				"B_ENV": "ccgvb",
			},
			expectedError: "",
		},
		{
			dir:           "C:\\Users\\Андрей\\go\\src\\homework_6",
			expectedMap:   map[string]string{},
			expectedError: "",
		},
		{
			dir:           "fff",
			expectedMap:   map[string]string{},
			expectedError: "Ошибка чтения пути:",
		},
	}

	for _, testCase := range TestTable {
		result, err := ReadDir(testCase.dir)
		t.Log(testCase.dir)
		t.Log(testCase.expectedMap)
		t.Log(err)
		if err != nil {
			require.ErrorContains(t, err, testCase.expectedError)
			require.Equal(t, testCase.expectedMap, result)
			continue
		}

		require.Equal(t, testCase.expectedMap, result)
	}
}
