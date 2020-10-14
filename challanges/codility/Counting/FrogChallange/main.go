package main

import (
	"fmt"
)

func main() {
	elements := []int{1, 2, 3, 4, 5, 6, 7, 8}
	k := 8
	fmt.Println("Solution: ", Solution(k, elements))
}

func Solution(X int, A []int) int {
	arraySum, check, perfectSum := 0, make(map[int]bool), (X * (X + 1) / 2)

	for index, value := range A {
		if !check[value] {
			arraySum += value
			check[value] = true
		}

		if arraySum == perfectSum {
			return index
		}
	}

	return -1
}
