package main

import (
	"fmt"
	"math"
)

// NOTE: https://www.hackerrank.com/challenges/sherlock-and-squares/problem?isFullScreen=true
func main() {
	var a, b int32 = 135964528, 323753587

	fmt.Printf("how many squares in: from %v to %v are %v squares\n", a, b, squares(a, b))
}

func squares(a int32, b int32) int32 {
	squaresTrack := int32(0)

	for number := a; number <= b; number++ {
		if isNumberSquarePerfect(number) {
			squaresTrack++
		}
	}

	return squaresTrack
}

var cacheSquares = make(map[int32]float64)

func isNumberSquarePerfect(number int32) bool {
	// sqrtN, ok := cacheSquares[number]

	// if ok {
	//   return sqrtN == math.Floor(sqrtN)
	// }

	sqrtN := math.Sqrt(float64(number))
	// cacheSquares[number] = sqrtN

	return sqrtN == math.Floor(sqrtN)
}
