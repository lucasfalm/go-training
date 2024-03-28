package linked_list_core

// NOTE: search among the nodes
func (l *LinkedList) Find(value any) *Node {
	if l.Any() {
		for _, node := range l.Nodes {
			if node.Value == value {
				return node
			}
		}
	}

	return nil
}
