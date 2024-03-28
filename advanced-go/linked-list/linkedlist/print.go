package linkedlist

import "fmt"

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
