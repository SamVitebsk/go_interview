package generics

type Set[T comparable] map[T]struct{}

func (s Set[T]) Add(val T) {
	s[val] = struct{}{}
}

func (s Set[T]) Contains(val T) bool {
	_, exists := s[val]
	return exists
}

func (s Set[T]) Delete(val T) {
	delete(s, val)
}

func (s Set[T]) Len() int {
	return len(s)
}

// Unique возвращает срез уникальных элементов с сохранением порядка их первого появления.
func Unique[T comparable](items []T) []T {
	if len(items) == 0 {
		return []T{}
	}

	seen := make(Set[T], len(items))
	res := make([]T, 0, len(items))

	for _, val := range items {
		if !seen.Contains(val) {
			seen.Add(val)
			res = append(res, val)
		}
	}

	return res
}

// Некомпилируемый пример использования slice в качестве ключа Set:
//
//var _ = Set[[]int]{}
//
//type Report struct {
//    Values []int
//}
//var _ = Set[Report]{}
//
// Ошибка компилятора:
//   "[]int does not satisfy comparable"
//   "Report does not satisfy comparable"
//
// Объяснение ошибки:
//   Срезы (slice), мапы (map) и функции (func) в Go являются несравнимыми типами (non-comparable).
//   Срез ссылается на мутабельный базовый массив: изменение элементов изменило бы хэш ключа в map,
//   сломав инварианты хэш-таблицы (значение стало бы невозможно найти). Кроме того, поэлементное
//   сравнение срезов требует O(N) времени вместо константного O(1), обязательного для операций с map.
