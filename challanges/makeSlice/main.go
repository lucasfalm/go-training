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
	tasks := make([]string, len(mMap), 10)
	// here my backing array is always the same, because the capacity of my backing array is 10
	copy(tasks, mMap)

	s.Show("Tasks", tasks)

	s.Show("mMap", mMap)
}
