package main

import (
	"fmt"
	"sort"
)

func main() {
	elements := []int{-1, 2, 3, 4, 5, 6, 7, 1000}
	fmt.Println("Solution: ", Solution(elements))
}

// 75% O(N) or O(N * log(N))
func Solution(A []int) int {
	totalSliceLenght, missing := len(A), 0

	if totalSliceLenght >= 1 {
		sort.Ints(A)
		start, stop := A[0], A[len(A)-1:]

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
			} else if start < 0 && stop[0] < 0 {
				return 1
			} else if start < 0 && stop[0] > 1 {
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

				lastValue := 0
				for _, v := range A {
					beforeSet := lastValue
					lastValue = v

					if beforeSet-lastValue <= -2 {
						if beforeSet+1 == 0 && lastValue != 1 {
							return 1
						} else if beforeSet+1 == 0 && lastValue == 1 {
							continue
						} else {
							return beforeSet + 1
						}
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
