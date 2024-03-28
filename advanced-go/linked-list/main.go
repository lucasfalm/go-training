package main

import (
	"fmt"

	linkedlist "github.com/lucasfalm/go-training/advanced-go/linked-list/linkedlist"
)

func main() {
	linkedList := linkedlist.LinkedList{}

	fmt.Println("pushing 32")
	linkedList.Push(32)
	linkedList.Print()

	fmt.Println("pushing 10")
	linkedList.Push(10)
	linkedList.Print()

	fmt.Println("pushing 5")
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

	fmt.Println("unshifting 88...")
	linkedList.Unshift(88)
	linkedList.Print()

	fmt.Println("searching value 1920...")
	linkedList.FindWithPrint(1920)

	fmt.Println("searching value 5...")
	linkedList.FindWithPrint(5)
}
