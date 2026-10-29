package deadlock_goroutine_leak_close_channel

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNumbers_FiveElementsAndCancellation(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	numbers := Numbers(ctx)

	want := []int{0, 1, 2, 3, 4}

	var got []int
	for range 5 {
		val, ok := <-numbers
		assert.True(t, ok)
		got = append(got, val)
	}

	assert.Equal(t, want, got)

	cancel()

	val, ok := <-numbers
	assert.False(t, ok)
	assert.Equal(t, 0, val)
}
