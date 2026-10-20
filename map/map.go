package _map

// Frequencies подсчитывает частоту вхождения каждого слова в срезе words и возвращает карту частот.
// Сложность: по времени O(n), по памяти O(n)
func Frequencies(words []string) map[string]int {
	res := make(map[string]int, len(words))
	for _, word := range words {
		res[word]++
	}

	return res
}

// MostFrequent возвращает наиболее часто встречающееся слово и число его повторений, разрешая равенство частот выбором лексикографически меньшего слова.
// Сложность: по времени O(n), по памяти O(1)
func MostFrequent(freq map[string]int) (word string, count int, ok bool) {
	for key, val := range freq {
		if !ok || val > count || (val == count && key < word) {
			word = key
			count = val
			ok = true
		}
	}

	return word, count, ok
}
