package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {

	var (
		args = os.Args
	)

	if len(args) >= 2 {
		n := args[1]
		age, err := strconv.Atoi(n)
		if err != nil {
			fmt.Println("Error! ", err)
			return
		} else {
			fmt.Printf("Success converion %q into %v\n", n, age)
		}
	} else {
		fmt.Println("Write an input to be converted")
		return
	}
}
