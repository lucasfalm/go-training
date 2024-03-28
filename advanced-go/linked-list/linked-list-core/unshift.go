package linked_list_core

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
