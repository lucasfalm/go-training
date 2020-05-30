package main

import (
	"fmt"
	"strconv"
)

func main() {
	const (
		winter = 1
		summer = 3
		yearly = summer + winter
	)

	var (
		books  [yearly]string
		sBooks [summer]string
		//wBooks [winter]string
	)

	for i := range books {
		books[i] = "Book ID: " + strconv.Itoa(i+1)
	}

	for _, v := range books {
		for y := range sBooks {
			sBooks[y] = v
		}
	}

	fmt.Printf("Books: %v\n", books)
	fmt.Printf("Summer Books: %v\n", sBooks)
}
