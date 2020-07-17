package main

import "fmt"

func main() {
	elements := []int{7}
	fmt.Println("Solution: ", Solution(elements))
}

func Solution(A []int) int {
	if len(A) >= 0 {
		if len(A) == 1 {
			return A[0] + 1
		} else {
			start, end, missing := A[:1], A[len(A)-1:], 0

			for searchedNumber := start[0]; searchedNumber <= end[0]; searchedNumber++ {
				found := false

				for _, number := range A {
					if number == searchedNumber {
						found = true
					}
				}
				if !found {
					missing = searchedNumber
				}

				if missing == 0 {
					missing = end[0] + 1
				}
			}

			return missing
		}
	} else {
		return 0
	}
}
