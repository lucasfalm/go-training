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
	// define lenght and capacity of slice, 0 and 10
	tasks := make([]string, 0, 10)
	// here my backing array is always the same, because the capacity of my backing array is 10
	for _, y := range mMap {
		tasks = append(tasks, y)
	}

	s.Show("Tasks", tasks)
}
