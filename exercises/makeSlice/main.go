package main

import (
	s "github.com/inancgumus/prettyslice"
)

func main() {
	s.PrintBacking = true
	s.MaxPerLine = 10

	var (
		mMap = []string{"a", "b", "c", "d", "e"}
	)

	tasks := make([]string, 0, 10)

	for _, y := range mMap {
		tasks = append(tasks, y)
	}

	s.Show("Tasks", tasks)
}
