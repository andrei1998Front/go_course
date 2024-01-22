package main

import (
	"fmt"
	"log"
)

func main() {
	list := List{}

	list.PushFront(3)
	list.PushBack(2)
	list.PushBack(1)
	list.PushBack("rdg")

	itemForRemove := list.First().Next().Next()

	err := list.Remove(itemForRemove)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(list.CheckItem(itemForRemove))
}
