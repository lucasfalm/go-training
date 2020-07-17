package main

import "fmt"

func main() {
	numbers := []int{1, 2, 3, 1, 3, 5, 4, 5, 4}
	fmt.Printf("Alone number: %v\n", Solution(numbers))
}

func Solution(A []int) int {
	occurrencesMap, alone := make(map[int]int), 0

	for _, value := range A {
		occurrencesMap[value]++
	}

	for number, occurrences := range occurrencesMap {
		if occurrences%2 != 0 {
			alone = number
		}
	}

	return alone
}
