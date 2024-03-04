package firstdayofmonth

import "time"

func GetFirstDayOfMonth(yy, mm int) (time.Time, error) {
	if mm <= 0 || mm > 12 {
		return time.Time{}, ErrMonthValue{}
	}

	if yy < 0 {
		return time.Time{}, ErrYearValue{}
	}

	return time.Date(yy, time.Month(mm), 1, 0, 0, 0, 0, time.UTC), nil
}
