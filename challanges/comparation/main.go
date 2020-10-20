package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/flucas97/go-trainning/challanges/comparation/calc"
)

func main() {
	var (
		args = os.Args
	)

	if len(args) != 2 {
		fn, _ := strconv.Atoi(args[1])
		ln, _ := strconv.Atoi(args[2])

		r, err := calc.Compare(fn, ln)
		if err == nil {
			fmt.Printf("%v is bigger than %v? -> %v\n", fn, ln, r)
		} else {
			fmt.Printf("Error! -> %v", err)
		}
	} else {
		fmt.Println("Pass two arguments")
	}
}
