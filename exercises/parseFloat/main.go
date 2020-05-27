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
		i := args[1]
		feet, err := strconv.ParseFloat(i, 64)
		if err == nil {
			meters := feet * 0.3048

			fmt.Printf("%g feets is %g meters\n", feet, meters)
		} else {
			fmt.Printf("Error -> %v", err)
			return
		}
	} else {
		fmt.Println("Write a number in feets")
		return
	}
}
