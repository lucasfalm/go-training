package linked_list_core

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
