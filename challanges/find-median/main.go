package main

import (
	"fmt"
	"sort"
)

// NOTE: https://www.hackerrank.com/challenges/find-the-median/problem?isFullScreen=true
func main() {
	fmt.Println(findMedian([]int32{6, 5, 4, 3, 2, 4, 5, 10, 5, 8}))
}

func findMedian(arr []int32) int32 {
	var intArr []int

	for _, n := range arr {
		intArr = append(intArr, int(n))
	}

	sort.Ints(intArr)

	return int32(intArr[len(intArr)/2])
}
