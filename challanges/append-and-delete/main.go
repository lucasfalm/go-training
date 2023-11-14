package main

import (
	"fmt"
	"math"
	"strings"
)

func main() {
	fmt.Println(appendAndDelete("abc", "def", 3))
}

// NOTE: https://www.hackerrank.com/challenges/append-and-delete/problem?isFullScreen=true - 10/13
func appendAndDelete(s string, t string, k int32) string {
	var y string = "Yes"
	var n string = "No"

	var diff int32 = int32(math.Abs(float64(len(s) - len(t))))

	if diff > k {
		return n
	}

	var p int = 0
	var c int = 0

	var sS = strings.Split(s, "")
	var sT = strings.Split(t, "")

	for {
		if len(sS)-1 < p && len(sT)-1 < p {
			break
		}

		// NOTE: sS bigger than sT
		if len(sT)-1 < p && len(sS)-1 >= p {
			c += len(sS[p:])
			break
		}

		// NOTE: sT bigger than sS
		if len(sT)-1 < p && len(sS)-1 >= p {
			c += len(sT[p:])
			break
		}

		if sS[p] != sT[p] {
			c += len(sS[p:])
			c += len(sT[p:])
			break
		}

		p++
	}

	if c > int(k) {
		return n
	}

	return y
}
