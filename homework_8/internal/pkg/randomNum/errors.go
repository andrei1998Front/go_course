package randomNum

type ErrorMaxValue struct{}

func (err ErrorMaxValue) Error() string {
	return "максимальное значение не может быть меньше минимального"
}
