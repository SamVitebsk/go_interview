package data_race

import (
	"sync"
	"sync/atomic"
)

// Count демонстрирует состояние гонки (data race) при конкурентном инкременте счетчика без синхронизации.
// Результат может быть меньше 100: операция counter++ неатомарна (чтение-изменение-запись), из-за чего горутины читают устаревшие значения и перезаписывают инкременты друг друга.
func Count() int {
	var counter int
	var wg sync.WaitGroup

	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			counter++
		}()
	}

	wg.Wait()
	return counter
}

// CountWithMutex устраняет data race с помощью [sync.Mutex], гарантируя монопольный доступ к счетчику.
// В production мьютекс выбирают тогда, когда критическая секция содержит сложную бизнес-логику или требует атомарного изменения нескольких связанных полей.
func CountWithMutex() int {
	var counter int
	var wg sync.WaitGroup
	var mu sync.Mutex

	for range 100 {
		wg.Go(func() {
			mu.Lock()
			counter++
			mu.Unlock()
		})
	}

	wg.Wait()
	return counter
}

// CountWithAtomic устраняет data race с помощью атомарного счетчика [atomic.Int64] без блокировок.
// Выбор для production: для простого счетчика предпочтителен именно вариант с atomic (CountWithAtomic).
// Причина: он lock-free, выполняется на уровне аппаратных CPU-инструкций, не создает оверхеда на планировщик
// и блокировку горутин, а также исключает риск взаимных блокировок (deadlock).
func CountWithAtomic() int {
	var counter atomic.Int64
	var wg sync.WaitGroup

	for range 100 {
		wg.Go(func() {
			counter.Add(1)
		})
	}

	wg.Wait()
	return int(counter.Load())
}
