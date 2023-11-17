package main

import (
	"fmt"
	"math"
)

// NOTE: https://www.hackerrank.com/challenges/picking-numbers/problem
func main() {
	fmt.Println(pickingNumbers([]int32{4, 6, 5, 3, 3, 1}))
}

func pickingNumbers(a []int32) int32 {
	var (
		pf, pw      int = 0, 0
		subs        [][]int32
		csub        []int32
		walkingBack bool = false
	)

	for pf <= len(a)-1 {
		c := a[pw]

		if canAppend(c, csub) {
			csub = append(csub, c)
		}

		pw++

		if pw > len(a)-1 && !walkingBack || walkingBack && pw == pf {

			if pf > 0 && !walkingBack {
				pw = 0
				walkingBack = true
				continue
			}

			if satisfySub(csub) {
				subs = append(subs, csub)
			}

			csub = []int32{}
			pf++
			pw = pf
			walkingBack = false
			continue
		}
	}

	fmt.Println(subs)

	var r int32 = -1

	for _, sub := range subs {
		cl := int32(len(sub))
		if r == -1 || cl > r {
			r = cl
		}
	}

	return r
}

func canAppend(n int32, sub []int32) bool {
	var r bool = true

	if len(sub) == 0 {
		return r
	}

	for _, subn := range sub {
		if math.Abs(float64(n-subn)) > 1 {
			r = false
			break
		}
	}

	return r
}

func satisfySub(sub []int32) bool {
	return len(sub) > 1
}
