package firstdayofmonth

type ErrMonthValue struct{}

func (err ErrMonthValue) Error() string {
	return "неверное значение месяца. Значение месяца должно быть в диапозоне от 1 до 12"
}

type ErrYearValue struct{}

func (err ErrYearValue) Error() string {
	return "неверно значение года. Должно быть не меньшу нуля"
}
