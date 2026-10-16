package temperature

import (
	"fmt"
	"strconv"
)

// ParseCelsius преобразует строковое представление температуры в int16.
func ParseCelsius(raw string) (int16, error) {
	intCelsius, err := strconv.ParseInt(raw, 10, 16)
	if err != nil {
		return 0, fmt.Errorf("parse %q as int16: %w", raw, err)
	}

	return int16(intCelsius), nil
}

// ToFahrenheit переводит градусы Цельсия в градусы Фаренгейта по формуле F = C × 9 / 5 + 32.
func ToFahrenheit(c int16) float64 {
	return float64(c)*9/5 + 32
}
