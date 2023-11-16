package sortingv1

import "github.com/lucasfalm/go-training/challanges/insert-sorting/printer"

// NOTE: https://www.hackerrank.com/challenges/insertionsort1/problem?isFullScreen=true
func InsertionSort(n int32, arr []int32) {
	var (
		pv int32 = arr[n-1]
		i  int32 = n - 1
	)

	arr[i] = arr[i-1]

	printer.PrintFormatted(arr)

	for {
		if i < 0 {
			break
		}

		if i == 1 || arr[i-2] < pv {
			arr[i-1] = pv
			printer.PrintFormatted(arr)
			break
		}

		i--
		arr[i] = arr[i-1]

		printer.PrintFormatted(arr)
	}
}
