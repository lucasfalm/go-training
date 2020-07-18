package main

import "fmt"

func main() {
	linkedSet := LinkedSet{}

	linkedSet.addLast(5)
	linkedSet.addLast(6)
	linkedSet.addLast(7)

	fmt.Println(linkedSet)
}

type Node struct {
	value int
	next  *Node
}

type LinkedSet struct {
	root *Node
	size int
}

func (l *LinkedSet) addLast(value int) {
	node := &Node{
		value: value,
	}

	if l.size == 0 {
		l.root = node
		l.size += 1
	} else {
		currentNode := l.root
		for currentNode.next != nil {
			currentNode = currentNode.next
		}

		currentNode.next = node
	}

	l.size++
}
