package hw04lrucache

type List interface {
	Len() int
	Front() *ListItem
	Back() *ListItem
	PushFront(v interface{}) *ListItem
	PushBack(v interface{}) *ListItem
	Remove(i *ListItem) int
	MoveToFront(i *ListItem) interface{}
	Clean()
}

type list struct {
	Head  *ListItem
	Tail  *ListItem
	Count int
}

type ListItem struct {
	Front *ListItem
	Back  *ListItem
	Value interface{}
}

func NewList() List {
	return new(list)
}

func (l *list) PushBack(v interface{}) *ListItem {
	i := ListItem{Value: v}
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

func (l *list) PushFront(v interface{}) *ListItem {
	i := ListItem{Value: v}
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

func (l *list) Front() *ListItem {
	return l.Head
}

func (l *list) Back() *ListItem {
	return l.Tail
}

func (i *ListItem) Next() *ListItem {
	return i.Back
}

func (i *ListItem) Prev() *ListItem {
	return i.Front
}

func (i *ListItem) ListItemValue() interface{} {
	return i.Value
}

func (l *list) Len() int {
	return l.Count
}

func (l *list) Remove(i *ListItem) int {
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

func (l *list) FindFirstValue(v int) *ListItem {
	var r *ListItem
	for p := l.Head; p != nil; p = p.Back {
		r = p
		if p.Value == v {
			return r
		}
	}
	return nil
}

func (l *list) MoveToFront(i *ListItem) interface{} {
	l.Remove(i)
	i.Front = nil
	i.Back = l.Head
	l.Head.Front = i
	l.Head = i
	l.Count++
	return i.Value
}

func (l *list) Clean() {
	l.Count = 0
	l.Head = nil
	l.Tail = nil
}
