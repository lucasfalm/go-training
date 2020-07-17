package main

import (
	"fmt"
	"sort"
)

func main() {
	elements := []int{-100, 1}
	fmt.Println("Solution: ", Solution(elements))
}

func Solution(A []int) int {
	sort.Ints(A)
	start, stop, missing, perfectCount, lowerValue := A[0], A[len(A)-1:], 0, make(map[int]bool), 0

	if len(A) == 1 {
		if A[0] > 2 {
			return A[0] - 1
		} else if A[0] == 1 {
			return 2
		} else {
			return 1
		}
	} else {
		if start < 0 && stop[0] < 0 {
			missing = 1
		} else if start < 0 && stop[0] == 1 {
			missing = stop[0] + 1
		} else {
			for i := start; i <= stop[0]; i++ {
				perfectCount[i] = false
			}

			for i := 0; i <= len(A)-1; i++ {
				actualValue := A[i]

				if _, found := perfectCount[actualValue]; found {
					perfectCount[actualValue] = true
				}
			}

			before := 0

			for missingValue, status := range perfectCount {
				if !status {
					if missing == 0 {
						missing = missingValue
						lowerValue = missing
					} else {
						before = missing
						missing = missingValue

						if before < missing {
							lowerValue = before
						} else {
							lowerValue = missing
						}
					}
				}
			}

			if start > 1 && lowerValue > 1 {
				lowerValue = 1
			}
			missing = lowerValue
		}
	}

	if missing == 0 {
		if start > 1 && stop[0] > 1 {
			missing = 1
		} else if start < 0 && stop[0] > 1 {
			missing = stop[0] + 1
		} else {
			missing = start + 1
		}
	}

	return missing
}
