package data_structures

import (
	"errors"
	"math/rand/v2"
	"slices"
	"testing"
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
				if !errors.Is(err, tt.wantErr) {
					t.Errorf("UniqueRandom(%d, %d, rnd) error = %v, want %v", tt.n, tt.maxVal, err, tt.wantErr)
				}
				if got != nil {
					t.Errorf("UniqueRandom(%d, %d, rnd) got = %v, want nil", tt.n, tt.maxVal, got)
				}
			})
		}
	})

	t.Run("пустой результат при n = 0", func(t *testing.T) {
		t.Parallel()

		rnd := rand.New(rand.NewPCG(10, 20))

		got, err := UniqueRandom(0, 10, rnd)
		if err != nil {
			t.Fatalf("UniqueRandom(0, 10, rnd) неожиданная ошибка: %v", err)
		}
		if len(got) != 0 {
			t.Errorf("len(got) = %d, want 0", len(got))
		}

		gotZeroMax, err := UniqueRandom(0, 0, rnd)
		if err != nil {
			t.Fatalf("UniqueRandom(0, 0, rnd) неожиданная ошибка: %v", err)
		}
		if len(gotZeroMax) != 0 {
			t.Errorf("len(gotZeroMax) = %d, want 0", len(gotZeroMax))
		}
	})

	t.Run("обычный случай: корректная длина, диапазон и отсутствие повторов", func(t *testing.T) {
		t.Parallel()

		rnd := rand.New(rand.NewPCG(42, 99))
		n := 25
		maxVal := 100

		got, err := UniqueRandom(n, maxVal, rnd)
		if err != nil {
			t.Fatalf("UniqueRandom(%d, %d, rnd) неожиданная ошибка: %v", n, maxVal, err)
		}

		if len(got) != n {
			t.Errorf("длина среза = %d, want %d", len(got), n)
		}

		seen := make(map[int]struct{}, n)
		for i, val := range got {
			if val < 0 || val >= maxVal {
				t.Errorf("элемент got[%d] = %d вне диапазона [0, %d)", i, val, maxVal)
			}
			if _, exists := seen[val]; exists {
				t.Errorf("обнаружен повтор числа %d на позиции %d", val, i)
			}
			seen[val] = struct{}{}
		}
	})

	t.Run("полный диапазон при n = max", func(t *testing.T) {
		t.Parallel()

		rnd := rand.New(rand.NewPCG(77, 88))
		n := 30
		maxVal := 30

		got, err := UniqueRandom(n, maxVal, rnd)
		if err != nil {
			t.Fatalf("UniqueRandom(%d, %d, rnd) неожиданная ошибка: %v", n, maxVal, err)
		}

		if len(got) != n {
			t.Fatalf("длина среза = %d, want %d", len(got), n)
		}

		sorted := slices.Clone(got)
		slices.Sort(sorted)

		want := make([]int, maxVal)
		for i := range maxVal {
			want[i] = i
		}

		if !slices.Equal(sorted, want) {
			t.Errorf("отсортированный срез = %v, want %v", sorted, want)
		}
	})

	t.Run("воспроизводимость результатов при одинаковом seed", func(t *testing.T) {
		t.Parallel()

		rnd1 := rand.New(rand.NewPCG(12345, 67890))
		rnd2 := rand.New(rand.NewPCG(12345, 67890))

		n := 15
		maxVal := 50

		got1, err1 := UniqueRandom(n, maxVal, rnd1)
		if err1 != nil {
			t.Fatalf("первый вызов вернул ошибку: %v", err1)
		}

		got2, err2 := UniqueRandom(n, maxVal, rnd2)
		if err2 != nil {
			t.Fatalf("второй вызов вернул ошибку: %v", err2)
		}

		if !slices.Equal(got1, got2) {
			t.Errorf("результаты не совпали для одинакового seed:\ngot1 = %v\ngot2 = %v", got1, got2)
		}
	})
}
