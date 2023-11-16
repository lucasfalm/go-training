package main

import (
	"fmt"
	"strconv"
	"strings"
)

// NOTE: https://www.hackerrank.com/challenges/insertionsort1/problem?isFullScreen=true
func main() {
	var arr = []int32{2, 3, 4, 5, 6, 7, 8, 9, 10, 1}

	printFormatted(arr)

	insertionSort1(10, arr)
}

func insertionSort1(n int32, arr []int32) {
	var (
		pv int32 = arr[n-1]
		i  int32 = n - 1
	)

	arr[i] = arr[i-1]

	printFormatted(arr)

	for {
		if i < 0 {
			break
		}

		if i == 1 || arr[i-2] < pv {
			arr[i-1] = pv
			printFormatted(arr)
			break
		}

		i--
		arr[i] = arr[i-1]

		printFormatted(arr)
	}
}

func printFormatted(arr []int32) {
	var f = make([]string, len(arr))

	for i, n := range arr {
		f[i] = strconv.Itoa(int(n))
	}

	fmt.Println(strings.Join(f, " "))
}
