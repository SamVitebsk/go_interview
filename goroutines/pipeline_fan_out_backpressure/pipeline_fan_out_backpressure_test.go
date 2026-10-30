package pipeline_fan_out_backpressure

import (
	"context"
	"sort"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestPipeline_Complete(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	ids := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	ints := gen(ctx, ids)
	users := load(ctx, ints)
	filtered := filter(ctx, users)
	res := collect(ctx, filtered)

	assert.Len(t, res, 5)

	var collectedIDs []int
	for _, u := range res {
		assert.True(t, u.Active)
		collectedIDs = append(collectedIDs, u.ID)
	}

	sort.Ints(collectedIDs)
	assert.Equal(t, []int{2, 4, 6, 8, 10}, collectedIDs)
}

func TestPipeline_Cancellation(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Millisecond)
	defer cancel()

	var ids []int
	for i := 1; i <= 1000; i++ {
		ids = append(ids, i)
	}

	ints := gen(ctx, ids)
	users := load(ctx, ints)
	filtered := filter(ctx, users)
	res := collect(ctx, filtered)

	assert.True(t, len(res) < 1000)
}
