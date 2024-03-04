package randomeventslist

import (
	"math/rand"
	"strconv"
	"time"

	firstdayofmonth "github.com/andrei1998Front/go_course/homework_8/internal/pkg/firstDayOfMonth"
	lastdayofmonth "github.com/andrei1998Front/go_course/homework_8/internal/pkg/lastDayOfMonth"
)

type EventTeplate struct {
	Title     string
	DateEvent time.Time
}

func checkMaxMin(minYY, maxYY int) error {
	if maxYY < minYY {
		return ErrMaxMinYY{}
	}

	if maxYY == 0 || minYY == 0 {
		return ErrYYZero{}
	}

	if maxYY < 0 || minYY < 0 {
		return ErrYYLessZero{}
	}

	return nil
}

func GetRandomDate(minYY int, maxYY int) (time.Time, error) {
	err := checkMaxMin(minYY, maxYY)

	if err != nil {
		return time.Time{}, err
	}

	yy := rand.Intn(maxYY-minYY) + minYY
	mm := rand.Intn(11) + 1

	leftBorder, err := firstdayofmonth.GetFirstDayOfMonth(yy, mm)

	if err != nil {
		return time.Time{}, err
	}

	rightBorder, err := lastdayofmonth.GetLastDayOfMonth(yy, mm)

	if err != nil {
		return time.Time{}, err
	}

	dd := rand.Intn(rightBorder.Day()-leftBorder.Day()) - leftBorder.Day()

	dt := time.Date(yy, time.Month(mm), dd, 0, 0, 0, 0, time.UTC)

	return dt, nil
}

func GetRandomEventsList(size int, minYY, maxYY int) ([]EventTeplate, error) {
	err := checkMaxMin(minYY, maxYY)

	if err != nil {
		return []EventTeplate{}, err
	}

	if size < 0 {
		return []EventTeplate{}, ErrSizeLessZero{}
	}

	if size == 0 {
		return []EventTeplate{}, ErrSizeZero{}
	}

	eventTemplate := []EventTeplate{}

	for i := 0; i < size; i++ {
		title := "Event #" + strconv.Itoa(i)
		dt, err := GetRandomDate(minYY, maxYY)

		if err != nil {
			return []EventTeplate{}, err
		}

		et := EventTeplate{Title: title, DateEvent: dt}
		eventTemplate = append(eventTemplate, et)
	}

	return eventTemplate, nil
}
