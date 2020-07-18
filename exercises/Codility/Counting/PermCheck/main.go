package main

import (
	"fmt"
	"sort"
)

func main() {
	elements := []int{1, 2, 3, 3, 3}
	fmt.Println("Solution: ", Solution(elements))
}

func Solution(A []int) int {
	sort.Ints(A)
	actual, result, last := 0, 0, 0

	if len(A) > 1 {
		first := A[0]

		if first > 1 {
			return 0
		}
		for _, value := range A {
			last = actual
			actual = value

			equation := last - actual

			if equation < -1 || equation > 1 {
				if actual == 1 && last == -1 {
					continue
				}
				return 0
			} else if equation == 0 {
				if actual == 1 {
					return 0
				}
				result = 1
			} else {
				result = 1
			}
		}
	} else if len(A) == 1 && A[0] == 1 {
		return 1
	} else {
		return 0
	}

	return result
}
