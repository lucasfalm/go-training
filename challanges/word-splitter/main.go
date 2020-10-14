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
	var wordGroups []string
	rangeValues := 2

	for index := 0; index <= len(str); index += 2 {
		if index == len(str) {
			break
		}

		var formedWord string

		if len(str) < rangeValues {
			rangeValues -= 1
		}
		formedWord = str[index:rangeValues]

		if formedWord == "" {
			break
		}

		if len(formedWord) < 2 {
			formedWord = formedWord + "_"
		}

		wordGroups = append(wordGroups, formedWord)
		rangeValues += 2
	}

	return wordGroups
}
