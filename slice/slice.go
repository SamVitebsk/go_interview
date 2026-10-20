package slice

import (
	"slices"
)

// CopyWithout возвращает новый срез без target: время O(n), память O(n).
func CopyWithout[T comparable](s []T, target T) []T {
	if s == nil {
		return nil
	}

	res := make([]T, 0, len(s))

	for _, val := range s {
		if val != target {
			res = append(res, val)
		}
	}

	if cap(res) > 1024 && len(res) < cap(res)/2 {
		return slices.Clone(res)
	}

	return res
}

// DeleteInPlace удаляет target на месте с занулением хвоста: время O(n), память O(1).
func DeleteInPlace[T comparable](s []T, target T) []T {
	if s == nil {
		return nil
	}

	i := 0
	for i < len(s) && s[i] != target {
		i++
	}
	if i == len(s) {
		return s
	}

	idx := i
	for _, val := range s[i+1:] {
		if target != val {
			s[idx] = val
			idx++
		}
	}

	clear(s[idx:])

	return s[:idx]
}
