package wait_group_mutex_once_pool

import (
	"sync"
	"sync/atomic"
)

// Counter реализует потокобезопасный счетчик на основе атомарных операций atomic.Int64
//
// Сравнение (Mutex vs Atomic):
//   - Atomic (Count/AtomicCounter) лучше использовать для простых числовых счетчиков, флагов и метрик.
//     Он lock-free, выполняется на уровне аппаратных CPU-инструкций, не создает оверхеда на планировщик
//     и блокировку горутин, а также исключает риск взаимных блокировок (deadlock).
//   - Mutex (MutexCounter) необходим, когда критическая секция содержит сложную бизнес-логику,
//     требует проверки условий перед записью, одновременного обновления нескольких
//     связанных полей или защиты неатомарных структур данных (например, map или slice).
type Counter struct {
	val atomic.Int64
}

// Inc атомарно увеличивает значение счетчика на 1
func (c *Counter) Inc() {
	c.val.Add(1)
}

// Add атомарно увеличивает значение счетчика на n
func (c *Counter) Add(n int) {
	c.val.Add(int64(n))
}

// Value атомарно возвращает текущее значение счетчика
func (c *Counter) Value() int {
	return int(c.val.Load())
}

// MutexCounter реализует потокобезопасный счетчик с защитой критической секции через sync.Mutex
type MutexCounter struct {
	mu  sync.Mutex
	val int
}

// Inc увеличивает значение счетчика на 1 под монопольной блокировкой мьютекса
func (c *MutexCounter) Inc() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.val++
}

// Add увеличивает значение счетчика на n под монопольной блокировкой мьютекса
func (c *MutexCounter) Add(n int) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.val += n
}

// Value возвращает текущее значение счетчика под блокировкой мьютекса
func (c *MutexCounter) Value() int {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.val
}
