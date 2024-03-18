package hw04lrucache

import "fmt"

type LRU struct {
	Capacity  int
	Queue     DoublyLinkedList
	HashTable map[string]interface{}
}

func (c *LRU) New(capacity int) {
	c.Capacity = capacity
	c.Queue = DoublyLinkedList{}
	c.HashTable = make(map[string]interface{}, capacity)
}

func (c *LRU) Set(key string, value interface{}) bool {
	if v, ok := c.HashTable[key]; ok {
		// fmt.Printf("(DEBUG): Key [%s] is exists and has pointer [%p]=>[%s]\n", key, v, v.Value)
		v.(*Item).Value = value
		c.Queue.MoveToFront(v.(*Item))
		return true // key exists
	}
	c.HashTable[key] = c.Queue.PushFront(value)
	if c.Queue.Count > c.Capacity {
		k := c.Queue.Back()
		c.Queue.Remove(k)
		delete(c.HashTable, key)
		// fmt.Printf("(DEBUG): BackAddr=%p\n", k)
	}
	return false // key doesn't exist

}

func (c *LRU) Get(key string) (interface{}, bool) {
	if v, ok := c.HashTable[key]; ok {
		// fmt.Printf("(DEBUG): Key [%s] is exists and has pointer [%p]=>[%s]\n", key, v, v.Value)
		return c.Queue.MoveToFront(v.(*Item)), true
	}
	return nil, false
}

func (c *LRU) PrintHashTable() {
	for k, v := range c.HashTable {
		fmt.Printf("%s => %p\n", k, v)
	}
}

func (c *LRU) Print() {
	fmt.Printf("LenOfList=%d HeadPointer=[%p] TailPointer=[%p]\n", c.Queue.Count, c.Queue.Head, c.Queue.Tail)
	for p := c.Queue.Head; p != nil; p = p.Back {
		fmt.Printf("\tItemAddr[%p] FrontAddr[%-12p] BackAddr[%-12p] ItemValueue[%v]\n", p, p.Front, p.Back, p.Value)
	}
}

func (c *LRU) Clean() {
	c.Queue.Head = nil
	c.Queue.Tail = nil
	c.Queue.Count = 0
	c.HashTable = make(map[string]interface{}, c.Capacity)
}
