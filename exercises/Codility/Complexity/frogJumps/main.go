package main

import "fmt"

func main() {
	fmt.Println(Solution(10, 1000, 1000))
}

func Solution(X int, Y int, D int) int {
	step := 0

	if X < Y {
		for result := X; result < Y; result += D {
			step++
		}
	} else {
		step = 0
	}

	return step

}
