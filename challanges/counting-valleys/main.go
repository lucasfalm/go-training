package main

import (
	"fmt"
	"strings"
)

/*
 * Complete the 'countingValleys' function below.
 *
 * The function is expected to return an INTEGER.
 * The function accepts following parameters:
 *  1. INTEGER steps
 *  2. STRING path
 */

func countingValleys(steps int32, path string) int32 {
	var (
		v     int32
		w     = make([]int32, 0, steps)
		ps    = make([]string, 0, steps)
		rules = map[string]int32{
			"D": -1,
			"U": +1,
		}
	)

	w = append(w, 0) // coming from sea level
	ps = strings.Split(path, "")
	for i, p := range ps {
		r := rules[p]
		var result int32

		if i == 0 {
			result = r
			w = append(w, result)
		}

		if i > 0 {
			result = w[i]
			result += r
			w = append(w, result)
		}

		if result < 0 && w[i] == 0 {
			v++
		}
	}

	return v
}

func main() {
	s := "DUUDDDUUUD"
	r := countingValleys(10, s)
	fmt.Printf("%v valeys in %s\n", r, s)
}
