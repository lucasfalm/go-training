package main

import (
	"fmt"
	"os"
	"strconv"
)

// https://www.hackerrank.com/challenges/drawing-book/problem
// Big O(n) time and O(n) space
func pageCount(n int32, p int32) int32 {
	var (
		frontCount int32
		backCount  = n

		totalFront int32
		totalBack  int32

		flag bool
	)

front:
	for frontCount < p {
		if flag {
			totalFront++
			flag = false
			frontCount++

			continue front
		}
		frontCount++
		flag = true
	}

back:
	for backCount > p {
		if flag {
			totalBack++
			flag = false
			backCount--

			continue back
		}
		backCount--
		flag = true
	}

	if totalBack < totalFront {
		fmt.Println("from back")
		return totalBack
	} else {
		fmt.Println("from front")
		return totalFront
	}
}

func main() {
	if len(os.Args) > 1 {
		var (
			totalPages, _   = strconv.Atoi(os.Args[1])
			searchedPage, _ = strconv.Atoi(os.Args[2])
		)

		result := pageCount(int32(totalPages), int32(searchedPage))
		fmt.Println(result)
	} else {
		fmt.Println("write pages and searched page")
	}
}
