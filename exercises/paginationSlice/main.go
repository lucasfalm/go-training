package main

import (
	s "github.com/inancgumus/prettyslice"
)

func main() {
	items := []string{"a", "b", "c", "d", "e", "f", "g", "h", "i"}

	const (
		pageSize = 4
	)
	var (
		l = len(items)
	)

	for from := 0; from < l; from += pageSize {
		to := from + pageSize
		if to > l {
			to = l
		}
		// slicing a slice
		currentPage := items[from:to]

		s.Show("", currentPage)
	}
}
