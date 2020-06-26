package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args[1:]
	for _, w := range args {
		words := Solution(w)
		fmt.Printf("the word '%v', is: %v\n", w, words)
	}
}
func Solution(str string) []string {
	var groups []string
	f := 2

	for i := 0; i <= len(str); i += 2 {
		if i == len(str) {
			break
		}

		var w string

		if len(str) < f {
			f -= 1
		}
		w = str[i:f]

		if w == "" {
			break
		}

		if len(w) < 2 {
			w = w + "_"
		}

		groups = append(groups, w)
		f += 2
	}

	return groups
}
