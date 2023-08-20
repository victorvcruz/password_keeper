package errors

import "fmt"

type NotFoundIdError struct {
	Path string
}

func (e *NotFoundIdError) Error() string {
	return fmt.Sprintf("Incorrect id" + e.Path)
}
