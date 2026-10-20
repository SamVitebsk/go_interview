package temperature

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
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

			if tt.wantErr != nil {
				assert.ErrorIs(t, err, tt.wantErr)
				if tt.raw != "" {
					assert.Contains(t, err.Error(), tt.raw)
				}
				return
			}

			require.NoError(t, err)
			assert.Equal(t, tt.want, got)
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
			assert.InDelta(t, tt.want, got, epsilon)
		})
	}
}
