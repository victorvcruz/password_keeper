package errors

import "fmt"

type NotFoundServiceError struct {
	Path string
}

func (e *NotFoundServiceError) Error() string {
	return fmt.Sprintf("Incorrect service " + e.Path)
}
