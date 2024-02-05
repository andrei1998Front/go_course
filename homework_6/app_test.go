package main

import (
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

type TestFiles struct {
	fileFromPath *os.File
	fileFromSize int64
	fileToPath   string
}

const (
	fileToPath   = "./testFileTo.txt"
	fileFromPath = "./testFileFrom.txt"
	errorPath    = "./dfdgfdsgfdg.txt"
	testText     = string("Равным образом рамки и место обучения кадров представляет собой интересный эксперимент проверки существующих финансовых и административных условий. Соображения высшего порядка, а также выбранный нами инновационный путь влечет за собой процесс внедрения и модернизации форм воздействия. Дорогие друзья, дальнейшее развитие различных форм деятельности представляет собой интересный эксперимент проверки экономической целесообразности принимаемых решений. Задача организации, в особенности же дальнейшее развитие различных форм деятельности играет важную роль в формировании направлений прогрессивного развития? Таким образом, выбранный нами инновационный путь создаёт предпосылки качественно новых шагов для дальнейших направлений развития проекта? Повседневная практика показывает, что выбранный нами инновационный путь обеспечивает широкому кругу специалистов участие в формировании экономической целесообразности принимаемых решений. Дорогие друзья, повышение уровня гражданского сознания обеспечивает широкому кругу специалистов участие в формировании соответствующих условий активизации. Не следует, однако, забывать о том, что сложившаяся структура организации влечет за собой процесс внедрения и модернизации дальнейших направлений развития проекта. Задача организации, в особенности же социально-экономическое развитие позволяет выполнить важнейшие задания по разработке ключевых компонентов планируемого обновления. Практический опыт показывает, что повышение уровня гражданского сознания влечет за собой процесс внедрения и модернизации системы масштабного изменения ряда параметров. Разнообразный и богатый опыт выбранный нами инновационный путь требует от нас анализа системы обучения кадров, соответствующей насущным потребностям? Таким образом, сложившаяся структура...")
)

func prepareTestFile() (TestFiles, error) {
	fileFrom, err := os.Create(fileFromPath)
	// defer fileFrom.Close()

	if err != nil {
		return TestFiles{}, err
	}

	buff := []byte(testText)

	_, err = fileFrom.Write(buff)

	if err != nil {
		return TestFiles{}, err
	}

	fileStat, err := fileFrom.Stat()

	if err != nil {
		return TestFiles{}, err
	}

	return TestFiles{
		fileFromPath: fileFrom,
		fileToPath:   fileToPath,
		fileFromSize: fileStat.Size(),
	}, nil
}

func TestSetLimit(t *testing.T) {
	testFiles, err := prepareTestFile()

	if err != nil {
		t.FailNow()
	}

	TestTable := []struct {
		fl             *os.File
		currentLimit   int64
		offset         int64
		mustBeError    bool
		err            string
		expectedOutput int64
	}{
		{
			fl:             testFiles.fileFromPath,
			currentLimit:   -1,
			offset:         100,
			mustBeError:    true,
			err:            "лимит не может быть отрицательным",
			expectedOutput: 0,
		},
		{
			fl:             testFiles.fileFromPath,
			currentLimit:   100,
			offset:         -1,
			mustBeError:    true,
			err:            "сдвиг не может быть отрицательным",
			expectedOutput: 0,
		},
		{
			fl:             testFiles.fileFromPath,
			currentLimit:   -1,
			offset:         -1,
			mustBeError:    true,
			err:            "лимит не может быть отрицательным",
			expectedOutput: 0,
		},
		{
			fl:             testFiles.fileFromPath,
			currentLimit:   1000,
			offset:         0,
			mustBeError:    false,
			err:            "",
			expectedOutput: 1000,
		},
		{
			fl:             testFiles.fileFromPath,
			currentLimit:   0,
			offset:         0,
			mustBeError:    false,
			err:            "",
			expectedOutput: testFiles.fileFromSize,
		},
		{
			fl:             testFiles.fileFromPath,
			currentLimit:   0,
			offset:         10,
			mustBeError:    false,
			err:            "",
			expectedOutput: testFiles.fileFromSize - 10,
		},
		{
			fl:             testFiles.fileFromPath,
			currentLimit:   0,
			offset:         10,
			mustBeError:    false,
			err:            "",
			expectedOutput: testFiles.fileFromSize - 10,
		},
		{
			fl:             testFiles.fileFromPath,
			currentLimit:   testFiles.fileFromSize + 200,
			offset:         10,
			mustBeError:    false,
			err:            "",
			expectedOutput: testFiles.fileFromSize - 10,
		},
		{
			fl:             testFiles.fileFromPath,
			currentLimit:   testFiles.fileFromSize + 200,
			offset:         0,
			mustBeError:    false,
			err:            "",
			expectedOutput: testFiles.fileFromSize,
		},
	}

	for _, testCase := range TestTable {
		result, err := setLimit(testCase.fl, testCase.currentLimit, testCase.offset)

		if testCase.mustBeError {
			require.EqualError(t, err, testCase.err)
			require.Equal(t, testCase.expectedOutput, result)
			continue
		}

		require.NoError(t, err)
		require.Equal(t, testCase.expectedOutput, result)
	}

	testFiles.fileFromPath.Close()
}

func TestCopy(t *testing.T) {
	testFiles, err := prepareTestFile()

	if err != nil {
		t.FailNow()
	}

	TestTable := []struct {
		from        string
		to          string
		limit       int64
		offset      int64
		mustBeError bool
		err         string
	}{
		{
			from:        fileFromPath,
			to:          fileToPath,
			limit:       -1,
			offset:      100,
			mustBeError: true,
			err:         "аргументы функции не могут быть отрицательными",
		},
		{
			from:        fileFromPath,
			to:          fileToPath,
			limit:       100,
			offset:      -1,
			mustBeError: true,
			err:         "аргументы функции не могут быть отрицательными",
		},
		{
			from:        fileFromPath,
			to:          fileToPath,
			limit:       100,
			offset:      -1,
			mustBeError: true,
			err:         "аргументы функции не могут быть отрицательными",
		},
		{
			from:        fileFromPath,
			to:          fileToPath,
			limit:       0,
			offset:      testFiles.fileFromSize + 200,
			mustBeError: true,
			err:         "значение сдвига не должно превышать размер файла",
		},
		{
			from:        errorPath,
			to:          fileToPath,
			limit:       0,
			offset:      testFiles.fileFromSize + 200,
			mustBeError: true,
			err:         "open " + errorPath + ": The system cannot find the file specified.",
		},
	}

	for _, testCase := range TestTable {
		err := Copy(testCase.from, testCase.to, testCase.limit, testCase.offset)

		if testCase.mustBeError {
			require.EqualError(t, err, testCase.err)
			continue
		}

		require.NoError(t, err)
	}

	testFiles.fileFromPath.Close()
}
