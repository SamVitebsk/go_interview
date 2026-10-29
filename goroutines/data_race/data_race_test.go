package data_race

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestCount тестирует функцию Count, содержащую намеренный data race.
// 1. При обычном запуске `go test` тест проходит и проверяет, что результат лежит в диапазоне (0, 100].
// 2. При запуске `go test -race` Race Detector обнаруживает состояние гонки и завершает тест с ошибкой "WARNING: DATA RACE".
func TestCount(t *testing.T) {
	res := Count()
	assert.True(t, res > 0 && res <= 100, "результат должен быть в диапазоне (0, 100]")
}

func TestCountWithMutex(t *testing.T) {
	for range 10 {
		assert.Equal(t, 100, CountWithMutex())
	}
}

func TestCountWithAtomic(t *testing.T) {
	for range 10 {
		assert.Equal(t, 100, CountWithAtomic())
	}
}
