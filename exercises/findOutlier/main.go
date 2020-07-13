package main

import (
	"fmt"
	"os"
	"strconv"
)

var (
	args = os.Args[1:]
)

func main() {
	numbersList := []int{}
	for _, v := range args {
		convert, _ := strconv.Atoi(v)
		numbersList = append(numbersList, convert)
	}

	result := FindOutlier(numbersList)
	fmt.Println("Outlier is:", result)
}
func FindOutlier(integers []int) int {
	odd, oddn, even, evenn := 0, 0, 0, 0

finding:
	for _, number := range integers {
		if number % 2 != 0 {
			odd++
			oddn = number
			continue finding
		}

		even++
		evenn = number
	}

	if odd > even {
		return evenn
	}

	return oddn
}
