package main

import "fmt"

func main() {
	numbers := []int{1, 2, 3, 4}
	key := 3

	fmt.Println(Solution(numbers, key))
}

func Solution(A []int, K int) []int {
	result := []int{}

	for i := 1; i <= K; i++ {
		if len(A) >= 1 {
			nextLength := len(A) - 1
			last := A[nextLength:]
			A = A[:nextLength]
			result = append([]int{last[0]}, result...)
		}
	}

	if len(A) > 0 {
		for _, number := range A {
			result = append(result, number)
		}
	}

	return result
}
