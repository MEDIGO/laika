package store

import (
	"fmt"
)

type CustomError struct {
	Message string
}

func (e CustomError) Error() string {
	return fmt.Sprintf("%v", e.Message)
}
