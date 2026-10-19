package main

import (
	"slices"
	"testing"
)

// TestDeferOrder проверяет обратный порядок выполнения трех defer (LIFO)
// без зависимости от вывода в консоль (Требование №4).
func TestDeferOrder(t *testing.T) {
	var got []int

	func() {
		defer func() { got = append(got, 1) }()
		defer func() { got = append(got, 2) }()
		defer func() { got = append(got, 3) }()
	}()

	want := []int{3, 2, 1}
	if !slices.Equal(got, want) {
		t.Errorf("got %v, want %v", got, want)
	}
}
