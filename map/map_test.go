package _map

import (
	"maps"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestFrequencies(t *testing.T) {
	t.Run("nil срез возвращает пустую карту", func(t *testing.T) {
		t.Parallel()

		got := Frequencies(nil)

		assert.NotNil(t, got)
		assert.Empty(t, got)
	})

	t.Run("пустой срез возвращает пустую карту", func(t *testing.T) {
		t.Parallel()

		got := Frequencies([]string{})

		assert.NotNil(t, got)
		assert.Empty(t, got)
	})

	t.Run("подсчёт частот с повторами", func(t *testing.T) {
		t.Parallel()

		words := []string{"яблоко", "банан", "яблоко", "апельсин", "банан", "яблоко"}
		want := map[string]int{
			"яблоко":   3,
			"банан":    2,
			"апельсин": 1,
		}

		got := Frequencies(words)

		assert.Equal(t, want, got)
	})

	t.Run("все слова уникальны", func(t *testing.T) {
		t.Parallel()

		words := []string{"golang", "rust", "python"}
		want := map[string]int{
			"golang": 1,
			"rust":   1,
			"python": 1,
		}

		got := Frequencies(words)

		assert.Equal(t, want, got)
	})
}

func TestMostFrequent(t *testing.T) {
	t.Run("nil карта возвращает ok = false", func(t *testing.T) {
		t.Parallel()

		word, count, ok := MostFrequent(nil)

		assert.False(t, ok)
		assert.Empty(t, word)
		assert.Zero(t, count)
	})

	t.Run("пустая карта возвращает ok = false", func(t *testing.T) {
		t.Parallel()

		word, count, ok := MostFrequent(map[string]int{})

		assert.False(t, ok)
		assert.Empty(t, word)
		assert.Zero(t, count)
	})

	t.Run("единственное слово", func(t *testing.T) {
		t.Parallel()

		freq := map[string]int{"единственное": 7}
		word, count, ok := MostFrequent(freq)

		require.True(t, ok)
		assert.Equal(t, "единственное", word)
		assert.Equal(t, 7, count)
	})

	t.Run("явный максимум частоты без ничьей", func(t *testing.T) {
		t.Parallel()

		freq := map[string]int{
			"кот":    1,
			"собака": 5,
			"птица":  3,
		}

		word, count, ok := MostFrequent(freq)

		require.True(t, ok)
		assert.Equal(t, "собака", word)
		assert.Equal(t, 5, count)
	})

	t.Run("равная максимальная частота: лексикографически меньшее слово", func(t *testing.T) {
		t.Parallel()

		freq := map[string]int{
			"zebra":  4,
			"apple":  4,
			"banana": 4,
			"cat":    2,
		}

		// Запускаем многократно, чтобы подтвердить независимость от случайного порядка обхода map
		for range 100 {
			word, count, ok := MostFrequent(freq)

			require.True(t, ok)
			assert.Equal(t, "apple", word)
			assert.Equal(t, 4, count)
		}
	})

	t.Run("равная максимальная частота: лексикографически меньшее слово (кириллица)", func(t *testing.T) {
		t.Parallel()

		freq := map[string]int{
			"яблоко": 10,
			"арбуз":  10,
			"груша":  10,
			"вишня":  5,
		}

		for range 100 {
			word, count, ok := MostFrequent(freq)

			require.True(t, ok)
			assert.Equal(t, "арбуз", word)
			assert.Equal(t, 10, count)
		}
	})

	t.Run("входная карта не изменяется", func(t *testing.T) {
		t.Parallel()

		freq := map[string]int{
			"first":  3,
			"second": 3,
			"third":  1,
		}
		snapshot := maps.Clone(freq)

		_, _, ok := MostFrequent(freq)

		require.True(t, ok)
		assert.Equal(t, snapshot, freq)
	})
}
