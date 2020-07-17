package main

import (
	"fmt"
	"sort"
)

func main() {
	elements := []int{1, 1, 2, 3, 4, 6, 7, 9}
	fmt.Println("Solution: ", Solution(elements))
}

func Solution(A []int) int {
	totalSliceLenght, missing := len(A), 0

	if totalSliceLenght >= 1 {
		sort.Ints(A)
		start, stop, perfectCount := A[0], A[len(A)-1:], make(map[int]bool)

		if totalSliceLenght == 1 {
			if start > 2 {
				return start - 1
			} else if start == 1 {
				return 2
			} else {
				return 1
			}
		} else if totalSliceLenght == 2 {
			if start > 1 && stop[0] > 1 {
				return 1
			} else {
				return start + 1
			}
		} else {
			if start < 0 && stop[0] < 0 {
				missing = 1
			} else {
				if start > 1 && stop[0] > 1 {
					return 1
				}

				for increment := start; increment <= stop[0]; increment++ {
					perfectCount[increment] = false
					if increment == 0 {
						perfectCount[increment] = true
					}
				}

				for i := 0; i <= len(A)-1; i++ {
					actualValue := A[i]

					if _, found := perfectCount[actualValue]; found {
						perfectCount[actualValue] = true
					}
				}

				for value, _ := range perfectCount {
					if !perfectCount[value] {
						missing = value
						break
					}
				}
			}

			if missing == 0 {
				if start > 1 && stop[0] > 1 {
					missing = 1
				} else if start < 0 && stop[0] > 1 {
					missing = stop[0] + 1
				} else {
					missing = stop[0] + 1
				}
			} else if missing < 0 {
				missing = 1
			}
		}
	} else {
		return 1
	}

	return missing
}
