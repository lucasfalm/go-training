package main

import (
	"fmt"
	"math"

	"github.com/lucasfalm/go-training/challanges/day-of-programmer/components"
)

// NOTE: https://www.hackerrank.com/challenges/day-of-the-programmer/problem?isFullScreen=true
func main() {
	fmt.Println(dayOfProgrammer(1918))
}

func dayOfProgrammer(year int32) string {
	var (
		r                   string
		gCalendarTransition int32 = 1918
		isTransitionYear    bool
	)

	if components.CalendarType(year) == "gregorian" {
		if year == gCalendarTransition {
			isTransitionYear = true
		}
	}

	r = findDay(year, map[string]interface{}{
		"isTransitionYear": isTransitionYear,
	})

	return r
}

func findDay(year int32, opts map[string]interface{}) string {
	var (
		r               string
		dayOfProgrammer int32 = 256
		sum             int32 = 0
	)

	for _, month := range components.ListMonths() {
		if month == "february" {
			if opts["isTransitionYear"].(bool) {
				tYopt := components.Calendar[month]["transitionYear"].(map[string]interface{})
				sum += (tYopt["to"].(int32) - tYopt["from"].(int32)) + 1
			} else {
				if components.IsLeap(year) {
					sum += components.Calendar[month]["leap"].(map[string]interface{})["to"].(int32)
				} else {
					sum += components.Calendar[month]["nonLeap"].(map[string]interface{})["to"].(int32)
				}
			}
		} else {
			if components.IsLeap(year) {
				sum += components.Calendar[month]["to"].(int32)
			} else {
				sum += components.Calendar[month]["to"].(int32)
			}
		}

		if sum >= dayOfProgrammer {
			var d float64 = math.Abs(float64(sum - components.Calendar[month]["to"].(int32) - dayOfProgrammer))
			var mI int32

			for i, m := range components.ListMonths() {
				if m == month {
					mI = int32(i) + 1
					break
				}
			}

			r = fmt.Sprintf("%v.%02d.%v", d, mI, year)
			break
		}
	}

	return r
}
