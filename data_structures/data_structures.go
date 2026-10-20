package data_structures

import (
	"errors"
	"math/rand/v2"
	"slices"
)

var (
	// ErrNegativeCount возвращается, когда запрошено отрицательное количество чисел.
	ErrNegativeCount = errors.New("количество чисел (n) не может быть отрицательным")

	// ErrNegativeMax возвращается, когда верхняя граница диапазона отрицательна.
	ErrNegativeMax = errors.New("верхняя граница диапазона (max) не может быть отрицательной")

	// ErrCountExceedsMax возвращается, когда запрошенное количество n превышает размер диапазона max.
	ErrCountExceedsMax = errors.New("количество запрашиваемых чисел (n) не может превышать диапазон (max)")

	// ErrNilRandomSource возвращается, когда передан неинициализированный генератор случайных чисел.
	ErrNilRandomSource = errors.New("источник случайных чисел (rnd) не может быть nil")
)

// UniqueRandom генерирует и возвращает срез из n уникальных псевдослучайных чисел
// в полуинтервале [0, max), используя переданный генератор rnd.
// По памяти O(n), по времени в среднем O(n), может и больше из-за частых коллизий
func UniqueRandom(n int, max int, rnd *rand.Rand) ([]int, error) {
	if n < 0 {
		return nil, ErrNegativeCount
	}
	if max < 0 {
		return nil, ErrNegativeMax
	}
	if n > max {
		return nil, ErrCountExceedsMax
	}
	if rnd == nil {
		return nil, ErrNilRandomSource
	}

	uniqueNums := make(map[int]struct{}, n)
	res := make([]int, 0, n)

	for len(uniqueNums) < n {
		num := rnd.IntN(max)
		if _, exists := uniqueNums[num]; !exists {
			uniqueNums[num] = struct{}{}
			res = append(res, num)
		}
	}

	return res, nil
}

// UniqueRandomWithoutMap генерирует и возвращает срез из n уникальных псевдослучайных чисел
// в полуинтервале [0, max), используя переданный генератор rnd.
// По памяти O(n) (доп. память O(1)), по времени в среднем O(n²), может и больше из-за частых коллизий
func UniqueRandomWithoutMap(n int, max int, rnd *rand.Rand) ([]int, error) {
	if n < 0 {
		return nil, ErrNegativeCount
	}
	if max < 0 {
		return nil, ErrNegativeMax
	}
	if n > max {
		return nil, ErrCountExceedsMax
	}
	if rnd == nil {
		return nil, ErrNilRandomSource
	}

	res := make([]int, 0, n)

	for len(res) < n {
		num := rnd.IntN(max)
		if !slices.Contains(res, num) {
			res = append(res, num)
		}
	}

	return res, nil
}
