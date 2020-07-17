package main

import (
	"fmt"
	"sort"
)

func main() {
	elements := []int{1, 2, 3, 4, 5, 6, 7, 8}
	fmt.Println("Solution: ", Solution(elements))
}

func Solution(A []int) int {
	sort.Ints(A)
	start, stop, missing, perfectCount := A[0], A[len(A)-1:], 0, make(map[int]bool)

	if start > 0 && stop[0] > 0 {
		for i := start; i <= stop[0]; i++ {
			perfectCount[i] = false
		}

		for i := 0; i <= len(A)-1; i++ {
			actualValue := A[i]

			if _, found := perfectCount[actualValue]; found {
				perfectCount[actualValue] = true
			}
		}

		for missingValue, status := range perfectCount {
			if !status {
				missing = missingValue
			}
		}
		if missing == 0 {
			missing = stop[0] + 1
		}

		return missing
	} else {
		return 1
	}
}
