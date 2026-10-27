package scope_closure_range

// Accumulator возвращает замыкание, накапливающее сумму переданных целых чисел, начиная с 0.
func Accumulator() func(int) int {
	sum := 0

	return func(n int) int {
		sum += n
		return sum
	}
}

// Readers возвращает слайс замыканий, возвращающих фиксированные значения элементов входного среза.
func Readers(values []int) []func() int {
	res := make([]func() int, 0, len(values))
	for _, val := range values {
		res = append(res, func() int {
			return val
		})
	}

	return res
}
