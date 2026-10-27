package escape_analysis_gc_memory

// Build собирает срез чисел от 0 до n-1 с опциональным предварительным выделением памяти.
// Предусловие: n >= 0.
func Build(n int, reserve bool) []int {
	var res []int
	if reserve {
		res = make([]int, 0, n)
	}

	for i := range n {
		res = append(res, i)
	}

	return res
}
