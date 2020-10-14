package main

import "fmt"

func main() {
	numbers := []int{} // [len(numbers)-y:len(numbers - i)]
	key := 13
	fmt.Println("Original", numbers)
	fmt.Println("Solution: ", Solution(numbers, key))
}

func Solution(A []int, K int) []int {
	result := []int{}

	if len(A) == 0 {
		return result
	}

	start, end := 1, 0

	for i := 1; i <= K; i++ {
		last := A[len(A)-start : len(A)-end]
		result = append([]int{last[0]}, result...)
		start++
		end++

		if end >= len(A) {
			start, end = 1, 0
		}
	}

	if len(A) >= K {
		missing := A[:len(A)-len(result)]
		for _, number := range missing {
			result = append(result, number)
		}
	}

	return result[:len(A)]
}
