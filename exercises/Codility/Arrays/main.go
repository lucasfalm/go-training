package main

import "fmt"

func main() {
	numbers := []int{1, 2, 3, 4} // [len(numbers)-y:len(numbers - i)]
	key := 13
	fmt.Println("Original", numbers)
	fmt.Println("Solution: ", Solution(numbers, key))
}

func Solution(A []int, K int) []int {
	result := []int{}
	x, y := 1, 0

	for i := 1; i <= K; i++ {
		last := A[len(A)-x : len(A)-y]
		fmt.Println(last)
		x++
		y++

		result = append([]int{last[0]}, result...)

		if y >= len(A) {
			x, y = 1, 0
		}
	}

	if len(A) >= K {
		less := A[:len(A)-len(result)]
		for _, v := range less {
			result = append(result, v)
		}
	}

	return result[:len(A)]
}
