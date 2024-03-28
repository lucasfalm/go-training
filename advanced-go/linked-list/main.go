package main

import (
	"fmt"

	linked_list_core "github.com/lucasfalm/go-training/advanced-go/linked-list/linked-list-core"
)

func main() {
	linkedList := linked_list_core.LinkedList{}

	fmt.Println("adding 32")
	linkedList.Push(32)
	linkedList.Print()

	fmt.Println("adding 10")
	linkedList.Push(10)
	linkedList.Print()

	fmt.Println("adding 5")
	linkedList.Push(5)
	linkedList.Print()

	fmt.Println("popping")
	linkedList.Pop()
	linkedList.Print()

	fmt.Println("popping")
	linkedList.Pop()
	linkedList.Print()

	fmt.Println("popping")
	linkedList.Pop()
	linkedList.Print()

	fmt.Println("populating linked list...")
	linkedList.Push(1)
	linkedList.Push(290)
	linkedList.Push(948)
	linkedList.Push(1920)
	linkedList.Push(100000)
	linkedList.Push(2)
	linkedList.Print()
}
