package data_structures

import (
	"math/rand/v2"
	"slices"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestUniqueRandom(t *testing.T) {
	t.Run("неверные аргументы", func(t *testing.T) {
		t.Parallel()

		dummyRnd := rand.New(rand.NewPCG(1, 2))

		tests := []struct {
			name    string
			n       int
			maxVal  int
			rnd     *rand.Rand
			wantErr error
		}{
			{
				name:    "отрицательное n (n < 0)",
				n:       -1,
				maxVal:  10,
				rnd:     dummyRnd,
				wantErr: ErrNegativeCount,
			},
			{
				name:    "отрицательный max (max < 0)",
				n:       5,
				maxVal:  -1,
				rnd:     dummyRnd,
				wantErr: ErrNegativeMax,
			},
			{
				name:    "n превышает max (n > max)",
				n:       11,
				maxVal:  10,
				rnd:     dummyRnd,
				wantErr: ErrCountExceedsMax,
			},
			{
				name:    "источник случайных чисел nil (rnd == nil)",
				n:       5,
				maxVal:  10,
				rnd:     nil,
				wantErr: ErrNilRandomSource,
			},
		}

		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				got, err := UniqueRandom(tt.n, tt.maxVal, tt.rnd)
				assert.ErrorIs(t, err, tt.wantErr)
				assert.Nil(t, got)
			})
		}
	})

	t.Run("пустой результат при n = 0", func(t *testing.T) {
		t.Parallel()

		rnd := rand.New(rand.NewPCG(10, 20))

		got, err := UniqueRandom(0, 10, rnd)
		require.NoError(t, err)
		assert.Empty(t, got)

		gotZeroMax, err := UniqueRandom(0, 0, rnd)
		require.NoError(t, err)
		assert.Empty(t, gotZeroMax)
	})

	t.Run("обычный случай: корректная длина, диапазон и отсутствие повторов", func(t *testing.T) {
		t.Parallel()

		rnd := rand.New(rand.NewPCG(42, 99))
		n := 25
		maxVal := 100

		got, err := UniqueRandom(n, maxVal, rnd)
		require.NoError(t, err)
		require.Len(t, got, n)

		seen := make(map[int]struct{}, n)
		for i, val := range got {
			assert.GreaterOrEqual(t, val, 0, "элемент got[%d] должен быть >= 0", i)
			assert.Less(t, val, maxVal, "элемент got[%d] должен быть < maxVal", i)
			assert.NotContains(t, seen, val, "обнаружен повтор числа %d на позиции %d", val, i)
			seen[val] = struct{}{}
		}
	})

	t.Run("полный диапазон при n = max", func(t *testing.T) {
		t.Parallel()

		rnd := rand.New(rand.NewPCG(77, 88))
		n := 30
		maxVal := 30

		got, err := UniqueRandom(n, maxVal, rnd)
		require.NoError(t, err)
		require.Len(t, got, n)

		sorted := slices.Clone(got)
		slices.Sort(sorted)

		want := make([]int, maxVal)
		for i := range maxVal {
			want[i] = i
		}

		assert.Equal(t, want, sorted)
	})

	t.Run("воспроизводимость результатов при одинаковом seed", func(t *testing.T) {
		t.Parallel()

		rnd1 := rand.New(rand.NewPCG(12345, 67890))
		rnd2 := rand.New(rand.NewPCG(12345, 67890))

		n := 15
		maxVal := 50

		got1, err1 := UniqueRandom(n, maxVal, rnd1)
		require.NoError(t, err1)

		got2, err2 := UniqueRandom(n, maxVal, rnd2)
		require.NoError(t, err2)

		assert.Equal(t, got1, got2)
	})
}

func TestUniqueRandomWithoutMap(t *testing.T) {
	t.Run("неверные аргументы", func(t *testing.T) {
		t.Parallel()

		dummyRnd := rand.New(rand.NewPCG(1, 2))

		got, err := UniqueRandomWithoutMap(-1, 10, dummyRnd)
		assert.ErrorIs(t, err, ErrNegativeCount)
		assert.Nil(t, got)

		got, err = UniqueRandomWithoutMap(5, -1, dummyRnd)
		assert.ErrorIs(t, err, ErrNegativeMax)
		assert.Nil(t, got)

		got, err = UniqueRandomWithoutMap(11, 10, dummyRnd)
		assert.ErrorIs(t, err, ErrCountExceedsMax)
		assert.Nil(t, got)

		got, err = UniqueRandomWithoutMap(5, 10, nil)
		assert.ErrorIs(t, err, ErrNilRandomSource)
		assert.Nil(t, got)
	})

	t.Run("совпадение результатов с версией с мапой при одинаковом seed", func(t *testing.T) {
		t.Parallel()

		rnd1 := rand.New(rand.NewPCG(42, 99))
		rnd2 := rand.New(rand.NewPCG(42, 99))
		n := 20
		maxVal := 50

		gotWithMap, err1 := UniqueRandom(n, maxVal, rnd1)
		require.NoError(t, err1)

		gotWithoutMap, err2 := UniqueRandomWithoutMap(n, maxVal, rnd2)
		require.NoError(t, err2)

		assert.Equal(t, gotWithMap, gotWithoutMap)
	})
}
