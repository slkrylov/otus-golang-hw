package hw04lrucache

import "fmt"

type DoublyLinkedList struct {
	Head  *Item
	Tail  *Item
	Count int
}

type Item struct {
	Front *Item
	Back  *Item
	Value interface{}
}

func (l *DoublyLinkedList) PushBack(v interface{}) *Item {
	i := Item{Value: v}
	l.Count++
	switch {
	case l.Tail == nil && l.Head == nil:
		l.Tail = &i
		l.Head = &i
	case l.Tail != nil:
		i.Front = l.Tail
		(l.Tail).Back = &i
		l.Tail = &i
	}
	return &i
}

func (l *DoublyLinkedList) PushFront(v interface{}) *Item {
	i := Item{Value: v}
	l.Count++
	switch {
	case l.Tail == nil && l.Head == nil:
		l.Tail = &i
		l.Head = &i
	case l.Head != nil:
		i.Back = l.Head
		(l.Head).Front = &i
		l.Head = &i
	}
	return &i
}

func (l *DoublyLinkedList) Front() *Item {
	return l.Head
}

func (l *DoublyLinkedList) Back() *Item {
	return l.Tail
}

func (i *Item) Next() *Item {
	return i.Back
}

func (i *Item) Prev() *Item {
	return i.Front
}

func (i *Item) ItemValue() interface{} {
	return i.Value
}

func (l *DoublyLinkedList) Len() int {
	return l.Count
}

func (l *DoublyLinkedList) Remove(i *Item) int {
	switch {
	case i.Front == nil:
		l.Head = i.Back
		(i.Back).Front = nil
	case i.Front != nil && i.Back != nil:
		(i.Front).Back = i.Back
		(i.Back).Front = i.Front
	case i.Back == nil:
		l.Tail = i.Front
		(i.Front).Back = nil
	}
	l.Count--
	return l.Count
}

func (l *DoublyLinkedList) FindFirstValue(v int) *Item {
	var r *Item
	for p := l.Head; p != nil; p = p.Back {
		r = p
		if p.Value == v {
			return r
		}
	}
	return nil
}

func (l *DoublyLinkedList) MoveToFront(i *Item) interface{} {
	l.Remove(i)
	i.Front = nil
	i.Back = l.Head
	l.Head.Front = i
	l.Head = i
	l.Count++
	return i.Value
}

func (l *DoublyLinkedList) Print() {
	fmt.Printf("LenOfList=%d HeadPointer=[%p] TailPointer=[%p]\n", l.Count, l.Head, l.Tail)
	for p := l.Head; p != nil; p = p.Back {
		fmt.Printf("\tItemAddr[%p] FrontAddr[%-12p] BackAddr[%-12p] ItemValueue[%s]\n", p, p.Front, p.Back, p.Value)
	}
}
