package main

import "fmt"

//https://www.hackerrank.com/challenges/equality-in-a-array/problem
// Big O(n) time, and O(2n) space

func equalizeArray(arr []int32) int32 {
	var (
		rankRepeated = map[int32]int{}
		maxRepeated  int
	)

	for _, actualNumber := range arr {
		rankRepeated[actualNumber]++
	}

	for _, repeated := range rankRepeated {
		if repeated > maxRepeated {
			maxRepeated = repeated
		}
	}

	return int32((len(arr)) - maxRepeated)
}

func main() {
	fmt.Println(equalizeArray([]int32{1, 2, 5, 9, -1}))
}
