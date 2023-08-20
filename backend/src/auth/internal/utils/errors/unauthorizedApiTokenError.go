package errors

import "fmt"

type UnauthorizedApiTokenError struct {
	Path string
}

func (e *UnauthorizedApiTokenError) Error() string {
	return fmt.Sprintf("Incorrect api-token" + e.Path)
}
