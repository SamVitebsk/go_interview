package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDeferOrder(t *testing.T) {
	var got []int

	func() {
		defer func() { got = append(got, 1) }()
		defer func() { got = append(got, 2) }()
		defer func() { got = append(got, 3) }()
	}()

	assert.Equal(t, []int{3, 2, 1}, got)
}
