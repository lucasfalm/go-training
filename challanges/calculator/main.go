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

func Calc(expression string) (result float64) {
	fmt.Println(expression)
	css := strings.Split(expression, "")
	exps, _ := searchExpressions(css)

	return exps[len(exps)-1].result
}

type expression struct {
	action      string
	firstNumber float64
	lastNumber  float64
	result      float64
}

func searchExpressions(css []string) (exps []expression, ok bool) {
	ok = true

	actionIndex := 0
	lastNumberIndex := 0
	running := true

	for running {
		var action string
		action, actionIndex, ok = defineAction(actionIndex, lastNumberIndex, css)
		if !ok {
			if len(exps) > 0 {
				return exps, ok
			}

			r, _ := strconv.ParseFloat(strings.Join(css, ""), 64)
			exps = append(exps, expression{result: r})
			ok = false

			return exps, ok
		}

		var fn float64

		if len(exps) == 0 {
			fn = extractFirstNumber(css, actionIndex)
		}

		if len(exps) > 0 {
			fn = exps[len(exps)-1].result
		}

		ln, lni, okk := extractLastNumber(css, actionIndex)
		if !okk {
			return exps, false
		}

		lastNumberIndex = lni + lastNumberIndex

		exps = append(exps, expression{
			firstNumber: fn,
			lastNumber:  ln,
			action:      action,
			result:      execAction(fn, ln, action),
		})
	}

	return
}

func defineAction(lastActionIndex int, lastNumberIndex int, css []string) (action string, actionIndex int, ok bool) {
	ok = true

	if lastActionIndex > len(css)-1 && lastNumberIndex <= len(css)-1 {
		ok = false
		return "", -1, ok
	}

	var okk bool

	if lastNumberIndex != 0 {
		action, lastActionIndex, okk = searchAction(css, lastNumberIndex)
		if !okk {
			ok = false
			return "", -1, ok
		}
	}

	if lastNumberIndex == 0 {
		action, lastActionIndex, okk = searchAction(css, lastActionIndex)
		if !okk {
			ok = false
			return "", -1, ok
		}
	}

	return action, lastActionIndex, ok
}

func searchAction(ss []string, start int) (a string, i int, ok bool) {
	ok = true

	for y, s := range ss {
		if y < start {
			continue
		}

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

func extractLastNumber(ss []string, actionIndex int) (ln float64, lni int, ok bool) {
	var (
		lns = []string{}
		ns  = []string{}
		lnf = false
		rm  = []string{}
	)

	start := 0

	if actionIndex != 0 {
		start = actionIndex + 1
	}

	rm = ss[start:]

	ok = true

	for i, s := range rm {
		if val, okk := validActions[s]; okk {
			if lnf {
				break
			}

			if rm[i+1] == " " {
				ok = false
				break
			}

			ns = append(ns, val)
		}

		if y, okk := validNumbers[s]; okk {
			lns = append(lns, s)
			lnf = true

			if actionIndex != 0 {
				lni = y + start + 1
			}

			if actionIndex == 0 {
				lni = y + start
			}
		}
	}

	if !ok {
		return
	}

	var ls string
	if len(ns) > 1 {
		var p = 0
		for p < len(ns) {
			if p == len(ns)-1 {
				break
			}

			fs := ns[p]
			ns := ns[p+1]

			if fs == "-" && ns == "-" {
				ls = "+"
			}

			if fs == "+" && ns == "-" || fs == "-" && ns == "+" {
				ls = "-"
			}

			if fs == "+" && ns == "+" {
				ls = "+"
			}

			p++
		}
	}

	if len(ns) == 1 {
		ls = ns[0]
	}

	if ls != "" {
		lns = append([]string{ls}, lns...)
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
