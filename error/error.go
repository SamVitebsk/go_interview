package _error

import "fmt"

/**

Сначала реализуйте handle, не используя пакеты errors и fmt:

Затем расширьте собственный тип ошибки полями Field и Reason типа string.
На первом этапе ограничение на пакеты проверяет понимание интерфейса error;
на этапе оборачивания пакеты fmt и errors уже разрешены.

Требования:

1) Создайте тип ошибки валидации с полями имени параметра и причины отказа.
2) Реализуйте Error() на подходящем receiver и объясните выбор.
3) Оборачивайте ошибку на границе бизнес-операции через fmt.Errorf и %w.
4) В тесте создайте original := &ValidationError{Field: "email", Reason: "empty"},
оберните именно этот экземпляр и проверьте errors.Is(wrapped, original).
Через errors.As извлеките ошибку и проверьте оба поля.
Отдельно проверьте, что новый экземпляр с теми же полями не совпадает через Is,
если пользовательский метод Is не реализован.
5) Не возвращайте typed nil как error.
6) Пришлите ссылку на коммит с выполненной работой.

*/

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
