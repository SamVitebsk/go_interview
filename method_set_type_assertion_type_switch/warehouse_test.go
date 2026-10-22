package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestStorage(t *testing.T) {
	t.Run("успешная запись и чтение склада по ID", func(t *testing.T) {
		cache := NewInMemoryCache()
		storage := NewStorage(cache)

		want := &Warehouse{ID: 42}
		storage.Set(want)

		got := storage.Get(42)

		require.NotNil(t, got)
		assert.Equal(t, want.ID, got.ID)
		assert.Same(t, want, got)
	})

	t.Run("чтение отсутствующего ключа возвращает nil", func(t *testing.T) {
		cache := NewInMemoryCache()
		storage := NewStorage(cache)

		got := storage.Get(999)

		assert.Nil(t, got)
	})

	t.Run("попытка записи nil склада не паникует и не сохраняется", func(t *testing.T) {
		cache := NewInMemoryCache()
		storage := NewStorage(cache)

		storage.Set(nil)

		assert.Nil(t, storage.Get(0))
	})

	t.Run("безопасность comma-ok при несовпадении типа в кеше", func(t *testing.T) {
		cache := NewInMemoryCache()
		storage := NewStorage(cache)

		cache.Put(100, "not a warehouse pointer")

		got := storage.Get(100)

		assert.Nil(t, got)
	})
}
