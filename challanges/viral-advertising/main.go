package main

import (
	"fmt"
	"math"
)

func main() {
	var days int32 = 10

	fmt.Printf("total %d likes in %d days\n", viralAdvertising(days), days)
}

// NOTE: https://www.hackerrank.com/challenges/strange-advertising/problem?isFullScreen=true
func viralAdvertising(days int32) (likes int32) {
	var shared int32

	var sharesPerLike int32 = 3

	for i := int32(1); i <= days; i++ {
		var newLikes int32 = calcLikes((shared))

		shared = newLikes * sharesPerLike
		likes += calcLikes((shared))

		if i == 1 {
			shared += 5
			likes += calcLikes((shared))
		}
	}

	return
}

func calcLikes(recipients int32) int32 {
	return (int32(math.Floor(float64(recipients / 2))))
}
