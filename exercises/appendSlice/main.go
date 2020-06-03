package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	var (
		args    = os.Args[1:]
		numbers []int
		uniques []int
	)

	if len(args) >= 1 {
		for _, n := range args {
			nm, _ := strconv.Atoi(n)
			numbers = append(numbers, nm)
		}
	}
	//trying to check if the number is already in uniques, if not true so append into it, but fails :(
loop:
	for _, m := range numbers {
		uniques = append(uniques, m)
		for _, v := range uniques {
			if m == v {
				continue loop
			} else {
				uniques = append(uniques, m)
				continue loop
			}
		}
	}

	fmt.Println(numbers)
	fmt.Println(uniques)
}
