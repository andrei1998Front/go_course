package main

import "fmt"

func main() {
	list := List{}

	list.PushFront(3)
	list.PushBack(2)
	list.PushBack(1)
	list.PushBack("rdg")

	itemForRemove := list.First().Next().Next()

	list.Remove(*itemForRemove)


	addrItemForCheck := list.First().Next()
	fmt.Println(list)
	fmt.Println(addrItemForCheck)

	fmt.Println(list.Len())
}
