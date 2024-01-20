package main

type Item struct {
	value interface{}
	next  *Item
	prev  *Item
}

func (item Item) Value() interface{} {
	return item.value
}

func (item Item) Next() *Item {
	return item.next
}

func (item Item) Prev() *Item {
	return item.prev
}
