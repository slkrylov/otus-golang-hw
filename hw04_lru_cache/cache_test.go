package hw04lrucache

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCache(t *testing.T) {
	t.Run("empty cache", func(t *testing.T) {
		c := NewCache(10)

		_, ok := c.Get("aaa")
		require.False(t, ok)

		_, ok = c.Get("bbb")
		require.False(t, ok)
	})

	t.Run("simple", func(t *testing.T) {
		c := NewCache(5)
		wasInCache := c.Set("aaa", 100)
		require.False(t, wasInCache)

		wasInCache = c.Set("bbb", 200)
		require.False(t, wasInCache)

		val, ok := c.Get("aaa")
		require.True(t, ok)
		require.Equal(t, 100, val.(lruValue).LruVal)

		val, ok = c.Get("bbb")
		require.True(t, ok)
		require.Equal(t, 200, val.(lruValue).LruVal)

		wasInCache = c.Set("aaa", 300)
		require.True(t, wasInCache)

		val, ok = c.Get("aaa")
		require.True(t, ok)
		require.Equal(t, 300, val.(lruValue).LruVal)

		val, ok = c.Get("ccc")
		require.False(t, ok)
		require.Nil(t, val)
	})

	t.Run("LRU Push OverSize", func(t *testing.T) {
		lru := NewCache(3)
		lru.Set("aaa", 10) // head -> 10 -> tail
		lru.Set("bbb", 20) // head -> 20 -> 10 -> tail
		lru.Set("ccc", 30) // head -> 30 -> 20 -> 10 -> tail
		lru.Set("ddd", 40) // head -> 40 -> 30 -> 20  -> tail

		_, ok := lru.Get("aaa")
		require.False(t, ok)
	})
}
