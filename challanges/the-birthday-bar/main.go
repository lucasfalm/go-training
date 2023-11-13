package main

import "fmt"

func main() {
	var sum int32 = 1
	var with int32 = 1

	var r int32 = birthday([]int32{1, 1, 2, 5, 2, 1, 3, 4, 5}, sum, with)

	fmt.Printf("you can have %v ways to divide the chocolate\n", r)
}

// NOTE: https://www.hackerrank.com/challenges/the-birthday-bar/problem?isFullScreen=true
func birthday(s []int32, sum int32, with int32) int32 {
	var r int32

	// NOTE: pointer zero & walker pointer
	var pZ, wP int = 0, 0

	// NOTE: current total integer's sum
	var iS int32 = 0

	// NOTE: current chocolate slices count
	var cS int32 = 0

	for pZ <= len(s)-1 {
		if iS < with && wP <= len(s)-1 {
			cS += s[wP]
			iS++

			// NOTE: moving forward
			wP++

		} else {
			// NOTE: achieve the desired sum with this amount of chocolate slices
			if cS == sum {
				r++
			}

			// NOTE: resetting
			cS = 0
			iS = 0

			// NOTE: moving pointer zero forward
			pZ++

			// NOTE: resetting walker pointer
			wP = pZ
		}
	}

	return r
}
