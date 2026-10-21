package _error

import (
	"errors"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestValidationError(t *testing.T) {
	t.Run("проверка метода Error и nil ресивера", func(t *testing.T) {
		valErr := &ValidationError{Field: "email", Reason: "empty"}
		var nilErr *ValidationError

		assert.Equal(t, "email: empty", valErr.Error())
		assert.Equal(t, "<nil>", nilErr.Error())
	})

	t.Run("проверка errors.Is с исходным экземпляром", func(t *testing.T) {
		original := &ValidationError{Field: "email", Reason: "empty"}
		wrapped := fmt.Errorf("register user: %w", original)

		assert.ErrorIs(t, wrapped, original)
	})

	t.Run("проверка errors.As и извлечение полей", func(t *testing.T) {
		original := &ValidationError{Field: "email", Reason: "empty"}
		wrapped := fmt.Errorf("register user: %w", original)

		var target *ValidationError
		ok := errors.As(wrapped, &target)

		require.True(t, ok)
		require.NotNil(t, target)
		assert.Equal(t, "email", target.Field)
		assert.Equal(t, "empty", target.Reason)
	})

	t.Run("проверка errors.Is с другим экземпляром с теми же полями", func(t *testing.T) {
		original := &ValidationError{Field: "email", Reason: "empty"}
		wrapped := fmt.Errorf("register user: %w", original)
		other := &ValidationError{Field: "email", Reason: "empty"}

		assert.False(t, errors.Is(wrapped, other))
	})
}

func TestHandle(t *testing.T) {
	t.Run("handle возвращает обернутую ошибку валидации", func(t *testing.T) {
		err := handle()

		require.Error(t, err)

		var target *ValidationError
		ok := errors.As(err, &target)

		require.True(t, ok)
		assert.Equal(t, "email", target.Field)
		assert.Equal(t, "invalid email", target.Reason)
	})
}

func TestTypedNilError(t *testing.T) {
	t.Run("демонстрация и предотвращение typed nil ловушки", func(t *testing.T) {
		// Корректный возврат без ошибки
		safeFunc := func() error {
			return nil
		}
		err := safeFunc()

		assert.Nil(t, err)
		assert.True(t, err == nil)

		// Опасный возврат типизированного nil
		unsafeFunc := func() error {
			var valErr *ValidationError = nil
			return valErr
		}
		badErr := unsafeFunc()

		assert.False(t, badErr == nil)
	})
}
