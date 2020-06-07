package main

import "fmt"

func main() {
	var (
		counter byte = 10
	)

	p := &counter
	v := *p

	fmt.Printf("The variable p stores the memory address %v, and its value is %v.\n", p, v)
	fmt.Printf("Its type is %T, and its value is of type %T.\n", p, v)
}
