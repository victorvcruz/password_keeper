package errors

import "fmt"

type UnauthorizedTokenError struct {
	Path string
}

func (e *UnauthorizedTokenError) Error() string {
	return fmt.Sprintf("Unauthorized token" + e.Path)
}
