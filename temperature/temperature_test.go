package temperature

import (
	"errors"
	"math"
	"strconv"
	"strings"
	"testing"
)

func TestParseCelsius(t *testing.T) {
	tests := []struct {
		name    string
		raw     string
		want    int16
		wantErr error
	}{
		{
			name:    "ноль",
			raw:     "0",
			want:    0,
			wantErr: nil,
		},
		{
			name:    "обычное положительное значение",
			raw:     "25",
			want:    25,
			wantErr: nil,
		},
		{
			name:    "отрицательное значение",
			raw:     "-25",
			want:    -25,
			wantErr: nil,
		},
		{
			name:    "нижняя граница int16",
			raw:     "-32768",
			want:    -32768,
			wantErr: nil,
		},
		{
			name:    "верхняя граница int16",
			raw:     "32767",
			want:    32767,
			wantErr: nil,
		},
		{
			name:    "переполнение сверху (> 32767)",
			raw:     "32768",
			want:    0,
			wantErr: strconv.ErrRange,
		},
		{
			name:    "переполнение снизу (< -32768)",
			raw:     "-32769",
			want:    0,
			wantErr: strconv.ErrRange,
		},
		{
			name:    "некорректная строка (буквы)",
			raw:     "cold",
			want:    0,
			wantErr: strconv.ErrSyntax,
		},
		{
			name:    "некорректная строка (дробное число)",
			raw:     "25.5",
			want:    0,
			wantErr: strconv.ErrSyntax,
		},
		{
			name:    "пустая строка",
			raw:     "",
			want:    0,
			wantErr: strconv.ErrSyntax,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ParseCelsius(tt.raw)

			// Кейс 1: Ожидали ошибку (tt.wantErr != nil)
			if tt.wantErr != nil {
				// Проверяем, что ошибка действительно возникла и распаковывается в tt.wantErr через errors.Is
				if !errors.Is(err, tt.wantErr) {
					t.Fatalf("ParseCelsius(%q) ждали ошибку %v, но получили: %v", tt.raw, tt.wantErr, err)
				}
				// Проверяем, что текст ошибки содержит исходную строку raw
				if tt.raw != "" && !strings.Contains(err.Error(), tt.raw) {
					t.Errorf("сообщение об ошибке %q не содержит исходную строку %q", err.Error(), tt.raw)
				}
				return
			}

			// Кейс 2: Ошибки быть не должно (tt.wantErr == nil)
			if err != nil {
				t.Fatalf("ParseCelsius(%q) неожиданная ошибка: %v", tt.raw, err)
			}
			if got != tt.want {
				t.Errorf("ParseCelsius(%q) got = %v, want %v", tt.raw, got, tt.want)
			}
		})
	}
}

func TestToFahrenheit(t *testing.T) {
	const epsilon = 1e-9

	tests := []struct {
		name string
		c    int16
		want float64
	}{
		{
			name: "ноль",
			c:    0,
			want: 32.0,
		},
		{
			name: "обычное положительное значение (25 C)",
			c:    25,
			want: 77.0,
		},
		{
			name: "отрицательное значение (-40 C == -40 F)",
			c:    -40,
			want: -40.0,
		},
		{
			name: "нижняя граница int16 (-32768)",
			c:    -32768,
			want: -58950.4,
		},
		{
			name: "верхняя граница int16 (32767)",
			c:    32767,
			want: 59012.6,
		},
		// --- Кейсы, где из-за погрешности float64 результат вычисления ToFahrenheit отличается от точного десятичного числа ---
		{
			name: "погрешность float64: -158 C (ожидаем -252.4, функция возвращает -252.39999999999998)",
			c:    -158,
			want: -252.4,
		},
		{
			name: "погрешность float64: -157 C (ожидаем -250.6, функция возвращает -250.60000000000002)",
			c:    -157,
			want: -250.6,
		},
		{
			name: "погрешность float64: -86 C (ожидаем -122.8, функция возвращает -122.80000000000001)",
			c:    -86,
			want: -122.8,
		},
		{
			name: "погрешность float64: -84 C (ожидаем -119.2, функция возвращает -119.19999999999999)",
			c:    -84,
			want: -119.2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ToFahrenheit(tt.c)
			if math.Abs(got-tt.want) > epsilon {
				t.Errorf("ToFahrenheit(%d) = %v, want %v (разница превышает epsilon %v)", tt.c, got, tt.want, epsilon)
			}
		})
	}
}
