package main

import "fmt"

func main() {
	const (
		max = 4
	)
	books := []string{"book A", "book B", "book C", "book D"}

	for i := 0; i < max; i++ {
		fmt.Println(books[i])
	}

	cutSlice := books[0:2]
	for _, v := range cutSlice {
		fmt.Println(v)
	}
}
