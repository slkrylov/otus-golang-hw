package hw04lrucache

type Key string

type Cache interface {
	Clear()
	Set(key Key, value interface{}) bool
	Get(key Key) (interface{}, bool)
}

type lruCache struct {
	capacity int
	queue    List
	items    map[Key]*ListItem
}

type lruValue struct {
	LruKey Key
	LruVal interface{}
}

func NewCache(capacity int) Cache {
	return &lruCache{
		capacity: capacity,
		queue:    NewList(),
		items:    make(map[Key]*ListItem, capacity),
	}
}

func (c *lruCache) Set(key Key, value interface{}) bool {
	lv := lruValue{LruKey: key, LruVal: value}
	if v, ok := c.items[key]; ok {
		v.Value = lv
		c.queue.MoveToFront(v)
		return true // key exists
	}
	if c.queue.Len()+1 > c.capacity {
		k := c.queue.Back() // Get l.Tail pointer
		c.queue.Remove(k)
		c.items[key] = c.queue.PushFront(lv)
		delete(c.items, k.Value.(lruValue).LruKey)
		return false // key doesn't exist
	}
	c.items[key] = c.queue.PushFront(lv)
	return false // key doesn't exist
}

func (c *lruCache) Get(key Key) (interface{}, bool) {
	if v, ok := c.items[key]; ok {
		return c.queue.MoveToFront(v), true
	}
	return nil, false
}

func (c *lruCache) Clear() {
	c.queue.Clean()
	c.items = make(map[Key]*ListItem, c.capacity)
}
