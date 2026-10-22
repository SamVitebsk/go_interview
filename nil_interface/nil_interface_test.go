package nil_interface

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestValidateName_TypedNilBug(t *testing.T) {
	t.Run("исходная функция возвращает typed nil при валидном имени", func(t *testing.T) {
		err := ValidateName("Alice")

		// Обычная проверка err != nil возвращает true, так как в интерфейс
		// упакована пара (*ValidationError, nil). Динамический тип не равен nil!
		// assert.NotNil из testify использует reflection IsNil() под капотом,
		// а стандартная языковая проверка Go err != nil оценивается как true.
		assert.True(t, err != nil, "языковая проверка err != nil вычисляется в true")
		assert.False(t, err == nil, "интерфейс с typed nil не равен nil")
	})
}

func TestValidateNameFixed(t *testing.T) {
	t.Run("успешная валидация непустого имени возвращает чистый nil", func(t *testing.T) {
		err := ValidateNameFixed("Alice")

		require.NoError(t, err)
		assert.Nil(t, err)
		assert.True(t, err == nil)
	})

	t.Run("пробелы считаются обычными символами и валидация успешна", func(t *testing.T) {
		err := ValidateNameFixed("   ")

		require.NoError(t, err)
		assert.Nil(t, err)
		assert.True(t, err == nil)
	})

	t.Run("ошибка при пустом имени и распаковка через errors.As", func(t *testing.T) {
		err := ValidateNameFixed("")

		require.Error(t, err)

		var targetErr *ValidationError
		require.True(t, errors.As(err, &targetErr))
		assert.Equal(t, "name", targetErr.Field)
		assert.Equal(t, "invalid field: name", targetErr.Error())
	})
}

func TestIsNil(t *testing.T) {
	t.Run("нетипизированный nil дает true", func(t *testing.T) {
		var v any = nil

		assert.True(t, IsNil(v))
	})

	t.Run("типизированный nil указатель (*int)(nil) дает true", func(t *testing.T) {
		var p *int = nil

		assert.True(t, IsNil(p))
	})

	t.Run("nil slice дает true", func(t *testing.T) {
		var s []int = nil

		assert.True(t, IsNil(s))
	})

	t.Run("пустой ненулевой slice []int{} дает false", func(t *testing.T) {
		s := []int{}

		assert.False(t, IsNil(s))
	})

	t.Run("число 0 типа int дает false", func(t *testing.T) {
		n := 0

		assert.False(t, IsNil(n))
	})
}
