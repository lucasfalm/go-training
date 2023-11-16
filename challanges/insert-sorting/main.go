package main

import (
	"github.com/lucasfalm/go-training/challanges/insert-sorting/printer"
	sortingv1 "github.com/lucasfalm/go-training/challanges/insert-sorting/v1"
	sortingv2 "github.com/lucasfalm/go-training/challanges/insert-sorting/v2"
)

func main() {
	var arr = []int32{2, 3, 4, 5, 6, 7, 8, 9, 10, 1}

	printer.PrintFormatted(arr)

	sortingv1.InsertionSort1(10, arr)
	sortingv2.InsertionSort2(10, arr)
}
