package scope_closure_range

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestAccumulator(t *testing.T) {
	t.Run("последовательное прибавление чисел", func(t *testing.T) {
		t.Parallel()

		acc := Accumulator()
		assert.Equal(t, 2, acc(2))
		assert.Equal(t, 1, acc(-1))
		assert.Equal(t, 6, acc(5))
	})

	t.Run("независимость разных экземпляров аккумуляторов", func(t *testing.T) {
		t.Parallel()

		acc1 := Accumulator()
		acc2 := Accumulator()

		assert.Equal(t, 10, acc1(10))
		assert.Equal(t, 100, acc2(100))

		assert.Equal(t, 15, acc1(5))
		assert.Equal(t, 100, acc2(0))
		assert.Equal(t, 50, acc2(-50))
	})
}

func TestReaders(t *testing.T) {
	t.Run("базовый сценарий с возвратом значений", func(t *testing.T) {
		t.Parallel()

		readers := Readers([]int{10, 20})
		require.Len(t, readers, 2)

		assert.Equal(t, 10, readers[0]())
		assert.Equal(t, 20, readers[1]())

		// Повторные вызовы возвращают те же зафиксированные значения
		assert.Equal(t, 10, readers[0]())
		assert.Equal(t, 20, readers[1]())
	})

	t.Run("изменение входного среза не меняет результат функций", func(t *testing.T) {
		t.Parallel()

		input := []int{10, 20}
		readers := Readers(input)
		require.Len(t, readers, 2)

		input[0] = 999
		input[1] = 888

		assert.Equal(t, 10, readers[0]())
		assert.Equal(t, 20, readers[1]())
	})

	t.Run("nil вход возвращает срез длины 0", func(t *testing.T) {
		t.Parallel()

		readers := Readers(nil)
		assert.Empty(t, readers)
		assert.Len(t, readers, 0)
	})

	t.Run("пустой срез возвращает срез длины 0", func(t *testing.T) {
		t.Parallel()

		readers := Readers([]int{})
		assert.Empty(t, readers)
		assert.Len(t, readers, 0)
	})
}

// TestLoopScopeDifference демонстрирует разницу работы range-циклов с := и = без горутин.
// 1. При объявлении через := (Go 1.22+) каждая итерация создаёт новую переменную -> замыкания возвращают [1, 2, 3].
// 2. При присваивании через = внешняя переменная переиспользуется -> все замыкания видят последнее значение [3, 3, 3].
func TestLoopScopeDifference(t *testing.T) {
	values := []int{1, 2, 3}

	// 1. Цикл с := (новая переменная на каждой итерации)
	var readersNew []func() int
	for _, val := range values {
		readersNew = append(readersNew, func() int { return val })
	}
	assert.Equal(t, 1, readersNew[0]())
	assert.Equal(t, 2, readersNew[1]())
	assert.Equal(t, 3, readersNew[2]())

	// 2. Цикл с = (одна внешняя переменная переиспользуется)
	var val int
	var readersSame []func() int
	for _, val = range values {
		readersSame = append(readersSame, func() int { return val })
	}
	assert.Equal(t, 3, readersSame[0]())
	assert.Equal(t, 3, readersSame[1]())
	assert.Equal(t, 3, readersSame[2]())
}

// TestUserSliceModification демонстрирует разницу изменения срезов структур []User и указателей []*User.
// 1. []User: range копирует структуру по значению, поэтому для мутации элемента нужно обращаться по индексу: users[i].Name = "Bob".
// 2. []*User: range копирует указатель, но он ссылается на тот же объект в памяти, поэтому user.Name = "Bob" меняет оригинал.
func TestUserSliceModification(t *testing.T) {
	type User struct{ Name string }

	// 1. []User: исправление изменения через обращение по индексу
	users := []User{{Name: "Alice"}}
	for i := range users {
		users[i].Name = "Bob"
	}
	assert.Equal(t, "Bob", users[0].Name)

	// 2. []*User: копия указателя указывает на ту же структуру в памяти
	userPointers := []*User{{Name: "Alice"}}
	for _, user := range userPointers {
		user.Name = "Bob"
	}
	assert.Equal(t, "Bob", userPointers[0].Name)
}
