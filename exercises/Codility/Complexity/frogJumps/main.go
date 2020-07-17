package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(Solution(10, 85, 30))
}

func Solution(X int, Y int, D int) int {
	step := 0

	if X < Y {
		step = int(math.Round((float64(Y-X) / float64(D))))
	} else {
		step = 0
	}

	return step

}
