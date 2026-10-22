package main

import "sync"

type Warehouse struct {
	ID int64
}

type Cache interface {
	Put(id int64, value any)
	Get(id int64) (any, bool)
}

type Storage struct {
	cache Cache
}

func NewStorage(cache Cache) *Storage {
	return &Storage{cache: cache}
}

// Баг: *warehouse сохранял значение Warehouse, а Get ожидал *Warehouse (несовпадение типов).
// Исправление: сохраняем сам указатель warehouse, согласовывая типы между Set и Get.
func (s *Storage) Set(warehouse *Warehouse) {
	if warehouse == nil {
		return
	}
	s.cache.Put(warehouse.ID, warehouse)
}

func (s *Storage) Get(id int64) *Warehouse {
	item, ok := s.cache.Get(id)
	if !ok {
		return nil
	}

	warehouse, ok := item.(*Warehouse)
	if !ok {
		return nil
	}
	return warehouse
}

// InMemoryCache — потокобезопасная in-memory реализация интерфейса Cache.
type InMemoryCache struct {
	mu    sync.RWMutex
	items map[int64]any
}

func NewInMemoryCache() *InMemoryCache {
	return &InMemoryCache{items: make(map[int64]any)}
}

func (c *InMemoryCache) Put(id int64, value any) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.items[id] = value
}

func (c *InMemoryCache) Get(id int64) (any, bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()
	val, ok := c.items[id]
	return val, ok
}
