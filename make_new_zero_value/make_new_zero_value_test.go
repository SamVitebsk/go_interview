package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBag(t *testing.T) {
	t.Run("чтение до первой записи (нулевое значение)", func(t *testing.T) {
		t.Parallel()

		var b Bag

		assert.Equal(t, 0, b.Len())

		testStrings := []string{"Go", "Rust", "", "any"}
		for _, str := range testStrings {
			assert.False(t, b.Contains(str))
		}
	})

	t.Run("первое добавление, повторы, пустая строка и разные строки", func(t *testing.T) {
		t.Parallel()

		var b Bag

		// Первое добавление
		b.Add("Go")
		assert.Equal(t, 1, b.Len())
		assert.True(t, b.Contains("Go"))

		// Повторное добавление того же элемента не увеличивает размер
		b.Add("Go")
		assert.Equal(t, 1, b.Len())

		// Добавление пустой строки
		b.Add("")
		assert.Equal(t, 2, b.Len())
		assert.True(t, b.Contains(""))
		assert.True(t, b.Contains("Go"))

		// Добавление нескольких разных строк
		otherLanguages := []string{"Rust", "Python", "Java"}
		for _, lang := range otherLanguages {
			b.Add(lang)
		}

		assert.Equal(t, 5, b.Len())

		allAdded := []string{"Go", "", "Rust", "Python", "Java"}
		for _, item := range allAdded {
			assert.True(t, b.Contains(item))
		}

		// Элемент, который не добавлялся
		assert.False(t, b.Contains("C++"))
	})

	t.Run("два отдельно объявленных var a, b Bag не разделяют элементы", func(t *testing.T) {
		t.Parallel()

		var a, b Bag

		a.Add("Go")

		// Набор a содержит добавленный элемент
		assert.Equal(t, 1, a.Len())
		assert.True(t, a.Contains("Go"))

		// Набор b остаётся нетронутым и пустым
		assert.Equal(t, 0, b.Len())
		assert.False(t, b.Contains("Go"))
	})

	t.Run("одинаковое поведение методов для var b Bag и new(Bag)", func(t *testing.T) {
		t.Parallel()

		var b1 Bag
		b2 := new(Bag)

		// Чтение до первой записи
		assert.Equal(t, b1.Len(), b2.Len())
		assert.Equal(t, b1.Contains("Go"), b2.Contains("Go"))

		// Добавление элементов работает идентично
		b1.Add("Go")
		b2.Add("Go")

		assert.Equal(t, 1, b1.Len())
		assert.Equal(t, 1, b2.Len())
		assert.True(t, b2.Contains("Go"))
	})
}
