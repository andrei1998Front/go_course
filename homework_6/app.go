package main

import (
	"errors"
	"flag"
	"io"
	"log"
	"os"
	"strconv"

	"github.com/cheggaaa/pb/v3"
)

type flags struct {
	from   string
	to     string
	limit  string
	offset string
}

type flgsToFunc struct {
	from   string
	to     string
	limit  int64
	offset int64
}

func setLimit(fl *os.File, currentLimit int64, offset int64) (int64, error) {
	stat, err := fl.Stat()

	if err != nil {
		return 0, err
	}

	endOfFile := stat.Size() - offset

	if currentLimit < 0 {
		return 0, errors.New("лимит не может быть отрицательным")
	} else if offset < 0 {
		return 0, errors.New("сдвиг не может быть отрицательным")
	} else if offset > stat.Size() {
		return 0, errors.New("сдвиг не может быть больше чем размер файла")
	} else if currentLimit == 0 || currentLimit > endOfFile {
		return endOfFile, nil
	}

	return currentLimit, nil
}

func Copy(from string, to string, limit int64, offset int64) error {
	if offset < 0 || limit < 0 {
		return errors.New("аргументы функции не могут быть отрицательными")
	}

	fileFrom, err := os.Open(from)

	if err != nil {
		return err
	}

	fileStat, err := fileFrom.Stat()

	if err != nil {
		return err
	}

	if fileStat.Size() < offset {
		return errors.New("значение сдвига не должно превышать размер файла")
	}

	fileTo, err := os.Create(to)

	if err != nil {
		return err
	}

	limit, err = setLimit(fileFrom, limit, offset)

	if err != nil {
		return err
	}

	bar := pb.Full.Start64(limit)

	defer func() {
		fileFrom.Close()
		fileTo.Close()
		bar.Finish()
	}()

	var buff = make([]byte, limit)

	if err != nil {
		return err
	}

	reader := io.NewSectionReader(fileFrom, offset, limit)
	barReader := bar.NewProxyReader(reader)

	_, err = barReader.Read(buff)

	if err != nil {
		return err
	}

	_, err = fileTo.Write(buff)

	if err != nil {
		return err
	}

	return nil
}

func parseFlags() {
	flag.String("from", "./for_test.txt", "Файл-источник")
	flag.String("to", "./test_file.txt", "Файл в который будет копироваться")
	flag.Int64("limit", 0, "Ограничение по количеству копируемых байт")
	flag.Int64("offset", 0, "Сдвиг файла-источника")

	flag.Parse()
}

func getFlags() flags {
	flgs := flags{}

	flag.VisitAll(func(f *flag.Flag) {
		switch f.Name {
		case "from":
			flgs.from = f.Value.String()

		case "to":
			flgs.to = f.Value.String()

		case "limit":
			flgs.limit = f.Value.String()

		case "offset":
			flgs.offset = f.Value.String()
		}
	})

	return flgs
}

func getFlagsToFunc(f *flags) (flgsToFunc, error) {
	var resultFlags flgsToFunc

	resultFlags.from = f.from
	resultFlags.to = f.to

	limit, err := strconv.ParseInt(f.limit, 10, 64)

	if err != nil {
		return flgsToFunc{}, err
	}

	resultFlags.limit = limit

	offset, err := strconv.ParseInt(f.offset, 10, 64)

	if err != nil {
		return flgsToFunc{}, err
	}

	resultFlags.offset = offset

	return resultFlags, nil
}

func main() {
	parseFlags()

	flgs := getFlags()
	toFunc, err := getFlagsToFunc(&flgs)

	if err != nil {
		log.Fatal(err)
	}

	err = Copy(toFunc.from, toFunc.to, toFunc.limit, toFunc.offset)

	if err != nil {
		log.Fatal(err)
	}
}
