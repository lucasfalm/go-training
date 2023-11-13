package main

import (
	"fmt"
	"math"
	"strings"
)

// NOTE: https://www.hackerrank.com/challenges/repeated-string/problem
func repeatedString(s string, n int64) int64 {
	// var v1r int64 = v1(s, n) // NOTE: 2020;
	var v2r int64 = v2(s, n) // NOTE: 2023;

	fmt.Printf("the letter 'a' appeared %v times in '%v' * %v\n", v2r, s, n)

	return v2r
}

/*
	To correctly solve this problem it's needed to check whether or not
	it's necessary to round the integer to up, or down. (0.7 or lower down, 0.8 or 0.9 up)

	Using math.Round(up) or just casting into int64 (down)
*/
// Big O(n) time( where n is length of string) and Big O(n) space. 16/23 tests
func v1(s string, n int64) int64 {
	var (
		lettersSlice = strings.Split(s, "")
		repeat       int
	)

	for _, letter := range lettersSlice {
		if letter == "a" {
			repeat++
		}
	}

	var r int64 = int64(math.Round(float64(repeat) / (float64(len(lettersSlice))) * float64(n)))

	return r
}

func v2(s string, n int64) int64 {
	var searchedLetter string = "a"

	var slc int64 = int64(0)
	var tl int64 = int64(0)

	var sS []string = strings.Split(s, "")

	if len(sS) == 1 {
		if s == searchedLetter {
			return 1 * n
		} else {
			return 0
		}
	}

	var x int = int(n) / len(sS)
	var xF float64 = float64(n) / float64(len(sS))

	var hasRest bool = xF-math.Trunc(xF) > 0

	for _, letter := range sS {
		if letter == searchedLetter {
			slc++
		}
	}

	slc = slc * int64(x)
	tl = int64(len(sS) * x)

	if hasRest {
	while:
		for {
			if tl == n {
				break while
			}

			for _, letter := range sS {
				if tl == n {
					break while
				}

				tl++

				if letter == searchedLetter {
					slc++
				}
			}
		}
	}

	return slc
}

func main() {
	repeatedString("abc", 100)
}
