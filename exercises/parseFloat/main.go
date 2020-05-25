package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	args := os.Args[1]

	feet, _ := strconv.ParseFloat(args, 64)

	meters := feet * 0.3048

	fmt.Printf("%g feets is %f meters\n", feet, meters)
}
