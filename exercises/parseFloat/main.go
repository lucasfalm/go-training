package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	args := os.Args[1]

	feet, _ := strconv.ParseFloat(args, 64)

	meters := feet * 0.3048

	fmt.Printf("%g feets is %g meters\n", feet, meters)

	intNumber := 2

	fmt.Println("O número em string é: " + strconv.Itoa(intNumber))

	secondArg := os.Args[2]
	fmt.Println(strings.ToUpper(secondArg))
}
