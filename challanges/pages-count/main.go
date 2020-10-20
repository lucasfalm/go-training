package main

// https://www.hackerrank.com/challenges/drawing-book/problem
// Big O(n) time and O(n) space
func pageCount(n int32, p int32) int32 {
	var (
		frontCount int32
		backCount  = n

		totalFront int32
		totalBack  int32

		flag bool
	)

front:
	for frontCount < p {
		if flag {
			totalFront++
			flag = false
			frontCount++

			continue front
		}
		frontCount++
		flag = true
	}

back:
	for backCount > p {
		if flag {
			totalBack++
			flag = false
			backCount--

			continue back
		}

		backCount--
		flag = true
	}

	if totalBack < totalFront {
		return totalBack
	} else {
		return totalFront
	}
}
