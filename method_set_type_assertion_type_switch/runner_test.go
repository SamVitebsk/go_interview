package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRunner_Execute(t *testing.T) {
	t.Run("вызов через интерфейс для типа с value receiver", func(t *testing.T) {
		r := ValueRunner{}
		got := Execute(r)
		want := "ValueRunner"

		assert.Equal(t, want, got)
	})

	t.Run("вызов через интерфейс для типа с pointer receiver", func(t *testing.T) {
		r := &PointerRunner{}
		got := Execute(r)
		want := "PointerRunner"

		assert.Equal(t, want, got)
	})

	t.Run("разница в наборах методов value и pointer receiver", func(t *testing.T) {
		var r1 Runner = ValueRunner{}
		var r2 Runner = &ValueRunner{}

		assert.Equal(t, "ValueRunner", r1.Run())
		assert.Equal(t, "ValueRunner", r2.Run())

		var r3 Runner = &PointerRunner{}

		assert.Equal(t, "PointerRunner", r3.Run())
	})
}

func TestDescribe(t *testing.T) {
	t.Run("диагностика *PointerRunner", func(t *testing.T) {
		p := &PointerRunner{}
		got := Describe(p)
		want := "PointerRunner"

		assert.Equal(t, want, got)
	})

	t.Run("диагностика ValueRunner", func(t *testing.T) {
		v := ValueRunner{}
		got := Describe(v)
		want := "ValueRunner"

		assert.Equal(t, want, got)
	})

	t.Run("диагностика неизвестного типа возвращает unknown", func(t *testing.T) {
		want := "unknown"

		assert.Equal(t, want, Describe("some string"))
		assert.Equal(t, want, Describe(123))
		assert.Equal(t, want, Describe(nil))
	})
}
