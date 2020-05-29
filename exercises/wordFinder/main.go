package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	const (
		corpus = "lazy cat jumps again and again and again"
	)
	var (
		words = strings.Fields(corpus)
		query = os.Args[1:]
		flag  int
	)
	if len(query) >= 1 {
		for _, v := range query {
			for i, w := range words {
				if v == w {
					fmt.Printf("The %q is at position %v: %q\n", v, i, w)
					flag++
				} else {
					if i == (len(words)-1) && flag == 0 {
						fmt.Printf("We dont find any word in match. Try again\n")
						break
					} else {
						continue
					}
				}
			}
		}
	} else {
		fmt.Println("Please, write at least one word to search.")
	}
}
