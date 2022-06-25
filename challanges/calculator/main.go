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
	startIndex  int
	endIndex    int
}

func searchExpressions(css []string) (exps []Expression, ok bool) {
	ok = true

	actionsIndex := map[string]int{}
	var calculated = false

	for i, a := range css {
		if i == 0 {
			continue
		}

		if _, okk := validActions[a]; okk {
			if _, okkk := actionsIndex[a]; !okkk {
				actionsIndex[a] = i
			}
		}
	}

restart:
	if _, okk := actionsIndex["*"]; okk {
		actionIndex := actionsIndex["*"]
		exps, ok = search(css, actionIndex, exps, "*")
		if !ok {
			return exps, ok
		}

		delete(actionsIndex, "*")
		calculated = true
		goto restart
	}

	if _, okk := actionsIndex["/"]; okk {
		actionIndex := actionsIndex["/"]
		exps, ok = search(css, actionIndex, exps, "/")
		if !ok {
			return exps, ok
		}

		delete(actionsIndex, "/")
		calculated = true
		goto restart
	}

	if _, o := actionsIndex["-"]; o {
		if _, ko := actionsIndex["+"]; ko {
			if actionsIndex["+"] < actionsIndex["-"] {
				actionIndex := actionsIndex["+"]
				exps, ok = search(css, actionIndex, exps, "+")
				if !ok {
					return exps, ok
				}

				delete(actionsIndex, "+")
				delete(actionsIndex, "-")
				calculated = true
				goto restart
			}

			if actionsIndex["-"] > actionsIndex["+"] {
				actionIndex := actionsIndex["-"]
				exps, ok = search(css, actionIndex, exps, "-")
				if !ok {
					return exps, ok
				}

				delete(actionsIndex, "+")
				delete(actionsIndex, "-")
				calculated = true
				goto restart
			}
		}
	}

	if _, ok := actionsIndex["-"]; ok {
		actionIndex := actionsIndex["-"]
		exps, ok = search(css, actionIndex, exps, "-")
		if !ok {
			return exps, ok
		}

		delete(actionsIndex, "-")
		calculated = true
		goto restart
	}

	if _, ok := actionsIndex["+"]; ok {
		actionIndex := actionsIndex["+"]
		exps, ok = search(css, actionIndex, exps, "+")
		if !ok {
			return exps, ok
		}

		delete(actionsIndex, "+")
		calculated = true
		goto restart
	}

	if len(actionsIndex) == 0 && !calculated {
		r, _ := strconv.ParseFloat(strings.Join(css, ""), 64)
		exps = append(exps, Expression{
			result: r,
		})

		return exps, false
	}

	lastIndexCalculated := 0
	lastExpression := Expression{}

	for _, exp := range exps {
		if exp.endIndex > lastIndexCalculated {
			lastIndexCalculated = exp.endIndex
			lastExpression = exp
		}

		if exp.endIndex == lastIndexCalculated {
			if exp.result > lastExpression.result {
				lastExpression = exp
			}
		}
	}

	if len(css[lastIndexCalculated:]) > 0 {
		fn := lastExpression.result
		var ln float64
		var lni int
		var okk bool

		ln, lni, okk = extractLastNumber(css, lastIndexCalculated)
		if !okk {
			return exps, false
		}

		exps = append(exps, Expression{
			firstNumber: fn,
			lastNumber:  ln,
			result:      execAction(fn, ln, "+"),
			endIndex:    lni,
		})
	}

	return
}

func search(css []string, actionIndex int, exps []Expression, action string) ([]Expression, bool) {
	for {
		var fn float64
		var fni int

		var fnAlreadyChecked = false
		var lnAlreadyChecked = false

		for _, e := range exps {
			if e.startIndex == actionIndex+1 {
				lnAlreadyChecked = true
				break
			}

			if e.startIndex == actionIndex-1 {
				lnAlreadyChecked = true
				break
			}
		}

		if !fnAlreadyChecked {
			if len(exps) == 0 {
				fn, fni = extractFirstNumber(css, actionIndex)
			}

			if len(exps) > 0 && exps[len(exps)-1].startIndex < actionIndex {
				fn = exps[len(exps)-1].result
				fni = exps[len(exps)-1].startIndex
			}

			if len(exps) > 0 && exps[len(exps)-1].startIndex > actionIndex {
				fn, fni = extractFirstNumber(css, actionIndex)
			}
		}

		if fnAlreadyChecked {
			for _, e := range exps {
				if e.startIndex == actionIndex-1 {
					fn = e.result
				}
			}
		}

		var ln float64
		var lni int

		if lnAlreadyChecked {
			for _, e := range exps {
				if e.startIndex == actionIndex+1 {
					ln = e.result
					lni = e.endIndex
				}
			}
		}

		if !lnAlreadyChecked {
			var okk = false

			ln, lni, okk = extractLastNumber(css, actionIndex)
			if !okk {
				return exps, false
			}
		}

		exps = append(exps, Expression{
			firstNumber: fn,
			lastNumber:  ln,
			action:      action,
			result:      execAction(fn, ln, action),
			startIndex:  fni,
			endIndex:    lni,
		})

		if lni >= len(css)-1 {
			return exps, false
		}

		for _, s := range css[lni+1:] {
			if s != " " && s != "(" && s != ")" {
				return exps, true
			}
		}

		return exps, false
	}
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

func extractFirstNumber(ss []string, actionIndex int) (fn float64, fni int) {
	fnActions := []string{}
	fNumbers := []string{}

	i := actionIndex - 1
	for {
		if i >= 0 {
			s := ss[i]
			if _, okk := validActions[s]; okk {
				if len(fnActions) == 0 {
					fnActions = append(fnActions, s)
				}

				if len(fnActions) > 0 {
					break
				}
			}

			if _, okk := validNumbers[s]; okk {
				fNumbers = append(fNumbers, s)
				fni = i
			}

			i--
		}

		if i < 0 {
			break
		}
	}

	revFnumbers := []string{}
	for i := len(fNumbers) - 1; i >= 0; i-- {
		revFnumbers = append(revFnumbers, fNumbers[i])
	}

	symbolSignal := extractSymbol(fnActions)

	if symbolSignal != "" {
		revFnumbers = append([]string{symbolSignal}, revFnumbers...)
	}

	fn, _ = strconv.ParseFloat(strings.Join(revFnumbers, ""), 64)
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
