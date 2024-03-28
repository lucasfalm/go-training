package linked_list_core

import "fmt"

// NOTE: search among the nodes and print result
func (l *LinkedList) FindWithPrint(value any) *Node {
	result := l.Find(value)

	msg := ""

	if result != nil {
		msg += fmt.Sprintf("found node: %v", result.Value)

		if result.Head != nil {
			msg += fmt.Sprintf(" | head: %v", result.Head.Value)
		}

		if result.Tail != nil {
			msg += fmt.Sprintf(" | tail: %v", result.Tail.Value)
		}

		fmt.Println(msg)
	} else {
		fmt.Println("node has not been found")
	}

	return result
}
