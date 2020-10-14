package main

import (
	"fmt"
	"os"
)

func main() {
	var (
		i int
		v string
	)
	for i, v = range os.Args[1:] {
		fmt.Printf("The index %v has the value %v\n", i+1, v)
	}
	if v != "" {
		fmt.Println("I Still have the value of last index: ", v)
	}
}
