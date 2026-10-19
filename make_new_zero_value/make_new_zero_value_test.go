package main

import "testing"

func TestBag(t *testing.T) {
	t.Run("чтение до первой записи (нулевое значение)", func(t *testing.T) {
		t.Parallel()

		var b Bag

		if got := b.Len(); got != 0 {
			t.Errorf("Len() для нулевого значения = %v, want 0", got)
		}

		testStrings := []string{"Go", "Rust", "", "any"}
		for _, str := range testStrings {
			if b.Contains(str) {
				t.Errorf("Contains(%q) для нулевого значения = true, want false", str)
			}
		}
	})

	t.Run("первое добавление, повторы, пустая строка и разные строки", func(t *testing.T) {
		t.Parallel()

		var b Bag

		// Первое добавление
		b.Add("Go")
		if b.Len() != 1 {
			t.Fatalf("после добавления первого элемента Len() = %v, want 1", b.Len())
		}
		if !b.Contains("Go") {
			t.Errorf("Contains(\"Go\") = false, want true")
		}

		// Повторное добавление того же элемента не увеличивает длину
		b.Add("Go")
		if b.Len() != 1 {
			t.Errorf("после повторного добавления \"Go\" Len() = %v, want 1", b.Len())
		}

		// Добавление пустой строки допустимо
		b.Add("")
		if b.Len() != 2 {
			t.Fatalf("после добавления пустой строки Len() = %v, want 2", b.Len())
		}
		if !b.Contains("") {
			t.Errorf("Contains(\"\") = false, want true")
		}

		// Проверяем, что предыдущий элемент сохранился
		if !b.Contains("Go") {
			t.Errorf("предыдущий элемент \"Go\" был потерян")
		}

		// Несколько разных строк
		otherLanguages := []string{"Rust", "Python", "Java"}
		for _, lang := range otherLanguages {
			b.Add(lang)
		}

		if b.Len() != 5 {
			t.Errorf("итоговая длина Len() = %v, want 5", b.Len())
		}

		allAdded := []string{"Go", "", "Rust", "Python", "Java"}
		for _, item := range allAdded {
			if !b.Contains(item) {
				t.Errorf("Contains(%q) = false, want true", item)
			}
		}

		// Элемент, который не добавлялся
		if b.Contains("C++") {
			t.Errorf("Contains(\"C++\") = true, want false")
		}
	})

	t.Run("два отдельно объявленных var a, b Bag не разделяют элементы", func(t *testing.T) {
		t.Parallel()

		var a, b Bag

		a.Add("Go")

		// Набор a содержит элемент
		if a.Len() != 1 || !a.Contains("Go") {
			t.Errorf("набор a должен содержать \"Go\", len = %v", a.Len())
		}

		// Набор b должен остаться нетронутым и пустым
		if b.Len() != 0 {
			t.Errorf("набор b не должен измениться: Len() = %v, want 0", b.Len())
		}
		if b.Contains("Go") {
			t.Errorf("набор b содержит \"Go\", хотя добавляли только в a")
		}
	})

	t.Run("одинаковое поведение методов для var b Bag и new(Bag)", func(t *testing.T) {
		t.Parallel()

		var b1 Bag
		b2 := new(Bag)

		// Чтение до первой записи
		if b1.Len() != b2.Len() {
			t.Errorf("начальная длина b1 (%v) != b2 (%v)", b1.Len(), b2.Len())
		}
		if b1.Contains("Go") != b2.Contains("Go") {
			t.Errorf("Contains до записи отличается у b1 и b2")
		}

		// Добавление элементов работает идентично
		b1.Add("Go")
		b2.Add("Go")

		if b1.Len() != b2.Len() || b1.Len() != 1 {
			t.Errorf("длина после добавления отличается: b1=%v, b2=%v", b1.Len(), b2.Len())
		}
		if !b2.Contains("Go") {
			t.Errorf("b2.Contains(\"Go\") = false, want true")
		}
	})
}
