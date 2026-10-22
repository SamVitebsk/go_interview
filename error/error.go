package _error

import "fmt"

type ValidationError struct {
	Field  string
	Reason string
}

// Pointer receiver нужен для identity-сравнения экземпляров в errors.Is и исключения копирования структуры
func (e *ValidationError) Error() string {
	if e == nil {
		return "<nil>"
	}
	return e.Field + ": " + e.Reason
}

var _ error = (*ValidationError)(nil)

func handle() error {
	validationError := &ValidationError{
		Field:  "email",
		Reason: "invalid email",
	}

	return fmt.Errorf("handle: %w", validationError)
}
