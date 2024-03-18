package hw04lrucache

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCache(t *testing.T) {
	t.Run("empty cache", func(t *testing.T) {
		var c LRU
		c.New(10)
		_, ok := c.Get("aaa")
		require.False(t, ok)

		_, ok = c.Get("bbb")
		require.False(t, ok)
	})

	t.Run("simple", func(t *testing.T) {
		var c LRU
		c.New(5)

		wasInCache := c.Set("aaa", 100)
		require.False(t, wasInCache)

		wasInCache = c.Set("bbb", 200)
		require.False(t, wasInCache)

		val, ok := c.Get("aaa")
		require.True(t, ok)
		require.Equal(t, 100, val)

		val, ok = c.Get("bbb")
		require.True(t, ok)
		require.Equal(t, 200, val)

		wasInCache = c.Set("aaa", 300)
		require.True(t, wasInCache)

		val, ok = c.Get("aaa")
		require.True(t, ok)
		require.Equal(t, 300, val)

		val, ok = c.Get("ccc")
		require.False(t, ok)
		require.Nil(t, val)
	})

	t.Run("LRU Clean", func(t *testing.T) {
		var lru LRU
		lru.New(3)
		cacheValues := make([]int, 0, lru.Capacity)

		lru.Set("aaa", 10) // head -> aaa(10) -> tail
		lru.Set("bbb", 20) // head -> bbb(20) -> aaa(10) -> tail
		lru.Set("ccc", 30) // head -> ccc(30) -> bbb(20) -> aaa(10) -> tail

		lru.Clean() // head==nil -> (emty ( len==0 )) -> tail==nil
		require.Equal(t, 0, lru.Queue.Len())

		lru.Set("xxx", 10) // head -> xxx(10) -> tail
		for p := lru.Queue.Head; p != nil; p = p.Back {
			cacheValues = append(cacheValues, p.Value.(int))
		}
		require.Equal(t, []int{10}, cacheValues)
	})

	t.Run("LRU Empty", func(t *testing.T) {
		var lru LRU
		lru.New(3)
		require.Equal(t, 3, lru.Capacity)
	})

	t.Run("LRU Push OverSize", func(t *testing.T) {
		var lru LRU
		lru.New(3)

		cacheValues := make([]int, 0, lru.Capacity)

		lru.Set("aaa", 10) // head -> 10 -> tail
		lru.Set("bbb", 20) // head -> 20 -> 10 -> tail
		lru.Set("ccc", 30) // head -> 30 -> 20 -> 10 -> tail
		lru.Set("ddd", 40) // head -> 40 -> 30 -> 20  -> tail
		for p := lru.Queue.Head; p != nil; p = p.Back {
			cacheValues = append(cacheValues, p.Value.(int))
		}
		require.Equal(t, []int{40, 30, 20}, cacheValues)
	})
}
