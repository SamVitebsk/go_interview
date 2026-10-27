package escape_analysis_gc_memory

import (
	"fmt"
	"slices"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBuild_Equality(t *testing.T) {
	t.Run("при n = 0 содержимое эквивалентно (nil против пустого среза)", func(t *testing.T) {
		noPrealloc := Build(0, false)
		prealloc := Build(0, true)

		// Технически в памяти они различаются (nil vs non-nil empty slice):
		assert.Nil(t, noPrealloc)
		assert.NotNil(t, prealloc)
		assert.NotEqual(t, noPrealloc, prealloc)

		// Но семантически по контракту задачи их содержимое эквивалентно:
		assert.Empty(t, noPrealloc)
		assert.Empty(t, prealloc)
		assert.Len(t, noPrealloc, 0)
		assert.Len(t, prealloc, 0)
		assert.True(t, slices.Equal(noPrealloc, prealloc))
	})

	t.Run("при n > 0 содержимое срезов полностью совпадает", func(t *testing.T) {
		cases := []int{1, 1000}

		for _, n := range cases {
			noPrealloc := Build(n, false)
			prealloc := Build(n, true)

			assert.Len(t, noPrealloc, n)
			assert.Len(t, prealloc, n)
			assert.Equal(t, noPrealloc, prealloc)
			assert.True(t, slices.Equal(noPrealloc, prealloc))
		}
	})
}

var sliceResult []int

func BenchmarkBuild(b *testing.B) {
	sizes := []int{1, 1000}

	for _, n := range sizes {
		b.Run(fmt.Sprintf("n=%d/NoPrealloc", n), func(b *testing.B) {
			b.ReportAllocs()
			for i := 0; i < b.N; i++ {
				sliceResult = Build(n, false)
			}
		})

		b.Run(fmt.Sprintf("n=%d/Prealloc", n), func(b *testing.B) {
			b.ReportAllocs()
			for i := 0; i < b.N; i++ {
				sliceResult = Build(n, true)
			}
		})
	}
}
