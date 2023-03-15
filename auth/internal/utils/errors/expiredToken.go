package errors

import "fmt"

type ExpiredTokenError struct {
	Path string
}

func (e *ExpiredTokenError) Error() string {
	return fmt.Sprintf("expired token" + e.Path)
}
