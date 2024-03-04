package randomeventslist

type ErrMaxMinYY struct{}

func (err ErrMaxMinYY) Error() string {
	return "максимальное значение года меньше минимального"
}

type ErrYYZero struct{}

func (err ErrYYZero) Error() string {
	return "значение года не может быть нулевым"
}

type ErrYYLessZero struct{}

func (err ErrYYLessZero) Error() string {
	return "значение года не может быть меньше нуля"
}

type ErrSizeZero struct{}

func (err ErrSizeZero) Error() string {
	return "размер слайса событий не может быть нулевым"
}

type ErrSizeLessZero struct{}

func (err ErrSizeLessZero) Error() string {
	return "размер слайса событий не может быть отрицательным"
}
