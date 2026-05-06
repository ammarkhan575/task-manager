package task

import (
	"errors"
	"fmt"
)

var (
	ErrNotFound     = errors.New("task not found")
	ErrInvalidInput = errors.New("invalid input")
)

type ValidationError struct {
	Field   string
	Message string
}

func (e *ValidationError) Error() string {
    return fmt.Sprintf("validation failed on %s: %s", e.Field, e.Message)
}
