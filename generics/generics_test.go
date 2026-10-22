package generics

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type UserKey struct {
	ID     int
	Region string
}

func TestSet_Int(t *testing.T) {
	t.Run("добавление, проверка наличия и удаление элементов int", func(t *testing.T) {
		s := make(Set[int])

		s.Add(10)
		s.Add(20)

		assert.True(t, s.Contains(10))
		assert.True(t, s.Contains(20))
		assert.False(t, s.Contains(30))
		assert.Equal(t, 2, s.Len())

		s.Add(10)

		assert.Equal(t, 2, s.Len())

		s.Delete(10)

		assert.False(t, s.Contains(10))
		assert.Equal(t, 1, s.Len())
	})
}

func TestSet_String(t *testing.T) {
	t.Run("работа множества со строками string", func(t *testing.T) {
		s := make(Set[string])

		s.Add("golang")
		s.Add("rust")

		assert.True(t, s.Contains("golang"))
		assert.True(t, s.Contains("rust"))
		assert.False(t, s.Contains("python"))
		assert.Equal(t, 2, s.Len())

		s.Delete("golang")

		assert.False(t, s.Contains("golang"))
		assert.Equal(t, 1, s.Len())
	})
}

func TestSet_ComparableStruct(t *testing.T) {
	t.Run("работа множества со сравнимой структурой", func(t *testing.T) {
		s := make(Set[UserKey])

		k1 := UserKey{ID: 1, Region: "EU"}
		k2 := UserKey{ID: 2, Region: "US"}

		s.Add(k1)
		s.Add(k2)

		assert.True(t, s.Contains(k1))
		assert.True(t, s.Contains(k2))
		assert.False(t, s.Contains(UserKey{ID: 3, Region: "ASIA"}))
		assert.Equal(t, 2, s.Len())

		s.Add(UserKey{ID: 1, Region: "EU"})

		assert.Equal(t, 2, s.Len())

		s.Delete(k1)

		assert.False(t, s.Contains(k1))
		assert.Equal(t, 1, s.Len())
	})
}

func TestSet_NilSafety(t *testing.T) {
	t.Run("безопасность методов Contains, Delete и Len для nil set", func(t *testing.T) {
		var s Set[int] = nil

		assert.False(t, s.Contains(10))
		assert.Equal(t, 0, s.Len())

		assert.NotPanics(t, func() {
			s.Delete(10)
		})
	})
}

func TestUnique(t *testing.T) {
	t.Run("сохранение порядка первого появления элементов", func(t *testing.T) {
		input := []int{2, 1, 2}
		want := []int{2, 1}
		got := Unique(input)

		assert.Equal(t, want, got)
		assert.Equal(t, []int{2, 1, 2}, input)
	})

	t.Run("nil вход дает результат длины 0", func(t *testing.T) {
		var input []string = nil
		got := Unique(input)

		require.NotNil(t, got)
		assert.Len(t, got, 0)
	})

	t.Run("пустой вход дает результат длины 0", func(t *testing.T) {
		input := []int{}
		got := Unique(input)

		require.NotNil(t, got)
		assert.Len(t, got, 0)
	})

	t.Run("уникализация элементов сравнимой структуры", func(t *testing.T) {
		input := []UserKey{
			{ID: 1, Region: "RU"},
			{ID: 2, Region: "KZ"},
			{ID: 1, Region: "RU"},
		}
		want := []UserKey{
			{ID: 1, Region: "RU"},
			{ID: 2, Region: "KZ"},
		}
		got := Unique(input)

		assert.Equal(t, want, got)
	})
}
