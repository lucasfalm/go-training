package main

import (
	"fmt"
	"os"
	"strings"
)

var (
	args    = os.Args[1:]
	heights = []int32{
		1, 3, 1, 3, 1, 4, 1, 3, 2, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5,
	}
)

func main() {
	if len(args) >= 1 {
		handleArgs()
	} else {
		fmt.Println("write a word")
	}
}

func handleArgs() {
	for i, value := range args {
		if i == 0 {
			word := value

			fmt.Println("--- result ---")
			fmt.Printf("word given: %v\n", word)
			fmt.Printf("the PDF mm2 is: %v\n", designerPdfViewer(heights, word))
			fmt.Println("--- result ---")
		}
	}
}

// NOTE: https://www.hackerrank.com/challenges/designer-pdf-viewer/problem?isFullScreen=true

/*
 * Complete the 'designerPdfViewer' function below.
 *
 * The function is expected to return an INTEGER.
 * The function accepts following parameters:
 *  1. INTEGER_ARRAY h
 *  2. STRING word
 */

// NOTE: sum of heights * highest height = result
func designerPdfViewer(h []int32, word string) int32 {
	totalHeightByLetter := getTotalHeightByLetter(h, word)

	max := calcMaxHeightOf(totalHeightByLetter)
	total := int32(len(totalHeightByLetter))

	return max * total
}

func getTotalHeightByLetter(h []int32, word string) map[int]int32 {
	alphabet := make(map[string]int32)

	for i := 'a'; i <= 'z'; i++ {
		alphabet[string(i)] = int32(i - 'a')
	}

	totalHeightByLetter := make(map[int]int32)

	for i, word := range strings.Split(word, "") {
		totalHeightByLetter[i] += h[alphabet[word]]
	}

	return totalHeightByLetter
}

func calcMaxHeightOf(totalHeightByLetter map[int]int32) int32 {
	var max int32

	for _, height := range totalHeightByLetter {
		if height > max {
			max = height
		}
	}

	return max
}

func calcTotalHeightOf(totalHeightByLetter map[int]int32) int32 {
	var total int32

	for _, height := range totalHeightByLetter {
		total += height
	}

	return total
}
