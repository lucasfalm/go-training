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

	// Time complexity of Big O(n^2) time, and space O(n^2)
	for i := range sBooks {
		for y, v := range books {
			if i == y {
				sBooks[i] = v
			} else {
				continue
			}
		}
	}

	fmt.Printf("Books: %v\n", books)
	fmt.Printf("Summer Books: %v\n", sBooks)
}
