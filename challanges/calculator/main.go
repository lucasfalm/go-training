package calculator

import (
	"fmt"
	"strconv"
	"strings"
)

var validActions = map[string]string{
	"+": "+",
	"-": "-",
	"*": "*",
	"/": "/",
}

var validNumbers = map[string]int{
	"0": 0,
	"1": 1,
	"2": 2,
	"3": 3,
	"4": 4,
	"5": 5,
	"6": 6,
	"7": 7,
	"8": 8,
	"9": 9,
}

func Calc(expression string) float64 {
	fmt.Println(expression)
	css := strings.Split(expression, "")
	action, actionIndex, ok := searchAction(css)
	if !ok {
		r, _ := strconv.ParseFloat(expression, 64)
		return r
	}

	fn := extractFirstNumber(css, actionIndex)
	ln, ok := extractLastNumber(css, actionIndex)
	if !ok {
		return 0.0
	}

	return execAction(fn, ln, action)
}

func searchAction(ss []string) (a string, i int, ok bool) {
	ok = true

	for y, s := range ss {
		if val, okk := validActions[s]; okk {
			if y == 0 {
				continue
			}

			a = val
			i = y
			break
		}
	}

	if a == "" {
		ok = false
	}

	return
}

func extractFirstNumber(ss []string, actionIndex int) (fn float64) {
	fn, _ = strconv.ParseFloat(strings.Join(gsub(ss[0:actionIndex]), ""), 64)
	return
}

func extractLastNumber(ss []string, actionIndex int) (ln float64, ok bool) {
	var (
		lns = []string{}
		rm  = ss[actionIndex+1:]
	)

	ok = true

	for i, s := range rm {
		if val, okk := validActions[s]; okk {
			if rm[i+1] == " " {
				ok = false
				break
			}

			lns = append(lns, val)
		}

		if _, okk := validNumbers[s]; okk {
			lns = append(lns, s)
		}
	}

	if !ok {
		return
	}

	ln, _ = strconv.ParseFloat(strings.Join(lns, ""), 64)
	return
}

func execAction(fn, ln float64, action string) (r float64) {
	switch action {
	case "+":
		r = float64(fn + ln)
	case "-":
		r = float64(fn - ln)
	case "*":
		r = float64(fn * ln)
	case "/":
		if fn == 0 || ln == 0 {
			r = 0.0
			return
		}

		r = float64(fn / ln)
	}

	return r
}

func gsub(ss []string) []string {
	css := make([]string, 0, len(ss))

	for _, s := range ss {
		if s != " " && s != "(" && s != ")" {
			css = append(css, s)
		}
	}

	return css
}
