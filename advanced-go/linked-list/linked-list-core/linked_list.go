package linked_list_core

import "fmt"

type LinkedList struct {
	Nodes []*Node
	Count int
}

type Node struct {
	Head  *Node
	Tail  *Node
	Value any
}

func (l *LinkedList) Last() *Node {
	if l.Any() {
		return l.Nodes[len(l.Nodes)-1]
	}

	return nil
}

func (l *LinkedList) First() *Node {
	if l.Any() {
		return l.Nodes[0]
	}

	return nil
}

func (l *LinkedList) Any() bool {
	return l.Count > 0
}

// NOTE: insert as the last node
func (l *LinkedList) Push(value any) *Node {
	newNode := Node{Value: value}

	if l.Any() {
		newNode.Head = l.Last()

		l.Last().Tail = &newNode
	}

	l.Nodes = append(l.Nodes, &newNode)

	l.Count++

	return &newNode
}

// NOTE: insert as the first node
func (l *LinkedList) Unshift(value any) *Node {
	newNode := &Node{Value: value}

	if l.Any() {
		previousFirst := l.First()

		previousFirst.Head = newNode
		newNode.Tail = previousFirst
	}

	l.Nodes = append([]*Node{newNode}, l.Nodes...)

	l.Count++

	return newNode
}

// NOTE: remove the first node
func (l *LinkedList) Shift(value any) *Node {
	return &Node{}
}

// NOTE: remove the last node
func (l *LinkedList) Pop() *Node {
	if l.Any() {
		poppedNode := l.Nodes[l.Count-1]

		if poppedNode.Head != nil {
			poppedNode.Head.Tail = nil
			poppedNode.Head = nil
		}

		l.Count--

		l.Nodes = l.Nodes[:l.Count]

		return poppedNode
	}

	return nil
}

// NOTE: search among the nodes
func (l *LinkedList) Find(value any) *Node {
	return &Node{}
}

// NOTE: insert the node at the position
func (l *LinkedList) Insert(value any, index int) *Node {
	return &Node{}
}

func (l *LinkedList) Print() {
	tree := ""

	for _, node := range l.Nodes {
		tree += fmt.Sprintf("%v <-> ", node.Value)
	}

	if len(tree) > 0 {
		tree = tree[:len(tree)-4]
	}

	fmt.Println(tree)
}
