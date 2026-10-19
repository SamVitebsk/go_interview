package main

import "fmt"

// Bag представляет собой набор уникальных строк (множество/Set) с полезным нулевым значением.
type Bag struct {
	items map[string]struct{}
}

// Add добавляет строку в набор. При первом вызове лениво инициализирует внутреннюю мапу.
func (b *Bag) Add(value string) {
	if b.items == nil {
		b.items = make(map[string]struct{})
	}
	b.items[value] = struct{}{}
}

// Contains проверяет присутствие строки в наборе. Безопасен для неинициализированного Bag.
func (b *Bag) Contains(value string) bool {
	_, ok := b.items[value]
	return ok
}

// Len возвращает количество уникальных строк в наборе.
func (b *Bag) Len() int {
	return len(b.items)
}

func main() {
	fmt.Println("\n--- 2. Сравнение map ---")

	// 1.1 var m map[string]int
	// Тип: map[string]int, Значение: nil
	var m1 map[string]int
	fmt.Printf("m1: тип=%T, nil=%t, len=%d, чтение m1[\"k\"]=%d\n", m1, m1 == nil, len(m1), m1["k"])
	// m1["k"] = 10 // panic при записи в nil мапу

	// 1.2 new(map[string]int)
	// Тип: *map[string]int, Указатель не nil, но указывает на nil мапу
	m2 := new(map[string]int)
	fmt.Printf("m2: тип=%T, p_nil=%t, *p_nil=%t, len=%d\n", m2, m2 == nil, *m2 == nil, len(*m2))
	// (*m2)["k"] = 10 // panic при записи в nil мапу

	// 1.3 make(map[string]int)
	// Тип: map[string]int, Значение: инициализированная пустая мапа, не nil
	m3 := make(map[string]int)
	m3["k"] = 10 // Безопасно: make инициализирует внутренние структуры хэш-таблицы
	fmt.Printf("m3: тип=%T, nil=%t, len=%d, m3[\"k\"]=%d\n", m3, m3 == nil, len(m3), m3["k"])

	// 1.4 map[string]int{}
	// Тип: map[string]int, Значение: пустая мапа через литерал, не nil
	m4 := map[string]int{}
	m4["k"] = 20 // Безопасно: аналогично make, мапа создана и готова к записи
	fmt.Printf("m4: тип=%T, nil=%t, len=%d, m4[\"k\"]=%d\n", m4, m4 == nil, len(m4), m4["k"])

	fmt.Println("\n--- 2. Сравнение slice ---")

	// 2.1 var s []int
	// Тип: []int, Значение: nil slice (len=0, cap=0, s1 == nil -> true)
	var s1 []int
	fmt.Printf("s1: тип=%T, nil=%t, len=%d, cap=%d\n", s1, s1 == nil, len(s1), cap(s1))
	s1 = append(s1, 10) // Безопасно: append сам выделяет базовый массив для nil slice
	// fmt.Println(s1[1]) // ПАНИКА: panic: runtime error: index out of range [1] with length 1

	// 2.2 new([]int)
	// Тип: *[]int, Указатель не nil, но разыменованный слайс — nil slice (*s2 == nil -> true)
	s2 := new([]int)
	fmt.Printf("s2: тип=%T, p_nil=%t, *p_nil=%t, len=%d, cap=%d\n", s2, s2 == nil, *s2 == nil, len(*s2), cap(*s2))
	*s2 = append(*s2, 20) // Безопасно: append над разыменованным слайсом создает базовый массив
	// fmt.Println((*s2)[1]) // ПАНИКА: panic: runtime error: index out of range [1] with length 1

	// 2.3 make([]int, 0, 2)
	// Тип: []int, Значение: не nil слайс с длиной 0 и емкостью 2 (s3 == nil -> false)
	s3 := make([]int, 0, 2)
	fmt.Printf("s3: тип=%T, nil=%t, len=%d, cap=%d\n", s3, s3 == nil, len(s3), cap(s3))
	s3 = append(s3, 30) // Безопасно: добавляется в уже созданный базовый массив
	fmt.Println(s3[0])  // До append s3[0] вызывал бы панику, так как len=0. Индекс 0 доступен только при len > 0.

	fmt.Println("\n--- 3. Демонстрация Bag ---")
	var b Bag
	fmt.Println("До записи:", b.Len(), b.Contains("Go")) // 0 false

	b.Add("Go")
	b.Add("Go")                                             // Повтор не увеличивает длину
	b.Add("")                                               // Пустая строка допустима
	fmt.Println("Len():", b.Len())                          // 2
	fmt.Println("Contains(\"Go\"): ", b.Contains("Go"))     // true
	fmt.Println("Contains(\"\"):   ", b.Contains(""))       // true
	fmt.Println("Contains(\"Rust\"): ", b.Contains("Rust")) // false
}
