package main

import "fmt"

func main() {
	var t int64 = 99999997668

	fmt.Println(strangeCounter(t))
}

// NOTE: https://www.hackerrank.com/challenges/strange-code/problem?isFullScreen=true
func strangeCounter(t int64) int64 {
	var r int64

	// r = v1(t)
	r = v2(t)

	return r
}

func v1(t int64) int64 {
	var cT int64 = 1
	var pInitialV int64 = 3
	var cV int64 = pInitialV

	var r int64

heart:
	for {
		if cT == t {
			r = cV
			break heart
		}

		if cV == 1 {
			pInitialV = pInitialV * 2
			cV = pInitialV
			cT++
			continue heart
		}

		cT++
		cV--

	}

	return r
}

func v2(t int64) int64 {
	var r int64

	var cV int64 = 3  // NOTE: initial value in the cycle
	var iCt int64 = 1 // NOTE: initial time in the cycle
	var xCt int64 = 0 // NOTE: final time in the cycle

heartLoop:
	for {
		if iCt == 1 && xCt == 0 {
			xCt = cV

			// NOTE: max time in the cycle covers the searched time?
			if hasFoundTime(xCt, t) {
				r = calcR(cV, t, iCt)
				break heartLoop
			}

			continue heartLoop
		}

		// NOTE: starting new cycle
		cV = cV * 2
		iCt = xCt + 1
		xCt = cV + xCt

		if hasFoundTime(xCt, t) {
			r = calcR(cV, t, iCt)
			break heartLoop
		}

	}

	return r
}

// NOTE: max time in the cycle covers the searched time?
func hasFoundTime(xCt, t int64) bool {
	return xCt >= t
}

func calcR(cV, t, iCt int64) int64 {
	return cV - (t - iCt)
}
