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

type Expression struct {
	action      string
	firstNumber float64
	lastNumber  float64
	result      float64
}

func searchExpressions(css []string) (exps []Expression, ok bool) {
	ok = true

	holdMult := false
	holdDiv := false

	actionIndex := 0
	lastNumberIndex := 0

	for {
		var action string
		action, actionIndex, ok = defineAction(actionIndex, lastNumberIndex, css)
		if !ok {
			if len(exps) > 0 {
				return exps, ok
			}

			r, _ := strconv.ParseFloat(strings.Join(css, ""), 64)
			exps = append(exps, Expression{result: r})
			ok = false

			return exps, ok
		}

		if action == "*" || action == "/" {
			if action == "*" {
				holdMult = true
			}

			if action == "/" {
				holdDiv = true
			}

			if len(exps) != 0 {
				actionIndex++
				lastNumberIndex = actionIndex
				action = ""
			}
		}

		var fn float64
		var fni int

		if len(exps) == 0 {
			if holdMult || holdDiv {
				if lastNumberIndex == 0 {
					fn, fni = extractFirstNumber(css, actionIndex, false)
				}

				if lastNumberIndex > 0 {
					fn, fni = extractFirstNumber(css, actionIndex, true)
				}
			}

			if !holdMult && !holdDiv {
				fn, fni = extractFirstNumber(css, actionIndex, false)
			}
		}

		if len(exps) > 0 && !holdMult {
			fn = exps[len(exps)-1].result
		}

		var ln float64
		var lni int

		var okk = false

		if fni > 0 {
			ln, lni, okk = extractLastNumber(css, fni)
			if !okk {
				return exps, false
			}
		}

		if fni == 0 {
			ln, lni, okk = extractLastNumber(css, actionIndex)
			if !okk {
				return exps, false
			}
		}

		if fni == lni && action == "" {
			if holdMult {
				exps = append(exps, Expression{
					result: execAction(fn, ln, "*"),
				})
			}

			if holdDiv {
				exps = append(exps, Expression{
					result: execAction(fn, ln, "/"),
				})
			}

			break
		}

		if action == "" {
			action = "+"
		}

		lastNumberIndex = lni

		exps = append(exps, Expression{
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

	if lastActionIndex > len(css)-1 || lastNumberIndex >= len(css)-1 {
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

func extractFirstNumber(ss []string, actionIndex int, holdMult bool) (fn float64, fni int) {
	if holdMult {
		fns := []string{}
		fnss := []string{}

		var ln = false

		for i, s := range ss[actionIndex:] {
			if val, okk := validActions[s]; okk {
				if ln {
					break
				}

				fnss = append(fnss, val)
				continue
			}

			if _, okk := validNumbers[s]; okk {
				fns = append(fns, s)
				ln = true
				fni = i + actionIndex
				continue
			}

			if s == "." {
				fns = append(fns, s)
			}
		}

		fs := extractSymbol(fnss)
		fns = append([]string{fs}, fns...)
		fn, _ = strconv.ParseFloat(strings.Join(fns, ""), 64)
	}

	if !holdMult {
		fn, _ = strconv.ParseFloat(strings.Join(gsub(ss[0:actionIndex]), ""), 64)
	}

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
				if len(rm)-1 < i+2 {
					ok = false
					break
				}
			}

			ns = append(ns, val)
		}

		if _, okk := validNumbers[s]; okk {
			lns = append(lns, s)
			lnf = true

			lni = i + start
		}

		if s == "." {
			lns = append(lns, s)

			lni = i + start + 1
		}
	}

	if !ok {
		return
	}

	ls := extractSymbol(ns)

	if ls != "" {
		lns = append([]string{ls}, lns...)
	}

	ln, _ = strconv.ParseFloat(strings.Join(lns, ""), 64)
	return
}

func extractSymbol(symbols []string) string {
	var ls string
	if len(symbols) > 1 {
		var p = 0
		for p < len(symbols) {
			if p == len(symbols)-1 {
				break
			}

			firstSymbol := symbols[p]
			lastSymbol := symbols[p+1]

			if firstSymbol == "-" && lastSymbol == "-" {
				ls = "+"
			}

			if firstSymbol == "+" && lastSymbol == "-" || firstSymbol == "-" && lastSymbol == "+" {
				ls = "-"
			}

			if firstSymbol == "+" && lastSymbol == "+" {
				ls = "+"
			}

			p++
		}
	}

	if len(symbols) == 1 {
		ls = symbols[0]
	}

	return ls
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
