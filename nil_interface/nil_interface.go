package nil_interface

import "reflect"

type ValidationError struct {
	Field string
}

func (e *ValidationError) Error() string {
	return "invalid field: " + e.Field
}

// ValidateName содержит исходную реализацию с ошибкой возврата типизированного nil
func ValidateName(name string) error {
	var err *ValidationError
	if name == "" {
		err = &ValidationError{Field: "name"}
	}
	return err // Возвращает пару (*ValidationError, nil), из-за чего interface != nil
}

// ValidateNameFixed реализует правильный контракт возврата ошибки:
// при валидном значении возвращается явный нетипизированный литерал nil
func ValidateNameFixed(name string) error {
	if name == "" {
		return &ValidationError{Field: "name"}
	}
	return nil
}

// IsNil проверяет значение на nil, включая typed nil для ссылочных типов
// Вызов rv.IsNil() безопасен только для nil-able kind
func IsNil(v any) bool {
	if v == nil {
		return true
	}

	rv := reflect.ValueOf(v)
	switch rv.Kind() {
	case reflect.Chan, reflect.Func, reflect.Interface, reflect.Map, reflect.Pointer, reflect.Slice, reflect.UnsafePointer:
		return rv.IsNil()
	default:
		return false
	}
}
