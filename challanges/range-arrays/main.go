package main

import (
	"fmt"
)

func main() {
	const (
		winter = 1
		summer = 3
		yearly = summer + winter
	)

	var (
		sBooks [summer]string
		//wBooks [winter]string
	)
	books := [yearly]string{
		"Book 1",
		"Book 2",
		"Book 3",
		"Book 4",
	}

	// Time complexity of Big O(n) time, and space O(n) - refactored
	for i := 0; i <= len(books); i++ {
		if sBooks[i] == books[i] {
			sBooks[i] = books[i]
		}
		i++
	}

	fmt.Printf("Books: %v\n", books)
	fmt.Printf("Summer Books: %v\n", sBooks)
}
