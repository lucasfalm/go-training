package main

import (
	"fmt"
	"sort"
)

func main() {
	elements := []int{5}
	fmt.Println("Solution: ", Solution(elements))
}

func Solution(A []int) int {
	perfect := []int{}

	if len(A) > 0 {
		if len(A) == 1 {
			return A[0] + 1
		}
		sort.Ints(A)
		perfectCount, missingCount := 0, 0
		first, last := A[:1], A[len(A)-1:]

		for i := first[0]; i <= last[0]; i++ {
			perfect = append(perfect, i)
		}

		for _, number := range perfect {
			perfectCount += number
		}

		for _, number := range A {
			missingCount += number
		}

		result := perfectCount - missingCount

		if result == 0 {
			if first[0] == 2 {
				return 1
			} else {
				return last[0] + 1
			}
		}

		return result

	} else {
		return 1
	}
}
