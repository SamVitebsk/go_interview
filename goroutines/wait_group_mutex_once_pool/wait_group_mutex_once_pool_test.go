package wait_group_mutex_once_pool

import (
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAtomicCounter_WithWaitGroup(t *testing.T) {
	c := &Counter{}
	var wg sync.WaitGroup

	const goroutines = 100
	const incPerGoroutine = 10

	for range goroutines {
		wg.Go(func() {
			for range incPerGoroutine {
				c.Inc()
			}
		})
	}

	wg.Wait()
	assert.Equal(t, goroutines*incPerGoroutine, c.Value())
}

func TestAtomicCounter_AddWithWaitGroup(t *testing.T) {
	c := &Counter{}
	var wg sync.WaitGroup

	const goroutines = 50

	for range goroutines {
		wg.Go(func() {
			c.Add(10)
		})
	}

	wg.Wait()
	assert.Equal(t, goroutines*10, c.Value())
}

func TestMutexCounter_WithWaitGroup(t *testing.T) {
	c := &MutexCounter{}
	var wg sync.WaitGroup

	const goroutines = 100
	const incPerGoroutine = 10

	for range goroutines {
		wg.Go(func() {
			for range incPerGoroutine {
				c.Inc()
			}
		})
	}

	wg.Wait()
	assert.Equal(t, goroutines*incPerGoroutine, c.Value())
}

func TestMutexCounter_AddWithWaitGroup(t *testing.T) {
	c := &MutexCounter{}
	var wg sync.WaitGroup

	const goroutines = 50

	for range goroutines {
		wg.Go(func() {
			c.Add(10)
		})
	}

	wg.Wait()
	assert.Equal(t, goroutines*10, c.Value())
}
