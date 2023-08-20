package errors

type NotFound struct {
	Msg string
}

func (e *NotFound) Error() string {
	return e.Msg
}
