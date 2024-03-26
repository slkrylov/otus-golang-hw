package hw04lrucache

import "fmt"

type Key string

type Cache interface {
	Clear()
	Len() int
	Print()
	PrintHashTable()
	Queue() *ListItem
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

func (c *lruCache) Len() int {
	return c.queue.Len()
}

func (c *lruCache) Queue() *ListItem {
	return c.queue.Front()
}

func (c *lruCache) Set(key Key, value interface{}) bool {
	lv := lruValue{LruKey: key, LruVal: value}
	if v, ok := c.items[key]; ok {
		// fmt.Printf("(DEBUG): Key [%s] is exists and has pointer [%p] to value [%v]\n", key, v, v.Value)
		v.Value = lv
		c.queue.MoveToFront(v)
		return true // key exists
	}
	if c.queue.Len()+1 > c.capacity {
		k := c.queue.Back() // Get l.Tail pointer
		c.queue.Remove(k)
		c.items[key] = c.queue.PushFront(lv)
		delete(c.items, k.Value.(lruValue).LruKey)
		// fmt.Printf("(DEBUG): Remove BackAddr=%p BackLRUKey=%v\n", k, k.Value.(lruValue).LruKey)
		return false // key doesn't exist
	}
	c.items[key] = c.queue.PushFront(lv)
	return false // key doesn't exist
}

func (c *lruCache) Get(key Key) (interface{}, bool) {
	if v, ok := c.items[key]; ok {
		// fmt.Printf("(DEBUG): Key [%s] is exists and has pointer [%p]=>[%v]\n", key, v, v.Value)
		return c.queue.MoveToFront(v), true
	}
	return nil, false
}

func (c *lruCache) PrintHashTable() {
	for k, v := range c.items {
		fmt.Printf("%s => %p [%v]\n", k, v, v.Value)
	}
}

func (c *lruCache) Print() {
	if c.queue.Len() > 0 {
		fmt.Printf("LenOfList=%d HeadPointer=[%p] TailPointer=[%p]\n", c.queue.Len(), c.queue.Front(), c.queue.Back())
		for p := c.queue.Front(); p != nil; p = p.Back {
			fmt.Printf("\tItemAddr[%p] FrontAddr[%-12p] BackAddr[%-12p] ItemValueue[%v]\n", p, p.Front, p.Back, p.Value)
		}
	}
}

func (c *lruCache) Clear() {
	c.queue.Clean()
	c.items = make(map[Key]*ListItem, c.capacity)
}
