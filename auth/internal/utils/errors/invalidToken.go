package errors

import "fmt"

type InvalidTokenError struct {
	Path string
}

func (e *InvalidTokenError) Error() string {
	return fmt.Sprintf("invalid token" + e.Path)
}
