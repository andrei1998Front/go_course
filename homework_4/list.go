package main

type List struct {
	len   int
	first *Item
	last  *Item
}

func (list List) Len() int {
	return list.len
}

func (list List) First() *Item {
	return list.first
}

func (list List) Last() *Item {
	return list.last
}

func (list *List) PushFront(v interface{}) {
	newItem := &Item{value: v}
	if list.first == nil {
		list.first = newItem
		list.last = newItem
		list.len = 1
		return
	}

	newItem.next = list.first
	list.first.prev = newItem
	list.first = newItem
	list.len++
}

func (list *List) PushBack(v interface{}) {
	newItem := &Item{value: v}
	if list.last == nil {
		list.first = newItem
		list.last = newItem
		list.len = 1
		return
	}

	newItem.prev = list.last
	list.last.next = newItem
	list.last = newItem
	list.len++
}

func (list *List) Remove(item Item) {
	if item.prev == nil {
		list.first = item.next
	} else {
		item.prev.next = item.next
	}

	if item.next == nil {
		list.last = item.prev
	} else {
		item.next.prev = item.prev
	}

	list.len--
}
