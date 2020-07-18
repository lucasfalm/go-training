package main

import "fmt"

func main() {
	linkedSet := LinkedSet{}

	linkedSet.addLast(5)
	linkedSet.addLast(6)
	linkedSet.removeLast()

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

func (l *LinkedSet) removeLast() {
	previusNode := &Node{}
	currentNode := l.root

	for l.root.next != nil {
		previusNode = currentNode

		currentNode = currentNode.next
	}
	fmt.Println(currentNode)
	previusNode.next = nil
	l.size--

	fmt.Println(currentNode)
}
