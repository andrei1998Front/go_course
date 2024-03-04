package event

type ErrNonExistentID struct{}

func (e ErrNonExistentID) Error() string {
	return "события с таким идентификатором не существует"
}

type ErrNonExistentEvent struct{}

func (e ErrNonExistentEvent) Error() string {
	return "такого события не существует"
}

type ErrDateBusy struct{}

func (e ErrDateBusy) Error() string {
	return "событие на данную дату уже выбрано"
}

type ErrNonExistentDate struct{}

func (e ErrNonExistentDate) Error() string {
	return "события с такой датой не существует"
}

type ErrExistentID struct{}

func (e ErrExistentID) Error() string {
	return "события с таким идентификатором уже существует"
}
