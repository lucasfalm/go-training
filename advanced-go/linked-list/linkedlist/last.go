package linkedlist

// NOTE: shows the last node
func (l *LinkedList) Last() *Node {
	if l.Any() {
		return l.Nodes[len(l.Nodes)-1]
	}

	return nil
}
