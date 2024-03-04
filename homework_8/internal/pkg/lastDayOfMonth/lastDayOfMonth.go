package lastdayofmonth

import (
	"time"

	firstdayofmonth "github.com/andrei1998Front/go_course/homework_8/internal/pkg/firstDayOfMonth"
)

func GetLastDayOfMonth(yy, mm int) (time.Time, error) {
	firstDay, err := firstdayofmonth.GetFirstDayOfMonth(yy, mm)

	if err != nil {
		return time.Time{}, err
	}

	firstDayNextMonth := firstDay.AddDate(0, 1, 0)

	return firstDayNextMonth.AddDate(0, 0, -1), nil
}
